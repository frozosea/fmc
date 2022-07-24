package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"schedule-tracking/internal/logging"
	"schedule-tracking/internal/scheduler"
	pb "schedule-tracking/pkg/proto"
)

type converter struct {
}

func (c *converter) convertAddOnTrack(r *pb.AddOnTrackRequest) BaseTrackReq {
	return BaseTrackReq{
		Numbers: r.GetNumber(),
		UserId:  r.GetUserId(),
		Country: "OTHER",
		Time:    r.GetTime(),
		Emails:  r.GetEmails(),
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
func (c *converter) convertAddEmails(r *pb.AddEmailRequest) AddEmailRequest {
	return AddEmailRequest{
		numbers: r.GetNumbers(),
		emails:  r.GetEmails(),
		userId:  r.GetUserId(),
	}
}
func (c *converter) convertDeleteEmails(r *pb.DeleteEmailFromTrackRequest) DeleteEmailFromTrack {
	return DeleteEmailFromTrack{
		number: r.GetNumber(),
		email:  r.GetEmail(),
		userId: r.GetUserId(),
	}
}
func (c *converter) convertInterfaceArrayToStringArray(r []interface{}) []string {
	var outputArr []string
	for _, v := range r {
		outputArr = append(outputArr, fmt.Sprintf(`%v`, v))
	}
	return outputArr
}

type Service struct {
	controller *Controller
	logger     logging.ILogger
	converter
	pb.UnimplementedScheduleTrackingServer
}

func NewService(controller *Controller, logger logging.ILogger) *Service {
	return &Service{controller: controller, logger: logger, converter: converter{}, UnimplementedScheduleTrackingServer: pb.UnimplementedScheduleTrackingServer{}}
}
func (s *Service) AddContainersOnTrack(ctx context.Context, r *pb.AddOnTrackRequest) (*pb.AddOnTrackResponse, error) {
	res, err := s.controller.AddContainerNumbersOnTrack(ctx, TrackByContainerNoReq{s.converter.convertAddOnTrack(r)})
	if err != nil {
		go func() {
			for _, v := range res.result {
				s.logger.FatalLog(fmt.Sprintf(`add container Numbers: %s for user-pb: %d failed: %s`, v.number, r.UserId, err.Error()))
			}
		}()
		switch err.(type) {
		case *scheduler.LookupJobError:
			return s.converter.convertAddOnTrackResponse(res), status.Error(codes.NotFound, "cannot find job with this id")
		default:
			return s.converter.convertAddOnTrackResponse(res), status.Error(codes.Internal, err.Error())
		}
	}
	go func() {
		jsonRepr, reprErr := json.Marshal(res)
		if reprErr != nil {
			return
		}
		s.logger.InfoLog(fmt.Sprintf(`add container Numbers on track request: %v to user-pb: %v, with result: %v`, r.Number, r.UserId, jsonRepr))
	}()
	return s.converter.convertAddOnTrackResponse(res), nil
}

func (s *Service) AddBillNosOnTrack(ctx context.Context, r *pb.AddOnTrackRequest) (*pb.AddOnTrackResponse, error) {
	res, err := s.controller.AddBillNumbersOnTrack(ctx, TrackByBillNoReq{s.converter.convertAddOnTrack(r)})
	if err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return s.converter.convertAddOnTrackResponse(res), status.Error(codes.NotFound, "cannot find job with this id")
		default:
			return s.converter.convertAddOnTrackResponse(res), status.Error(codes.Internal, err.Error())
		}
	}
	go func() {
		jsonRepr, err := json.Marshal(res)
		if err != nil {
			return
		}
		s.logger.InfoLog(fmt.Sprintf(`add container Numbers on track request: %v to user-pb: %v, with result: %v`, r.Number, r.UserId, jsonRepr))
	}()
	return s.converter.convertAddOnTrackResponse(res), nil
}
func (s *Service) UpdateTrackingTime(ctx context.Context, r *pb.UpdateTrackingTimeRequest) (*pb.RepeatedBaseAddOnTrackResponse, error) {
	resp, err := s.controller.UpdateTrackingTime(ctx, r.GetNumbers(), r.GetTime(), r.GetUserId())
	if err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return &pb.RepeatedBaseAddOnTrackResponse{}, status.Error(codes.NotFound, "cannot find job with this id ")
		case *NumberDoesntBelongThisUserError:
			return &pb.RepeatedBaseAddOnTrackResponse{}, status.Error(codes.PermissionDenied, err.Error())
		default:
			return &pb.RepeatedBaseAddOnTrackResponse{}, status.Error(codes.Internal, err.Error())
		}
	}
	go func() {
		for _, v := range resp {
			s.logger.InfoLog(fmt.Sprintf(`task with id: %s new Time: %s`, v.number, r.Time))
		}
	}()
	return &pb.RepeatedBaseAddOnTrackResponse{Response: s.convertBaseAddOnTrackResponse(resp)}, nil

}
func (s *Service) AddEmailsOnTracking(ctx context.Context, r *pb.AddEmailRequest) (*emptypb.Empty, error) {
	if err := s.controller.AddEmailToTracking(ctx, s.converter.convertAddEmails(r)); err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return &emptypb.Empty{}, status.Error(codes.NotFound, "cannot find job with this id")
		case *NumberDoesntBelongThisUserError:
			return &emptypb.Empty{}, status.Error(codes.PermissionDenied, err.Error())
		default:
			go func() {
				s.logger.ExceptionLog(fmt.Sprintf(`add Emails: %v for Numbers: %v err: %s`, r.GetEmails(), r.GetNumbers(), err.Error()))
			}()
			return &emptypb.Empty{}, err
		}
	}
	return &emptypb.Empty{}, nil
}
func (s *Service) DeleteEmailFromTrack(ctx context.Context, r *pb.DeleteEmailFromTrackRequest) (*emptypb.Empty, error) {
	if err := s.controller.DeleteEmailFromTrack(ctx, s.converter.convertDeleteEmails(r)); err != nil {
		switch err.(type) {
		case *scheduler.LookupJobError:
			return &emptypb.Empty{}, status.Error(codes.NotFound, "cannot find job with this id")
		case *CannotFindEmailError:
			return &emptypb.Empty{}, status.Error(codes.NotFound, "cannot find email in job args")
		case *NumberDoesntBelongThisUserError:
			return &emptypb.Empty{}, status.Error(codes.PermissionDenied, err.Error())
		default:
			go s.logger.ExceptionLog(fmt.Sprintf(`delete email: %s for Number: %s err: %s`, r.GetEmail(), r.GetNumber(), err.Error()))
			return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())

		}
	}
	return &emptypb.Empty{}, nil

}
func (s *Service) deleteFromTracking(ctx context.Context, r *pb.DeleteFromTrackRequest, isContainer bool) (*emptypb.Empty, error) {
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
func (s *Service) DeleteContainersFromTrack(ctx context.Context, r *pb.DeleteFromTrackRequest) (*emptypb.Empty, error) {
	return s.deleteFromTracking(ctx, r, true)
}
func (s *Service) DeleteBillNosFromTrack(ctx context.Context, r *pb.DeleteFromTrackRequest) (*emptypb.Empty, error) {
	return s.deleteFromTracking(ctx, r, false)
}
func (s *Service) GetInfoAboutTrack(ctx context.Context, r *pb.GetInfoAboutTrackRequest) (*pb.GetInfoAboutTrackResponse, error) {
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
		Number:      resp.number,
		Emails:      s.converter.convertInterfaceArrayToStringArray(resp.emails),
		NextRunTime: resp.nextRunTime.Unix(),
	}, nil
}
