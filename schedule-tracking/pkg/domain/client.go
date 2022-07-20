package domain

import (
	"context"
	"fmt"
	"schedule-tracking/internal/logging"
	user_pb "schedule-tracking/internal/user-pb"
)

type Client struct {
	cli    user_pb.ScheduleTrackingClient
	logger logging.ILogger
}

func NewClient(cli user_pb.ScheduleTrackingClient, logger logging.ILogger) *Client {
	return &Client{cli: cli, logger: logger}
}

func (c *Client) MarkBillNoOnTrack(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillNoOnTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark bill with Number %s no add on track for user %d failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *Client) MarkContainerOnTrack(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillNoOnTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark container with Number %s no add on track for user %d failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *Client) MarkContainerWasArrived(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkContainerWasArrived(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark container with Number %s no was arrived for user %d failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *Client) MarkBillNoWasArrived(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillNoWasArrived(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark bill with Number %s no was arrived for user %d failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *Client) MarkContainerWasRemovedFromTrack(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkContainerWasRemovedFromTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark container with Number %s for user %d no was removed from tracking failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *Client) MarkBillNoWasRemovedFromTrack(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillNoWasRemovedFromTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark container with Number %s for user %d no was removed from tracking failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
