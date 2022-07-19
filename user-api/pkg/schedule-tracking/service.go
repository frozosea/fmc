package schedule_tracking

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/internal/logging"
	pb "user-api/pkg/proto"
)

type Service struct {
	repository IRepository
	logger     logging.ILogger
	pb.UnimplementedScheduleTrackingServer
}

func NewService(repository IRepository, logger logging.ILogger) *Service {
	return &Service{repository: repository, logger: logger, UnimplementedScheduleTrackingServer: pb.UnimplementedScheduleTrackingServer{}}
}

func (s *Service) MarkBillNoOnTrack(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkBillNoOnTrack(ctx, r.GetNumber(), int(r.GetUserId())); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark bill no with number %s for user %d on track error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d add on track`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Service) MarkContainerOnTrack(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkContainerOnTrack(ctx, r.GetNumber(), int(r.GetUserId())); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark container with number %s for user %d on track error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d add on track`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Service) MarkContainerWasArrived(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkContainerWasArrived(ctx, r.GetNumber(), int(r.GetUserId())); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark container with number %s for user %d was arrived error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d was arrived`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Service) MarkBillNoWasArrived(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkBillNoWasArrived(ctx, r.GetNumber(), int(r.GetUserId())); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark bill no with number %s for user %d was arrived error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d was arrived`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Service) MarkBillNoWasRemovedFromTrack(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkBillNoWasRemovedFromTrack(ctx, r.GetNumber(), int(r.GetUserId())); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark bill no with number %s for user %d remove from track error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark bill no with number %s for user %d was removed from track`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
func (s *Service) MarkContainerWasRemovedFromTrack(ctx context.Context, r *pb.AddMarkOnTrackingRequest) (*emptypb.Empty, error) {
	if err := s.repository.AddMarkContainerWasRemovedFromTrack(ctx, r.GetNumber(), int(r.GetUserId())); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`mark container with number %s for user %d remove from track error: %s`, r.GetNumber(), r.GetUserId(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go s.logger.InfoLog(fmt.Sprintf(`mark container with number %s for user %d was removed from track`, r.GetNumber(), r.GetUserId()))
	return &emptypb.Empty{}, nil
}
