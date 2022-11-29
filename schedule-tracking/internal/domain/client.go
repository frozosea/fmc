package domain

import (
	"context"
	"fmt"
	user_pb "github.com/frozosea/fmc-proto/user"
	"schedule-tracking/pkg/logging"
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
	if _, err := c.cli.MarkContainerOnTrack(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
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
func (c *UserClient) CheckNumberBelongUser(ctx context.Context, number string, userId int64, isContainer bool) bool {
	existStruct, err := c.cli.CheckNumberExists(ctx, &user_pb.CheckNumberExistsRequest{
		UserId:      userId,
		Number:      number,
		IsContainer: isContainer,
	})
	if err != nil || !existStruct.GetExists() {
		return false
	}
	return true
}
