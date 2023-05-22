package schedule_tracking

import (
	"context"
	"fmt"
	pb "github.com/frozosea/fmc-pb/v2/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/pkg/cache"
	"user-api/pkg/logging"
)

type Grpc struct {
	repository IRepository
	logger     logging.ILogger
	cache      cache.ICache
	pb.UnimplementedScheduleTrackingServer
}

func NewGrpc(repository IRepository, logger logging.ILogger, cache cache.ICache) *Grpc {
	return &Grpc{repository: repository, logger: logger, cache: cache, UnimplementedScheduleTrackingServer: pb.UnimplementedScheduleTrackingServer{}}
}
func (s *Grpc) MarkBillNoOnTrack(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkBillNoOnTrack(ctx, r.GetNumber(), r.GetUserId()); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark bill no with number %s for user %d on track error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d add on track`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Grpc) MarkContainerOnTrack(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkContainerOnTrack(ctx, r.GetNumber(), r.GetUserId()); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark container no with number %s for user %d on track error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark container no with number %s for user %d add on track`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Grpc) MarkContainerWasArrived(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkContainerWasArrived(ctx, r.GetNumber(), r.GetUserId()); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark container with number %s for user %d was arrived error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d was arrived`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil

}
func (s *Grpc) MarkBillNoWasArrived(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkBillNoWasArrived(ctx, r.GetNumber(), r.GetUserId()); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark bill no with number %s for user %d was arrived error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d was arrived`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil

}
func (s *Grpc) MarkBillNoWasRemovedFromTrack(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkBillNoWasRemovedFromTrack(ctx, r.GetNumber(), r.GetUserId()); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark bill no with number %s for user %d remove from track error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d was removed from track`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Grpc) MarkContainerWasRemovedFromTrack(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkContainerWasRemovedFromTrack(ctx, r.GetNumber(), r.GetUserId()); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark container with number %s for user %d remove from track error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark container with number %s for user %d was removed from track`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Grpc) CheckNumberExists(ctx context.Context, r *pb.CheckNumberExistsRequest) (*pb.CheckNumberExistsResponse, error) {
	exist, err := s.repository.CheckNumberExists(ctx, r.GetNumber(), r.GetUserId(), r.GetIsContainer())
	if err != nil {
		return &pb.CheckNumberExistsResponse{}, status.Error(codes.NotFound, err.Error())
	}
	return &pb.CheckNumberExistsResponse{Exists: exist}, nil
}
func (s *Grpc) MarkContainerIsNotArrived(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkContainerNumberWasNotArrived(ctx, r.GetNumber(), r.GetUserId()); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark container no with number %s for user %d was not arrived error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark container with number %s for user %d was not arrived`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}

func (s *Grpc) MarkBillIsNotArrived(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkBillNumberWasNotArrived(ctx, r.GetNumber(), r.GetUserId()); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark bill no with number %s for user %d was not arrived error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill with number %s for user %d was not arrived`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
