package user

import (
	"context"
	pb "github.com/frozosea/fmc-pb/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/internal/domain"
)

type converter struct{}

func (c *converter) convertContainerOrBillToGrpc(r []*domain.Container) []*pb.ContainerResponse {
	var arr []*pb.ContainerResponse
	for _, v := range r {
		pbResp := &pb.ContainerResponse{
			Number:      v.Number,
			IsOnTrack:   v.IsOnTrack,
			IsContainer: v.IsContainer,
		}
		if v.IsOnTrack {
			pbResp.ScheduleTrackingObject = &pb.ScheduleTrackingObject{
				Time:    v.ScheduleTrackingInfo.Time,
				Emails:  v.ScheduleTrackingInfo.Emails,
				Subject: v.ScheduleTrackingInfo.Subject,
			}
		}
		arr = append(arr, pbResp)
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
	service   *Service
	converter converter
	pb.UnimplementedUserServer
}

func NewGrpc(controller *Service) *Grpc {
	return &Grpc{service: controller, converter: converter{}, UnimplementedUserServer: pb.UnimplementedUserServer{}}
}

func (s *Grpc) AddContainerToAccount(ctx context.Context, r *pb.AddContainerToAccountRequest) (*emptypb.Empty, error) {
	if err := s.service.AddContainerToAccount(ctx, int(r.GetUserId()), s.converter.addContainerOrBillConvert(r)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (s *Grpc) AddBillNumberToAccount(ctx context.Context, r *pb.AddContainerToAccountRequest) (*emptypb.Empty, error) {
	if err := s.service.AddBillNumberToAccount(ctx, int(r.GetUserId()), s.converter.addContainerOrBillConvert(r)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (s *Grpc) DeleteContainersFromAccount(ctx context.Context, r *pb.DeleteContainersFromAccountRequest) (*emptypb.Empty, error) {
	if err := s.service.DeleteContainersFromAccount(ctx, int(r.GetUserId()), r.GetNumbers()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil

}
func (s *Grpc) DeleteBillNumbersFromAccount(ctx context.Context, r *pb.DeleteContainersFromAccountRequest) (*emptypb.Empty, error) {
	if err := s.service.DeleteBillNumbersFromAccount(ctx, int(r.GetUserId()), r.GetNumbers()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil

}
func (s *Grpc) GetAll(ctx context.Context, r *pb.GetAllContainersFromAccountRequest) (*pb.GetAllContainersResponse, error) {
	res, err := s.service.GetAllContainers(ctx, int(r.GetUserId()))
	if err != nil {
		return &pb.GetAllContainersResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetAllContainersResponse{
		BillNumbers: s.converter.convertContainerOrBillToGrpc(res.BillNumbers),
		Containers:  s.converter.convertContainerOrBillToGrpc(res.Containers),
	}, nil
}
