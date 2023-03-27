package domain

import (
	"context"
	"fmt"
	pb "github.com/frozosea/fmc-pb/schedule-tracking"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/scheduler"
	"schedule-tracking/pkg/util"
	"time"
)

type converter struct {
}

func (c *converter) convertAddOnTrack(r *pb.AddOnTrackRequest, userId int) *BaseTrackReq {
	return &BaseTrackReq{
		Numbers:             r.GetNumbers(),
		UserId:              int64(userId),
		Time:                r.GetTime(),
		Emails:              r.GetEmails(),
		EmailMessageSubject: r.GetEmailMessageSubject(),
	}
}
func (c *converter) convertBaseAddOnTrackResponse(r []*BaseAddOnTrackResponse) []*pb.BaseAddOnTrackResponse {
	var res []*pb.BaseAddOnTrackResponse
	for _, v := range r {
		res = append(res, &pb.BaseAddOnTrackResponse{
			Success:     v.success,
			Number:      v.number,
			NextRunTime: v.nextRunTime.Unix(),
		})
	}
	return res
}
func (c *converter) convertAddOnTrackResponse(r *AddOnTrackResponse) *pb.AddOnTrackResponse {
	return &pb.AddOnTrackResponse{
		BaseResponse:   c.convertBaseAddOnTrackResponse(r.result),
		AlreadyOnTrack: r.alreadyOnTrack,
	}
}

type Grpc struct {
	service      *Service
	logger       logging.ILogger
	converter    *converter
	tokenManager util.ITokenManager
	pb.UnimplementedScheduleTrackingServer
}

func NewGrpc(controller *Service, logger logging.ILogger, manager util.ITokenManager) *Grpc {
	return &Grpc{
		service:                             controller,
		logger:                              logger,
		converter:                           &converter{},
		tokenManager:                        manager,
		UnimplementedScheduleTrackingServer: pb.UnimplementedScheduleTrackingServer{},
	}
}

func (s *Grpc) AddContainersOnTrack(ctx context.Context, r *pb.AddOnTrackRequest) (*pb.AddOnTrackResponse, error) {
	userId, err := s.tokenManager.GetUserIdFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	res, err := s.service.AddContainerNumbersOnTrack(ctx, s.converter.convertAddOnTrack(r, userId))
	if err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return s.converter.convertAddOnTrackResponse(res), status.Error(codes.NotFound, "cannot find job with this id")
		case *NumberDoesntBelongThisUserError:
			return s.converter.convertAddOnTrackResponse(res), status.Error(codes.PermissionDenied, "cannot find number in your account")
		default:
			return s.converter.convertAddOnTrackResponse(res), status.Error(codes.Internal, err.Error())
		}
	}
	return s.converter.convertAddOnTrackResponse(res), nil
}

func (s *Grpc) AddBillNosOnTrack(ctx context.Context, r *pb.AddOnTrackRequest) (*pb.AddOnTrackResponse, error) {
	userId, err := s.tokenManager.GetUserIdFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	res, err := s.service.AddBillNumbersOnTrack(ctx, s.converter.convertAddOnTrack(r, userId))
	if err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return &pb.AddOnTrackResponse{}, status.Error(codes.NotFound, "cannot find job with this id")
		case *NumberDoesntBelongThisUserError:
			return &pb.AddOnTrackResponse{}, status.Error(codes.PermissionDenied, "cannot find number in your account")
		case *scheduler.TimeParseError:
			return &pb.AddOnTrackResponse{}, status.Error(codes.InvalidArgument, err.Error())
		default:
			return &pb.AddOnTrackResponse{}, status.Error(codes.Internal, err.Error())
		}
	}
	return s.converter.convertAddOnTrackResponse(res), nil
}

func (s *Grpc) deleteFromTracking(ctx context.Context, r *pb.DeleteFromTrackingRequest, isContainer bool) (*emptypb.Empty, error) {
	userId, err := s.tokenManager.GetUserIdFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if err := s.service.DeleteFromTracking(ctx, int64(userId), isContainer, r.GetNumbers()); err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
		case *NumberDoesntBelongThisUserError:
			return &emptypb.Empty{}, status.Error(codes.PermissionDenied, err.Error())
		default:
			go func() {
				for _, v := range r.GetNumbers() {
					s.logger.ExceptionLog(fmt.Sprintf(`delete Number: %s for user-pb %d from tracking err: %s`, v, userId, err.Error()))
				}
			}()
			return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *Grpc) DeleteContainersFromTracking(ctx context.Context, r *pb.DeleteFromTrackingRequest) (*emptypb.Empty, error) {
	return s.deleteFromTracking(ctx, r, true)
}

func (s *Grpc) DeleteBillsFromTracking(ctx context.Context, r *pb.DeleteFromTrackingRequest) (*emptypb.Empty, error) {
	return s.deleteFromTracking(ctx, r, false)
}

func (s *Grpc) GetInfoAboutTrack(ctx context.Context, r *pb.GetInfoAboutTrackRequest) (*pb.GetInfoAboutTrackResponse, error) {
	resp, err := s.service.GetInfoAboutTracking(ctx, r.GetNumber(), r.GetUserId())
	if err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return nil, status.Error(codes.NotFound, "task with this id was not found")
		case *NumberDoesntBelongThisUserError:
			return &pb.GetInfoAboutTrackResponse{}, status.Error(codes.PermissionDenied, err.Error())
		default:
			go s.logger.ExceptionLog(fmt.Sprintf(`get info about tracking err: %s`, err.Error()))
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &pb.GetInfoAboutTrackResponse{
		Number:      resp.Number,
		IsOnTrack:   resp.IsOnTrack,
		IsContainer: resp.IsContainer,
		ScheduleTrackingInfo: &pb.ScheduleTrackingInfo{
			Time:    resp.ScheduleTrackingInfo.Time,
			Subject: resp.ScheduleTrackingInfo.Subject,
			Emails:  resp.ScheduleTrackingInfo.Emails,
		},
	}, nil
}

func (s *Grpc) GetTimeZone(context.Context, *emptypb.Empty) (*pb.GetTimeZoneResponse, error) {
	t := time.Now()
	zone, _ := t.Zone()
	return &pb.GetTimeZoneResponse{TimeZone: fmt.Sprintf(`UTC%s`, zone)}, nil
}

func (s *Grpc) update(ctx context.Context, r *pb.AddOnTrackRequest, isContainer bool) (*emptypb.Empty, error) {
	userId, err := s.tokenManager.GetUserIdFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if err := s.service.Update(ctx, s.converter.convertAddOnTrack(r, userId), isContainer); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *Grpc) UpdateContainer(ctx context.Context, r *pb.AddOnTrackRequest) (*emptypb.Empty, error) {
	return s.update(ctx, r, true)
}

func (s *Grpc) UpdateBill(ctx context.Context, r *pb.AddOnTrackRequest) (*emptypb.Empty, error) {
	return s.update(ctx, r, false)
}
