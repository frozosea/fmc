package schedule_tracking

import (
	"context"
	"errors"
	"fmc-gateway/internal/logging"
	pb "fmc-gateway/internal/schedule-tracking-pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	conn      *grpc.ClientConn
	cli       pb.ScheduleTrackingClient
	logger    logging.ILogger
	converter converter
}

func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{conn: conn, cli: pb.NewScheduleTrackingClient(conn), logger: logger, converter: converter{}}
}

func (c *Client) AddContainersOnTrack(ctx context.Context, userId int, req *AddOnTrackRequest) (*AddOnTrackResponse, error) {
	response, err := c.cli.AddContainersOnTrack(ctx, c.converter.addOnTrackRequestConvert(userId, req))
	if err != nil {
		//go func() {
		//	for _, v := range req.Numbers {
		//		c.logger.ExceptionLog(fmt.Sprintf(`add container on track with number: %s err: %s`, v, err.Error()))
		//	}
		//}()
		return nil, err
	}
	return c.converter.addOnTrackResponseConvert(response), nil
}

func (c *Client) AddBillNosOnTrack(ctx context.Context, userId int, req *AddOnTrackRequest) (*AddOnTrackResponse, error) {
	response, err := c.cli.AddBillNosOnTrack(ctx, c.converter.addOnTrackRequestConvert(userId, req))
	if err != nil {
		//go func() {
		//	for _, v := range req.Numbers {
		//		c.logger.ExceptionLog(fmt.Sprintf(`add bill number on track with number: %s err: %s`, v, err.Error()))
		//	}
		//}()
		return nil, err
	}
	return c.converter.addOnTrackResponseConvert(response), nil
}
func (c *Client) UpdateTrackingTime(ctx context.Context, req UpdateTrackingTimeRequest) ([]BaseAddOnTrackResponse, error) {
	result, err := c.cli.UpdateTrackingTime(ctx, c.converter.updateTrackingTimeRequestConvert(req))
	if err != nil {
		var numbers []BaseAddOnTrackResponse
		statusOfRequest := status.Convert(err)
		switch statusOfRequest.Code() {
		case codes.NotFound:
			return numbers, errors.New("cannot lookup job with this id")
		case codes.PermissionDenied:
			return numbers, errors.New("number does not belong to this user or cannot find job by your params")
		}
		go func() {
			for _, v := range req.Numbers {
				c.logger.ExceptionLog(fmt.Sprintf(`update tracking time on track with number: %s err: %s`, v, err.Error()))
			}
		}()
		return numbers, err
	}
	return c.converter.baseAddOnTrackResponseConver(result.GetResponse()), nil
}

func (c *Client) AddEmailsOnTracking(ctx context.Context, req AddEmailRequest) error {
	_, err := c.cli.AddEmailsOnTracking(ctx, c.converter.AddEmailRequestConvert(req))
	if err != nil {
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.NotFound:
			return errors.New("cannot find job with this id")
		case codes.PermissionDenied:
			return errors.New("number does not belong to this user or cannot find job by your params")
		default:
			return err
		}
	}
	return err
}
func (c *Client) DeleteEmailFromTrack(ctx context.Context, r DeleteEmailFromTrackRequest) error {
	_, err := c.cli.DeleteEmailFromTrack(ctx, &pb.DeleteEmailFromTrackRequest{
		Number: r.Number,
		Email:  r.Email,
		UserId: r.UserId,
	})
	if err != nil {
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.NotFound:
			return errors.New("cannot find job or email with this params")
		case codes.PermissionDenied:
			return errors.New("number does not belong to this user or cannot find job by your params")

		default:
			return err
		}
	}
	return err
}

func (c *Client) DeleteFromTracking(ctx context.Context, isContainer bool, userId int64, req DeleteFromTrackRequest) error {
	if isContainer {
		_, err := c.cli.DeleteContainersFromTrack(ctx, &pb.DeleteFromTrackRequest{
			UserId: userId,
			Number: req.Numbers,
		})
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.NotFound:
			return errors.New("cannot find job with this id")
		case codes.PermissionDenied:
			return errors.New("number does not belong to this user or cannot find job by your params")
		}
		return err
	} else {
		_, err := c.cli.DeleteBillNosFromTrack(ctx, &pb.DeleteFromTrackRequest{
			UserId: userId,
			Number: req.Numbers,
		})
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.NotFound:
			return errors.New("cannot find job with this id")
		case codes.PermissionDenied:
			return errors.New("number does not belong to this user or cannot find job by your params")
		}
		return err
	}
}

func (c *Client) GetInfoAboutTrack(ctx context.Context, r GetInfoAboutTrackRequest) (GetInfoAboutTrackResponse, error) {
	resp, err := c.cli.GetInfoAboutTrack(ctx, &pb.GetInfoAboutTrackRequest{Number: r.Number, UserId: r.UserId})
	if err != nil {
		var s GetInfoAboutTrackResponse
		code := status.Convert(err).Code()
		switch code {
		case codes.NotFound:
			return s, errors.New("cannot find job with this id")
		case codes.PermissionDenied:
			return s, errors.New("number does not belong to this user or cannot find job by your params")
		default:
			return s, err
		}
	}
	return GetInfoAboutTrackResponse{
		Number:      resp.GetNumber(),
		Emails:      resp.GetEmails(),
		NextRunTime: resp.GetNextRunTime(),
	}, nil
}

type converter struct {
}

func (c *converter) addOnTrackRequestConvert(userId int, r *AddOnTrackRequest) *pb.AddOnTrackRequest {
	return &pb.AddOnTrackRequest{
		UserId: int64(userId),
		Number: r.Numbers,
		Emails: r.Emails,
		Time:   r.Time,
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
func (c *converter) updateTrackingTimeRequestConvert(r UpdateTrackingTimeRequest) *pb.UpdateTrackingTimeRequest {
	return &pb.UpdateTrackingTimeRequest{
		Numbers: r.Numbers,
		Time:    r.NewTime,
		UserId:  r.UserId,
	}
}
func (c *converter) AddEmailRequestConvert(r AddEmailRequest) *pb.AddEmailRequest {
	return &pb.AddEmailRequest{
		Numbers: r.Numbers,
		Emails:  r.Emails,
		UserId:  r.UserId,
	}
}
