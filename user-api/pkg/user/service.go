package user

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/pkg/domain"
	pb "user-api/pkg/proto"
)

type Service struct {
	controller *Controller
	converter  converter
	pb.UnimplementedUserServer
}

func NewService(controller *Controller) *Service {
	return &Service{controller: controller, converter: converter{}, UnimplementedUserServer: pb.UnimplementedUserServer{}}
}

func (s *Service) AddContainerToAccount(ctx context.Context, r *pb.AddContainerToAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.AddContainerToAccount(ctx, int(r.GetUserId()), s.converter.addContainerOrBillConvert(r))
}
func (s *Service) AddBillNumberToAccount(ctx context.Context, r *pb.AddContainerToAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.AddBillNumberToAccount(ctx, int(r.GetUserId()), s.converter.addContainerOrBillConvert(r))
}
func (s *Service) DeleteContainersFromAccount(ctx context.Context, r *pb.DeleteContainersFromAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.DeleteContainersFromAccount(ctx, int(r.GetUserId()), r.GetNumberIds())

}
func (s *Service) DeleteBillNumbersFromAccount(ctx context.Context, r *pb.DeleteContainersFromAccountRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.DeleteBillNumbersFromAccount(ctx, int(r.GetUserId()), r.GetNumberIds())

}
func (s *Service) GetAll(ctx context.Context, r *pb.GetAllContainersFromAccountRequest) (*pb.GetAllContainersResponse, error) {
	res, err := s.controller.GetAllContainers(ctx, int(r.GetUserId()))
	if err != nil {
		fmt.Println(err.Error())
		return &pb.GetAllContainersResponse{}, err
	}
	return &pb.GetAllContainersResponse{
		BillNumbers: s.converter.convertContainerOrBillToGrpc(res.BillNumbers),
		Containers:  s.converter.convertContainerOrBillToGrpc(res.Containers),
	}, nil
}

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
