package schedule_tracking

import (
	"context"
	"fmc-with-git/internal/logging"
	pb "fmc-with-git/internal/user-pb"
	"fmt"
	"google.golang.org/grpc"
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
		go func() {
			for _, v := range req.Numbers {
				c.logger.ExceptionLog(fmt.Sprintf(`add container on track with number: %s err: %s`, v, err.Error()))
			}
		}()
		return nil, err
	}
	return c.converter.addOnTrackResponseConvert(response), nil
}

func (c *Client) AddBillNosOnTrack(ctx context.Context, userId int, req *AddOnTrackRequest) (*AddOnTrackResponse, error) {
	response, err := c.cli.AddBillNosOnTrack(ctx, c.converter.addOnTrackRequestConvert(userId, req))
	if err != nil {
		go func() {
			for _, v := range req.Numbers {
				c.logger.ExceptionLog(fmt.Sprintf(`add bill number on track with number: %s err: %s`, v, err.Error()))
			}
		}()
		return nil, err
	}
	return c.converter.addOnTrackResponseConvert(response), nil
}
func (c *Client) UpdateTrackingTime(ctx context.Context, req UpdateTrackingTimeRequest) ([]BaseAddOnTrackResponse, error) {
	result, err := c.cli.UpdateTrackingTime(ctx, c.converter.updateTrackingTimeRequestConvert(req))
	if err != nil {
		go func() {
			for _, v := range req.Numbers {
				c.logger.ExceptionLog(fmt.Sprintf(`update tracking time on track with number: %s err: %s`, v, err.Error()))
			}
		}()
		var numbers []BaseAddOnTrackResponse
		return numbers, err
	}
	return c.converter.baseAddOnTrackResponseConver(result.GetResponse()), nil
}

func (c *Client) AddEmailsOnTracking(ctx context.Context, req AddEmailRequest) error {
	_, err := c.cli.AddEmailsOnTracking(ctx, c.converter.AddEmailRequestConvert(req))
	if err != nil {
		go func() {
			for index, v := range req.Emails {
				c.logger.ExceptionLog(fmt.Sprintf(`add email on track with number: %s email: %s err: %s`, v, req.Emails[index], err.Error()))
			}
		}()
	}
	return err
}
func (c *Client) DeleteEmailFromTrack(ctx context.Context, r DeleteEmailFromTrackRequest) error {
	_, err := c.cli.DeleteEmailFromTrack(ctx, &pb.DeleteEmailFromTrackRequest{
		Number: r.Number,
		Email:  r.Email,
	})
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`delete email from track with number: %s email: %s err: %s`, r.Number, r.Email, err.Error()))
	}
	return err
}

func (c *Client) DeleteFromTracking(ctx context.Context, isContainer bool, userId int64, req DeleteFromTrackRequest) error {
	if isContainer {
		_, err := c.cli.DeleteContainersFromTrack(ctx, &pb.DeleteFromTrackRequest{
			UserId: userId,
			Number: req.Numbers,
		})
		go func() {
			for _, v := range req.Numbers {
				c.logger.ExceptionLog(fmt.Sprintf(`delete from track with number: %s for user: %d err: %s`, v, userId, err.Error()))
			}
		}()
		return err
	} else {
		_, err := c.cli.DeleteBillNosFromTrack(ctx, &pb.DeleteFromTrackRequest{
			UserId: userId,
			Number: req.Numbers,
		})
		return err
	}
}

func (c *Client) GetInfoAboutTrack(ctx context.Context, r GetInfoAboutTrackRequest) (GetInfoAboutTrackResponse, error) {
	resp, err := c.cli.GetInfoAboutTrack(ctx, &pb.GetInfoAboutTrackRequest{Number: r.Number})
	if err != nil {
		var s GetInfoAboutTrackResponse
		go c.logger.ExceptionLog(fmt.Sprintf(`get info about track for number: %s err: %s`, r.Number, err.Error()))
		return s, err
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
	}
}
func (c *converter) AddEmailRequestConvert(r AddEmailRequest) *pb.AddEmailRequest {
	return &pb.AddEmailRequest{
		Numbers: r.Numbers,
		Emails:  r.Emails,
	}
}
