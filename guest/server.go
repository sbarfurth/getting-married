package guest

import (
	"context"
	"fmt"
	guestv1 "getting-married/gen/guest/v1"

	"connectrpc.com/connect"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var bucketName = []byte("guests")

type Server struct {
	boltDB *bolt.DB
}

func NewServer(boltDB *bolt.DB) (*Server, error) {
	if err := boltDB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(bucketName); err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &Server{
		boltDB: boltDB,
	}, nil
}

func (s *Server) GetParty(ctx context.Context, req *connect.Request[guestv1.GetPartyRequest]) (*connect.Response[guestv1.GetPartyResponse], error) {
	var partyBytes []byte
	if err := s.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		copy(partyBytes, b.Get([]byte(req.Msg.GetName())))
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

func (s *Server) CreateParty(ctx context.Context, req *connect.Request[guestv1.CreatePartyRequest]) (*connect.Response[guestv1.CreatePartyResponse], error) {
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

	if err := s.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		b.Put([]byte(party.GetName()), partyBytes)
		return nil
	}); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&guestv1.CreatePartyResponse{
		Party: party,
	})
	res.Header().Set("Guest-Version", "v1")
	return res, nil
}
