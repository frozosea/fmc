package domain

import (
	"context"
	"fmt"
	user_pb "github.com/frozosea/fmc-pb/v2/user"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/util"
)

type IUserClient interface {
	MarkBillNoOnTrack(ctx context.Context, userId int64, number string) error
	MarkContainerOnTrack(ctx context.Context, userId int64, number string) error
	MarkContainerWasArrived(ctx context.Context, userId int64, number string) error
	MarkBillNoWasArrived(ctx context.Context, userId int64, number string) error
	MarkContainerWasRemovedFromTrack(ctx context.Context, userId int64, number string) error
	MarkBillNoWasRemovedFromTrack(ctx context.Context, userId int64, number string) error
	CheckNumberBelongUser(ctx context.Context, number string, userId int64, isContainer bool) bool
	MarkContainerWasNotArrived(ctx context.Context, userId int64, number string) error
	MarkBillWasNotArrived(ctx context.Context, userId int64, number string) error
	SubFromBalance(ctx context.Context, userId int64, number string) error
	CheckAccessToPaidTracking(ctx context.Context, userId int64) (bool, error)
}

type UserClient struct {
	cli        user_pb.ScheduleTrackingClient
	balanceCli user_pb.BalanceClient
	token      util.ITokenManager
	logger     logging.ILogger
}

func NewClient(cli user_pb.ScheduleTrackingClient, balanceCli user_pb.BalanceClient, logger logging.ILogger) *UserClient {
	return &UserClient{cli: cli, balanceCli: balanceCli, logger: logger}
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
func (c *UserClient) MarkContainerWasNotArrived(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkContainerIsNotArrived(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		return err
	}
	return nil
}
func (c *UserClient) MarkBillWasNotArrived(ctx context.Context, userId int64, number string) error {
	if _, err := c.cli.MarkBillIsNotArrived(ctx, &user_pb.AddMarkOnTrackingRequest{UserId: userId, Number: number}); err != nil {
		return err
	}
	return nil
}
func (c *UserClient) SubFromBalance(ctx context.Context, userId int64, number string) error {
	if _, err := c.balanceCli.SubOneDayTrackingPriceFromBalance(ctx, &user_pb.SubBalanceServiceRequest{
		UserId: userId,
		Number: number,
	}); err != nil {
		return err
	}
	return nil
}
func (c *UserClient) CheckAccessToPaidTracking(ctx context.Context, userId int64) (bool, error) {
	r, err := c.balanceCli.CheckAccessToPaidTracking(ctx, &user_pb.CheckAccessToPaidTrackingRequest{UserId: userId})
	if err != nil {
		return false, err
	}
	return r.HasAccess, nil
}
