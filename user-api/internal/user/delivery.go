package user

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/internal/domain"
	pb "user-api/pkg/proto"
)

type converter struct{}

func (c *converter) convertContainerOrBillToGrpc(r []*domain.Container) []*pb.ContainerResponse {
	var arr []*pb.ContainerResponse
	for _, v := range r {
		arr = append(arr, &pb.ContainerResponse{
			Id:        v.Id,
			Number:    v.Number,
			IsOnTrack: v.IsOnTrack,
		})
	}
	return arr
}
func (c *converter) addContainerOrBillConvert(r *pb.AddContainerToAccountRequest) []string {
	var containers []string
	for _, v := range r.GetContainer() {
		containers = append(containers, v.GetNumber())
	}
	return containers
}

type Grpc struct {
	controller *Service
	converter  converter
	pb.UnimplementedUserServer
}

func NewGrpc(controller *Service) *Grpc {
	return &Grpc{controller: controller, converter: converter{}, UnimplementedUserServer: pb.UnimplementedUserServer{}}
}

func (s *Grpc) AddContainerToAccount(ctx context.Context, r *pb.AddContainerToAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.AddContainerToAccount(ctx, int(r.GetUserId()), s.converter.addContainerOrBillConvert(r))
}
func (s *Grpc) AddBillNumberToAccount(ctx context.Context, r *pb.AddContainerToAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.AddBillNumberToAccount(ctx, int(r.GetUserId()), s.converter.addContainerOrBillConvert(r))
}
func (s *Grpc) DeleteContainersFromAccount(ctx context.Context, r *pb.DeleteContainersFromAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.DeleteContainersFromAccount(ctx, int(r.GetUserId()), r.GetNumberIds())

}
func (s *Grpc) DeleteBillNumbersFromAccount(ctx context.Context, r *pb.DeleteContainersFromAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.DeleteBillNumbersFromAccount(ctx, int(r.GetUserId()), r.GetNumberIds())

}
func (s *Grpc) GetAll(ctx context.Context, r *pb.GetAllContainersFromAccountRequest) (*pb.GetAllContainersResponse, error) {
	res, err := s.controller.GetAllContainers(ctx, int(r.GetUserId()))
	if err != nil {
		return &pb.GetAllContainersResponse{}, err
	}
	return &pb.GetAllContainersResponse{
		BillNumbers: s.converter.convertContainerOrBillToGrpc(res.BillNumbers),
		Containers:  s.converter.convertContainerOrBillToGrpc(res.Containers),
	}, nil
}
