package scac

import (
	"context"
	pb "github.com/frozosea/fmc-pb/v2/tracking"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Grpc struct {
	service *Service
	*pb.UnimplementedScacServiceServer
}

func NewGrpc(service *Service) *Grpc {
	return &Grpc{service: service, UnimplementedScacServiceServer: &pb.UnimplementedScacServiceServer{}}
}

func (g *Grpc) GetContainerScac(_ context.Context, _ *emptypb.Empty) (*pb.GetAllScacResponse, error) {
	linesList, err := g.service.GetContainerLines()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return linesList.ToGRPC(), nil
}
func (g *Grpc) GetBillScac(_ context.Context, _ *emptypb.Empty) (*pb.GetAllScacResponse, error) {
	linesList, err := g.service.GetBillLines()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return linesList.ToGRPC(), nil
}
