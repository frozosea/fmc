package user

import (
	"context"
	"fmc-gateway/internal/logging"
	pb "fmc-gateway/internal/user-pb"
	"fmt"
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
func (c *converter) containerFromGrpc(r []*pb.ContainerResponse) []*Container {
	var containers []*Container
	for _, item := range r {
		containers = append(containers, &Container{
			Id:        item.GetId(),
			Number:    item.GetNumber(),
			IsOnTrack: item.GetIsOnTrack(),
		})
	}
	return containers
}

type Client struct {
	conn      *grpc.ClientConn
	cli       pb.UserClient
	logger    logging.ILogger
	converter converter
}

func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{conn: conn, cli: pb.NewUserClient(conn), logger: logger, converter: converter{}}
}

func (c *Client) AddContainerToAccount(ctx context.Context, userId int64, r *AddContainers) error {
	_, err := c.cli.AddContainerToAccount(ctx, &pb.AddContainerToAccountRequest{
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
	_, err := c.cli.DeleteContainersFromAccount(ctx, &pb.DeleteContainersFromAccountRequest{
		UserId:    userId,
		NumberIds: r.Numbers,
	})
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`delete containers from account: %d for numbers: %v failed: %s`, userId, r.Numbers, err.Error()))
		return err
	}
	return nil
}
func (c *Client) DeleteBillNumbersFromAccount(ctx context.Context, userId int64, r *DeleteNumbers) error {
	_, err := c.cli.DeleteBillNumbersFromAccount(ctx, &pb.DeleteContainersFromAccountRequest{
		UserId:    userId,
		NumberIds: r.Numbers,
	})
	if err != nil {
		//go c.logger.ExceptionLog(fmt.Sprintf(`delete bill numbers from account: %d for numbers: %v failed: %s`, userId, r.numberIds, err.Error()))
		return err
	}
	return nil
}
func (c *Client) AddBillNumbersToAccount(ctx context.Context, userId int64, r *AddContainers) error {
	_, err := c.cli.AddBillNumberToAccount(ctx, &pb.AddContainerToAccountRequest{
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
	result, err := c.cli.GetAll(ctx, &pb.GetAllContainersFromAccountRequest{UserId: userId})
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
