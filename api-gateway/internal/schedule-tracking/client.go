package schedule_tracking

import (
	"context"
	"errors"
	"fmc-gateway/pkg/logging"
	pb "fmc-gateway/pkg/schedule-tracking-pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IClient interface {
	AddContainersOnTrack(ctx context.Context, userId int, req *AddOnTrackRequest) (*AddOnTrackResponse, error)
	AddBillNosOnTrack(ctx context.Context, userId int, req *AddOnTrackRequest) (*AddOnTrackResponse, error)
	Update(ctx context.Context, userId int, isContainer bool, req *AddOnTrackRequest) error
	DeleteFromTracking(ctx context.Context, isContainer bool, userId int64, req *DeleteFromTrackRequest) error
	GetInfoAboutTrack(ctx context.Context, r GetInfoAboutTrackRequest) (GetInfoAboutTrackResponse, error)
	GetTimeZone(ctx context.Context) (*TimeZoneResponse, error)
}
type Client struct {
	conn       *grpc.ClientConn
	cli        pb.ScheduleTrackingClient
	archiveCli pb.ArchiveClient
	logger     logging.ILogger
	converter  converter
}

func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{conn: conn, cli: pb.NewScheduleTrackingClient(conn), logger: logger, converter: converter{}}
}

func (c *Client) AddContainersOnTrack(ctx context.Context, userId int, req *AddOnTrackRequest) (*AddOnTrackResponse, error) {
	response, err := c.cli.AddContainersOnTrack(ctx, c.converter.addOnTrackRequestConvert(userId, req))
	if err != nil {
		statusOfRequest := status.Convert(err)
		switch statusOfRequest.Code() {
		case codes.PermissionDenied:
			return &AddOnTrackResponse{}, errors.New("number does not belong to this user or cannot find job by your params")
		case codes.InvalidArgument:
			return &AddOnTrackResponse{}, errors.New(err.Error())
		}
		return nil, err
	}
	return c.converter.addOnTrackResponseConvert(response), nil
}

func (c *Client) AddBillNosOnTrack(ctx context.Context, userId int, req *AddOnTrackRequest) (*AddOnTrackResponse, error) {
	response, err := c.cli.AddBillNosOnTrack(ctx, c.converter.addOnTrackRequestConvert(userId, req))
	if err != nil {
		statusOfRequest := status.Convert(err)
		switch statusOfRequest.Code() {
		case codes.PermissionDenied:
			return &AddOnTrackResponse{}, errors.New("number does not belong to this user or cannot find job by your params")
		case codes.InvalidArgument:
			return &AddOnTrackResponse{}, errors.New(err.Error())
		}
		return nil, err
	}
	return c.converter.addOnTrackResponseConvert(response), nil
}

//func (c *Client) UpdateTrackingTime(ctx context.Context, req UpdateTrackingTimeRequest) ([]BaseAddOnTrackResponse, error) {
//	result, err := c.cli.UpdateTrackingTime(ctx, c.converter.updateTrackingTimeRequestConvert(req))
//	if err != nil {
//		var numbers []BaseAddOnTrackResponse
//		statusOfRequest := status.Convert(err)
//		switch statusOfRequest.Code() {
//		case codes.NotFound:
//			return numbers, errors.New("cannot lookup job with this id")
//		case codes.PermissionDenied:
//			return numbers, errors.New("number does not belong to this user or cannot find job by your params")
//		case codes.InvalidArgument:
//			return numbers, errors.New(err.Error())
//		}
//		go func() {
//			for _, v := range req.Numbers {
//				c.logger.ExceptionLog(fmt.Sprintf(`update tracking time on track with number: %s err: %s`, v, err.Error()))
//			}
//		}()
//		return numbers, err
//	}
//	return c.converter.baseAddOnTrackResponseConver(result.GetResponse()), nil
//}
//
//func (c *Client) AddEmailsOnTracking(ctx context.Context, req AddEmailRequest) error {
//	_, err := c.cli.AddEmailsOnTracking(ctx, c.converter.AddEmailRequestConvert(req))
//	if err != nil {
//		statusCode := status.Convert(err).Code()
//		switch statusCode {
//		case codes.NotFound:
//			return errors.New("cannot find job with this id")
//		case codes.PermissionDenied:
//			return errors.New("number does not belong to this user or cannot find job by your params")
//		case codes.InvalidArgument:
//			return errors.New(err.Error())
//		default:
//			return err
//		}
//	}
//	return err
//}
//func (c *Client) DeleteEmailFromTrack(ctx context.Context, r DeleteEmailFromTrackRequest) error {
//	_, err := c.cli.DeleteEmailFromTrack(ctx, &__.DeleteEmailFromTrackRequest{
//		Number: r.Number,
//		Email:  r.Email,
//		UserId: r.userId,
//	})
//	if err != nil {
//		statusCode := status.Convert(err).Code()
//		switch statusCode {
//		case codes.NotFound:
//			return errors.New("cannot find job or email with this params")
//		case codes.PermissionDenied:
//			return errors.New("number does not belong to this user or cannot find job by your params")
//		case codes.InvalidArgument:
//			return errors.New(err.Error())
//		default:
//			return err
//		}
//	}
//	return err
//}

