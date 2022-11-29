package user

import (
	"context"
	"fmc-gateway/pkg/logging"
	"fmt"
	pb "github.com/frozosea/fmc-proto/user"
	"google.golang.org/grpc"
)

type converter struct{}

func (c *converter) ConvertToGrpcContainer(r []string) []*pb.Container {
	var containers []*pb.Container
	for _, item := range r {
		containers = append(containers, &pb.Container{Number: item})
	}
	return containers
}
func (c *converter) scheduleTrackingInfoObjectFromGrpc(r *pb.ScheduleTrackingObject) *ScheduleTrackingInfoObject {
	return &ScheduleTrackingInfoObject{
		Time:    r.GetTime(),
		Emails:  r.GetEmails(),
		Subject: r.GetSubject(),
	}
}
func (c *converter) containerFromGrpc(r []*pb.ContainerResponse) []*Container {
	var containers []*Container
	for _, item := range r {
		containers = append(containers, &Container{
			Number:               item.GetNumber(),
			IsOnTrack:            item.GetIsOnTrack(),
			IsContainer:          item.GetIsContainer(),
			ScheduleTrackingInfo: c.scheduleTrackingInfoObjectFromGrpc(item.GetScheduleTrackingObject()),
		})
	}
	return containers
}

type Client struct {
	conn        *grpc.ClientConn
	userCli     pb.UserClient
	feedbackCli pb.UserFeedbackClient
	logger      logging.ILogger
	converter   converter
}

func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{conn: conn, userCli: pb.NewUserClient(conn), feedbackCli: pb.NewUserFeedbackClient(conn), logger: logger, converter: converter{}}
}

func (c *Client) AddContainerToAccount(ctx context.Context, userId int64, r *AddContainers) error {
	_, err := c.userCli.AddContainerToAccount(ctx, &pb.AddContainerToAccountRequest{
		Container: c.converter.ConvertToGrpcContainer(r.Numbers),
		UserId:    userId,
	})
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add container to account: %d failed: %s`, userId, err.Error()))
		return err
	}
	return nil
}
func (c *Client) DeleteContainersFromAccount(ctx context.Context, userId int64, r *DeleteNumbers) error {
	_, err := c.userCli.DeleteContainersFromAccount(ctx, &pb.DeleteContainersFromAccountRequest{
		UserId:  userId,
		Numbers: r.Numbers,
	})
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`delete containers from account: %d for numbers: %v failed: %s`, userId, r.Numbers, err.Error()))
		return err
	}
	return nil
}
func (c *Client) DeleteBillNumbersFromAccount(ctx context.Context, userId int64, r *DeleteNumbers) error {
	_, err := c.userCli.DeleteBillNumbersFromAccount(ctx, &pb.DeleteContainersFromAccountRequest{
		UserId:  userId,
		Numbers: r.Numbers,
	})
	if err != nil {
		//go c.logger.ExceptionLog(fmt.Sprintf(`delete bill numbers from account: %d for numbers: %v failed: %s`, userId, r.numberIds, err.Error()))
		return err
	}
	return nil
}
func (c *Client) AddBillNumbersToAccount(ctx context.Context, userId int64, r *AddContainers) error {
	_, err := c.userCli.AddBillNumberToAccount(ctx, &pb.AddContainerToAccountRequest{
		Container: c.converter.ConvertToGrpcContainer(r.Numbers),
		UserId:    userId,
	})

	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`add bill numbers to account: %d failed: %s`, userId, err.Error()))
		return err
	}
	return nil
}
func (c *Client) GetAll(ctx context.Context, userId int64) (*AllContainersAndBillNumbers, error) {
	result, err := c.userCli.GetAll(ctx, &pb.GetAllContainersFromAccountRequest{UserId: userId})
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`get all containers and bills for user: %d failed: %s`, userId, err.Error()))
		return &AllContainersAndBillNumbers{}, err
	}
	result.GetBillNumbers()
	return &AllContainersAndBillNumbers{
		Containers:  c.converter.containerFromGrpc(result.GetContainers()),
		BillNumbers: c.converter.containerFromGrpc(result.GetBillNumbers()),
	}, nil
}
func (c *Client) AddFeedback(ctx context.Context, email, message string) error {
	_, err := c.feedbackCli.AddFeedback(ctx, &pb.AddFeedbackRequest{
		Email:   email,
		Message: message,
	})
	return err
}
