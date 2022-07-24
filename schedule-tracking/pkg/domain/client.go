package domain

import (
	"context"
	"fmt"
	"schedule-tracking/internal/logging"
	user_pb "schedule-tracking/internal/user-pb"
)

type UserClient struct {
	cli    user_pb.ScheduleTrackingClient
	logger logging.ILogger
}

func NewClient(cli user_pb.ScheduleTrackingClient, logger logging.ILogger) *UserClient {
	return &UserClient{cli: cli, logger: logger}
}

func (c *UserClient) MarkBillNoOnTrack(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillNoOnTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark bill with Number %s no add on track for user %d failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *UserClient) MarkContainerOnTrack(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillNoOnTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark container with Number %s no add on track for user %d failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *UserClient) MarkContainerWasArrived(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkContainerWasArrived(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark container with Number %s no was arrived for user %d failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *UserClient) MarkBillNoWasArrived(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillNoWasArrived(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark bill with Number %s no was arrived for user %d failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *UserClient) MarkContainerWasRemovedFromTrack(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkContainerWasRemovedFromTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark container with Number %s for user %d no was removed from tracking failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}
func (c *UserClient) MarkBillNoWasRemovedFromTrack(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillNoWasRemovedFromTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`mark container with Number %s for user %d no was removed from tracking failed: %s`, number, userId, err.Error()))
		return err
	}
	return nil
}

type AuthClient struct {
	client user_pb.AuthClient
	logger logging.ILogger
}
