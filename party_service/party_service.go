package partyservice

import (
	"context"
	"getting-married/proto_out/github.com/sbarfurth/getting-married/partypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ partypb.PartyServiceServer = &Server{}

type Server struct {
	partypb.UnimplementedPartyServiceServer
}

func (s *Server) GetParty(ctx context.Context, req *partypb.GetPartyRequest) (*partypb.Party, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *Server) ListParties(ctx context.Context, req *partypb.ListPartiesRequest) (*partypb.ListPartiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *Server) CreateParty(ctx context.Context, req *partypb.CreatePartyRequest) (*partypb.Party, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *Server) UpdateParty(ctx context.Context, req *partypb.UpdatePartyRequest) (*partypb.Party, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}

func (s *Server) DeleteParty(ctx context.Context, req *partypb.DeletePartyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}
