package tracking

import (
	"context"
	"fmt"
	pb "github.com/frozosea/fmc-pb/tracking"
	"google.golang.org/grpc"
	"schedule-tracking/pkg/logging"
	"time"
)

type Client struct {
	conn              *grpc.ClientConn
	billNoClient      pb.TrackingByBillNumberClient
	containerNoClient pb.TrackingByContainerNumberClient
	logger            logging.ILogger
	Converter
}

func (c *Client) TrackByBillNumber(ctx context.Context, track *Track) (BillNumberResponse, error) {
	request := pb.Request{
		Number:  track.Number,
		Scac:    track.Scac,
		Country: pb.Country(pb.Country_value[track.Country]),
	}
	response, err := c.billNoClient.TrackByBillNumber(ctx, &request)
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`trackingByBillNumber error: %s`, err.Error()))
		return BillNumberResponse{}, err
	}
	return c.convertGrpcBlNoResponse(response), nil
}

func (c *Client) TrackByContainerNumber(ctx context.Context, track Track) (ContainerNumberResponse, error) {
	request := pb.Request{
		Number:  track.Number,
		Scac:    track.Scac,
		Country: pb.Country(pb.Country_value[track.Country]),
	}
	response, err := c.containerNoClient.TrackByContainerNumber(ctx, &request)
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`trackingByContainerNumber error: %s`, err.Error()))
		return ContainerNumberResponse{}, err
	}
	return c.convertGrpcContainerNoResponse(response), nil
}
func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{
		conn:              conn,
		billNoClient:      pb.NewTrackingByBillNumberClient(conn),
		containerNoClient: pb.NewTrackingByContainerNumberClient(conn),
		logger:            logger,
	}
}

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) convertGrpcInfoAboutMoving(resp []*pb.InfoAboutMoving) []BaseInfoAboutMoving {
	var infoAboutMoving []BaseInfoAboutMoving
	for _, v := range resp {
		infoAboutMoving = append(infoAboutMoving, BaseInfoAboutMoving{Time: time.UnixMilli(v.GetTime()), Location: v.GetLocation(), OperationName: v.GetOperationName(), Vessel: v.GetVessel()})
	}
	return infoAboutMoving
}
func (c *Converter) convertGrpcBlNoResponse(response *pb.TrackingByBillNumberResponse) BillNumberResponse {
	return BillNumberResponse{
		BillNo:           response.GetBillNo(),
		Scac:             response.GetScac(),
		InfoAboutMoving:  c.convertGrpcInfoAboutMoving(response.InfoAboutMoving),
		EtaFinalDelivery: time.UnixMilli(response.GetEtaFinalDelivery()),
	}
}
func (c *Converter) convertGrpcContainerNoResponse(response *pb.TrackingByContainerNumberResponse) ContainerNumberResponse {
	return ContainerNumberResponse{
		Container:       response.GetContainer(),
		ContainerSize:   response.GetContainerSize(),
		Scac:            response.GetScac(),
		InfoAboutMoving: c.convertGrpcInfoAboutMoving(response.GetInfoAboutMoving()),
	}
}