func (c *Client) DeleteFromTracking(ctx context.Context, isContainer bool, userId int64, req *DeleteFromTrackRequest) error {
	if isContainer {
		_, err := c.cli.DeleteFromTracking(ctx, &pb.DeleteFromTrackRequest{
			UserId: userId,
			Number: req.Numbers,
		})
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.NotFound:
			return errors.New("cannot find job with this id")
		case codes.PermissionDenied:
			return errors.New("number does not belong to this user or cannot find job by your params")
		case codes.InvalidArgument:
			return errors.New(err.Error())
		}
		return err
	} else {
		_, err := c.cli.DeleteFromTracking(ctx, &pb.DeleteFromTrackRequest{
			UserId: userId,
			Number: req.Numbers,
		})
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.NotFound:
			return errors.New("cannot find job with this id")
		case codes.PermissionDenied:
			return errors.New("number does not belong to this user or cannot find job by your params")
		case codes.InvalidArgument:
			return errors.New(err.Error())
		}
		return err
	}
}

func (c *Client) GetInfoAboutTrack(ctx context.Context, r GetInfoAboutTrackRequest) (GetInfoAboutTrackResponse, error) {
	resp, err := c.cli.GetInfoAboutTrack(ctx, &pb.GetInfoAboutTrackRequest{Number: r.Number, UserId: r.userId})
	if err != nil {
		var s GetInfoAboutTrackResponse
		code := status.Convert(err).Code()
		switch code {
		case codes.NotFound:
			return s, errors.New("cannot find job with this id")
		case codes.PermissionDenied:
			return s, errors.New("number does not belong to this user or cannot find job by your params")
		case codes.InvalidArgument:
			return s, errors.New(err.Error())
		default:
			return s, err
		}
	}
	return GetInfoAboutTrackResponse{
		Number:      resp.GetNumber(),
		IsContainer: resp.GetIsContainer(),
		IsOnTrack:   resp.GetIsOnTrack(),
		ScheduleTrackingInfo: &ScheduleTrackingInfo{
			Time:    resp.GetScheduleTrackingInfo().GetTime(),
			Subject: resp.GetScheduleTrackingInfo().GetSubject(),
			Emails:  resp.GetScheduleTrackingInfo().GetEmails(),
		},
	}, nil
}
func (c *Client) GetTimeZone(ctx context.Context) (*TimeZoneResponse, error) {
	timeZone, err := c.cli.GetTimeZone(ctx, &emptypb.Empty{})
	if err != nil {
		return &TimeZoneResponse{}, err
	}
	return &TimeZoneResponse{TimeZone: timeZone.GetTimeZone()}, nil
}
func (c *Client) Update(ctx context.Context, userId int, isContainer bool, req *AddOnTrackRequest) error {
	if _, err := c.cli.Update(ctx, &pb.UpdateTaskRequest{
		Req:          req.ToGrpc(userId),
		IsContainers: isContainer,
	}); err != nil {
		return err
	}
	return nil
}

type converter struct {
}

func (c *converter) addOnTrackRequestConvert(userId int, r *AddOnTrackRequest) *pb.AddOnTrackRequest {
	return &pb.AddOnTrackRequest{
		UserId:              int64(userId),
		Numbers:             r.Numbers,
		Emails:              r.Emails,
		EmailMessageSubject: r.EmailMessageSubject,
	}
}
func (c *converter) baseAddOnTrackResponseConver(r []*pb.BaseAddOnTrackResponse) []BaseAddOnTrackResponse {
	var result []BaseAddOnTrackResponse
	for _, v := range r {
		result = append(result, BaseAddOnTrackResponse{
			Success:     v.GetSuccess(),
			Number:      v.GetNumber(),
			NextRunTime: v.GetNextRunTime(),
		})
	}
	return result
}
func (c *converter) addOnTrackResponseConvert(r *pb.AddOnTrackResponse) *AddOnTrackResponse {
	return &AddOnTrackResponse{
		Result:         c.baseAddOnTrackResponseConver(r.GetBaseResponse()),
		AlreadyOnTrack: r.GetAlreadyOnTrack(),
	}
}
