package guest

import (
	"context"
	"fmt"
	guestv1 "getting-married/gen/guest/v1"
	"net/http"
	"slices"

	"connectrpc.com/connect"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var bucketName = []byte("guests")

type Server struct {
	boltDB        *bolt.DB
	adminPassword string
}

func NewServer(boltDB *bolt.DB, adminPassword string) (*Server, error) {
	if err := boltDB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(bucketName); err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &Server{
		boltDB:        boltDB,
		adminPassword: adminPassword,
	}, nil
}

func (s *Server) isAdmin(headers http.Header) bool {
	return headers.Get("authorization") == s.adminPassword
}

func (s *Server) GetParty(ctx context.Context, req *connect.Request[guestv1.GetPartyRequest]) (*connect.Response[guestv1.GetPartyResponse], error) {
	var partyBytes []byte
	if err := s.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		partyBytes = b.Get([]byte(req.Msg.GetName()))
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if partyBytes == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("party not found"))
	}
	party := &guestv1.Party{}
	if err := proto.Unmarshal(partyBytes, party); err != nil {
		return nil, connect.NewError(connect.CodeDataLoss, err)
	}
	res := connect.NewResponse(&guestv1.GetPartyResponse{
		Party: party,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) UpdatePartyContactInfo(ctx context.Context, req *connect.Request[guestv1.UpdatePartyContactInfoRequest]) (*connect.Response[guestv1.UpdatePartyContactInfoResponse], error) {
	updatedParty, err := s.updateParty(req.Msg.GetName(), func(party *guestv1.Party) error {
		party.Address = req.Msg.GetAddress()
		party.Contact = req.Msg.GetContact()
		return nil
	})
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&guestv1.UpdatePartyContactInfoResponse{
		Party: updatedParty,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) UpdatePartyRsvpResponse(ctx context.Context, req *connect.Request[guestv1.UpdatePartyRsvpResponseRequest]) (*connect.Response[guestv1.UpdatePartyRsvpResponseResponse], error) {
	updatedParty, err := s.updateParty(req.Msg.GetName(), func(party *guestv1.Party) error {
		now := timestamppb.Now()
		for _, guest := range party.Guests {
			if guest.GetRsvp() == nil {
				guest.Rsvp = &guestv1.RSVP{
					Response:  req.Msg.GetResponse(),
					UpdatedAt: now,
				}
			} else {
				guest.GetRsvp().Response = req.Msg.GetResponse()
				guest.GetRsvp().UpdatedAt = now
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&guestv1.UpdatePartyRsvpResponseResponse{
		Party: updatedParty,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) UpdateGuestRsvp(ctx context.Context, req *connect.Request[guestv1.UpdateGuestRsvpRequest]) (*connect.Response[guestv1.UpdateGuestRsvpResponse], error) {
	updatedParty, err := s.updateParty(req.Msg.GetName(), func(party *guestv1.Party) error {
		if int(req.Msg.GetIndex()) >= len(party.GetGuests()) {
			return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unknown guest index"))
		}
		party.GetGuests()[req.Msg.GetIndex()].Rsvp = req.Msg.GetRsvp()
		party.GetGuests()[req.Msg.GetIndex()].Rsvp.UpdatedAt = timestamppb.Now()
		return nil
	})
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&guestv1.UpdateGuestRsvpResponse{
		Party: updatedParty,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) AddGuest(ctx context.Context, req *connect.Request[guestv1.AddGuestRequest]) (*connect.Response[guestv1.AddGuestResponse], error) {
	updatedParty, err := s.updateParty(req.Msg.GetName(), func(party *guestv1.Party) error {
		if !party.GetAllowGuestSelfService() {
			return connect.NewError(connect.CodePermissionDenied, fmt.Errorf("guest self-service is not allowed"))
		}
		if int(party.GetMaxGuests()) <= len(party.GetGuests()) {
			return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("maximum number of guests (%d) reached", party.GetMaxGuests()))
		}
		party.Guests = append(party.GetGuests(), &guestv1.Guest{
			FirstName: req.Msg.GetGuest().GetFirstName(),
			LastName:  req.Msg.GetGuest().GetLastName(),
			Rsvp: &guestv1.RSVP{
				Response:  guestv1.RSVP_RESPONSE_ACCEPTED,
				UpdatedAt: timestamppb.Now(),
			},
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&guestv1.AddGuestResponse{
		Party: updatedParty,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) RemoveGuest(ctx context.Context, req *connect.Request[guestv1.RemoveGuestRequest]) (*connect.Response[guestv1.RemoveGuestResponse], error) {
	updatedParty, err := s.updateParty(req.Msg.GetName(), func(party *guestv1.Party) error {
		if !party.GetAllowGuestSelfService() {
			return connect.NewError(connect.CodePermissionDenied, fmt.Errorf("guest self-service is not allowed"))
		}
		idx := int(req.Msg.GetIndex())
		if idx >= len(party.GetGuests()) {
			return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("specified index does not exist"))
		}
		if len(party.GetGuests()) <= 1 {
			return connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("cannot remove last guest"))
		}
		party.Guests = slices.Delete(party.GetGuests(), idx, idx+1)
		return nil
	})
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&guestv1.RemoveGuestResponse{
		Party: updatedParty,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) ListParties(ctx context.Context, req *connect.Request[guestv1.ListPartiesRequest]) (*connect.Response[guestv1.ListPartiesResponse], error) {
	if !s.isAdmin(req.Header()) {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("admin access required"))
	}
	var parties []*guestv1.Party
	if err := s.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()

		for name, partyBytes := c.First(); name != nil; name, partyBytes = c.Next() {
			party := &guestv1.Party{}
			if err := proto.Unmarshal(partyBytes, party); err != nil {
				return err
			}
			parties = append(parties, party)
		}

		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&guestv1.ListPartiesResponse{
		Parties: parties,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) CreateParty(ctx context.Context, req *connect.Request[guestv1.CreatePartyRequest]) (*connect.Response[guestv1.CreatePartyResponse], error) {
	if !s.isAdmin(req.Header()) {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("admin access required"))
	}
	party := req.Msg.GetParty()

	if party == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("party must be specified"))
	}

	party.Name = uuid.NewString()
	party.UpdatedAt = timestamppb.Now()

	partyBytes, err := proto.Marshal(party)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	if err := s.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Put([]byte(party.GetName()), partyBytes)
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&guestv1.CreatePartyResponse{
		Party: party,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) UpdateParty(ctx context.Context, req *connect.Request[guestv1.UpdatePartyRequest]) (*connect.Response[guestv1.UpdatePartyResponse], error) {
	if !s.isAdmin(req.Header()) {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("admin access required"))
	}
	party := req.Msg.GetParty()

	if party == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("party must be specified"))
	}

	if party.GetName() == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("party must have a name"))
	}

	party.UpdatedAt = timestamppb.Now()

	partyBytes, err := proto.Marshal(party)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, err)
	}

	if err := s.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b.Get([]byte(party.GetName())) == nil {
			return fmt.Errorf("party %q does not exist", party.GetName())
		}
		return b.Put([]byte(party.GetName()), partyBytes)
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&guestv1.UpdatePartyResponse{
		Party: party,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) DeleteParty(ctx context.Context, req *connect.Request[guestv1.DeletePartyRequest]) (*connect.Response[emptypb.Empty], error) {
	if !s.isAdmin(req.Header()) {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("admin access required"))
	}
	if err := s.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Delete([]byte(req.Msg.GetName()))
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&emptypb.Empty{})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}

func (s *Server) updateParty(name string, updater func(*guestv1.Party) error) (*guestv1.Party, error) {
	var partyBytes []byte
	var updatedParty *guestv1.Party
	if err := s.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		partyBytes = b.Get([]byte(name))
		if partyBytes == nil {
			return connect.NewError(connect.CodeNotFound, fmt.Errorf("party not found"))
		}
		party := &guestv1.Party{}
		if err := proto.Unmarshal(partyBytes, party); err != nil {
			return connect.NewError(connect.CodeDataLoss, err)
		}
		if err := updater(party); err != nil {
			return err
		}
		party.UpdatedAt = timestamppb.Now()
		updatedBytes, err := proto.Marshal(party)
		if err != nil {
			return connect.NewError(connect.CodeUnknown, err)
		}
		if err := b.Put([]byte(name), updatedBytes); err != nil {
			return connect.NewError(connect.CodeUnknown, err)
		}
		updatedParty = party
		return nil
	}); err != nil {
		return nil, err
	}
	return updatedParty, nil

}
