package domain

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"schedule-tracking/pkg/logging"
	pb "schedule-tracking/pkg/proto"
	"schedule-tracking/pkg/scheduler"
	"time"
)

type converter struct {
}

func (c *converter) convertAddOnTrack(r *pb.AddOnTrackRequest, country string) *BaseTrackReq {
	return &BaseTrackReq{
		Numbers:             r.GetNumbers(),
		UserId:              r.GetUserId(),
		Country:             country,
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
func (c *converter) convertInterfaceArrayToStringArray(r []interface{}) []string {
	var outputArr []string
	for _, v := range r {
		outputArr = append(outputArr, fmt.Sprintf(`%v`, v))
	}
	return outputArr
}
func (c *converter) convertUpdateTaskRequest(r *pb.UpdateTaskRequest, country string) *BaseTrackReq {
	return &BaseTrackReq{
		Numbers:             r.Req.GetNumbers(),
		UserId:              r.Req.GetUserId(),
		Country:             country,
		Time:                r.Req.GetTime(),
		Emails:              r.Req.GetEmails(),
		EmailMessageSubject: r.Req.GetEmailMessageSubject(),
	}
}

type Grpc struct {
	controller *Service
	logger     logging.ILogger
	converter
	pb.UnimplementedScheduleTrackingServer
}

func NewGrpc(controller *Service, logger logging.ILogger) *Grpc {
	return &Grpc{controller: controller, logger: logger, converter: converter{}, UnimplementedScheduleTrackingServer: pb.UnimplementedScheduleTrackingServer{}}
}
func (s *Grpc) AddContainersOnTrack(ctx context.Context, r *pb.AddOnTrackRequest) (*pb.AddOnTrackResponse, error) {
	res, err := s.controller.AddContainerNumbersOnTrack(ctx, s.converter.convertAddOnTrack(r, "OTHER"))
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
	res, err := s.controller.AddBillNumbersOnTrack(ctx, s.converter.convertAddOnTrack(r, "RU"))
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
func (s *Grpc) deleteFromTracking(ctx context.Context, r *pb.DeleteFromTrackRequest, isContainer bool) (*emptypb.Empty, error) {
	if err := s.controller.DeleteFromTracking(ctx, r.GetUserId(), isContainer, r.GetNumber()...); err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
		case *NumberDoesntBelongThisUserError:
			return &emptypb.Empty{}, status.Error(codes.PermissionDenied, err.Error())
		default:
			go func() {
				for _, v := range r.GetNumber() {
					s.logger.ExceptionLog(fmt.Sprintf(`delete Number: %s for user-pb %d from tracking err: %s`, v, r.GetUserId(), err.Error()))
				}
			}()
			return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
		}
	}
	return &emptypb.Empty{}, nil
}
func (s *Grpc) DeleteContainersFromTrack(ctx context.Context, r *pb.DeleteFromTrackRequest) (*emptypb.Empty, error) {
	return s.deleteFromTracking(ctx, r, true)
}
func (s *Grpc) DeleteBillNosFromTrack(ctx context.Context, r *pb.DeleteFromTrackRequest) (*emptypb.Empty, error) {
	return s.deleteFromTracking(ctx, r, false)
}
func (s *Grpc) GetInfoAboutTrack(ctx context.Context, r *pb.GetInfoAboutTrackRequest) (*pb.GetInfoAboutTrackResponse, error) {
	resp, err := s.controller.GetInfoAboutTracking(ctx, r.GetNumber(), r.GetUserId())
	if err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return &pb.GetInfoAboutTrackResponse{
				Number:      resp.number,
				Emails:      []string{},
				NextRunTime: 0,
			}, status.Error(codes.NotFound, "task with this id was not found")
		case *NumberDoesntBelongThisUserError:
			return &pb.GetInfoAboutTrackResponse{}, status.Error(codes.PermissionDenied, err.Error())
		default:
			go s.logger.ExceptionLog(fmt.Sprintf(`get info about tracking err: %s`, err.Error()))
			return &pb.GetInfoAboutTrackResponse{
				Number:      resp.number,
				Emails:      []string{},
				NextRunTime: 0,
			}, status.Error(codes.Internal, err.Error())
		}
	}
	return &pb.GetInfoAboutTrackResponse{
		Number:              resp.number,
		Emails:              s.converter.convertInterfaceArrayToStringArray(resp.emails),
		NextRunTime:         resp.nextRunTime.Unix(),
		EmailMessageSubject: resp.emailMessageSubject,
		Time:                resp.time,
	}, nil
}
func (s *Grpc) GetTimeZone(context.Context, *emptypb.Empty) (*pb.GetTimeZoneResponse, error) {
	t := time.Now()
	zone, _ := t.Zone()
	return &pb.GetTimeZoneResponse{TimeZone: fmt.Sprintf(`UTC%s`, zone)}, nil
}
func (s *Grpc) Update(ctx context.Context, r *pb.UpdateTaskRequest) (*emptypb.Empty, error) {
	if err := s.controller.Update(ctx, s.converter.convertUpdateTaskRequest(r, "RU"), r.GetIsContainers()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
