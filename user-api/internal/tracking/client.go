package tracking

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
	"user-api/internal/logging"
)

type Client struct {
	conn              *grpc.ClientConn
	billNoClient      trackingByBillNumberClient
	containerNoClient trackingByContainerNumberClient
	logger            logging.ILogger
	converter
}

func (c *Client) TrackByBillNumber(ctx context.Context, track *Track) (BillNumberResponse, error) {
	request := Request{
		Number:  track.Number,
		Scac:    Scac(Scac_value[track.Scac]),
		Country: Country(Country_value[track.Country]),
	}
	response, err := c.billNoClient.TrackByBillNumber(ctx, &request)
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`trackingByBillNumber error: %s`, err.Error()))
		return c.convertGrpcBlNoResponse(response), err
	}
	return c.convertGrpcBlNoResponse(response), nil
}

func (c *Client) TrackByContainerNumber(ctx context.Context, track Track) (ContainerNumberResponse, error) {
	request := Request{
		Number:  track.Number,
		Scac:    Scac(Scac_value[track.Scac]),
		Country: Country(Country_value[track.Country]),
	}
	response, err := c.containerNoClient.TrackByContainerNumber(ctx, &request)
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`trackingByContainerNumber error: %s`, err.Error()))
		return c.convertGrpcContainerNoResponse(response), err
	}
	return c.convertGrpcContainerNoResponse(response), nil
}
func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{
		conn:              conn,
		billNoClient:      trackingByBillNumberClient{cc: conn},
		containerNoClient: trackingByContainerNumberClient{cc: conn},
		logger:            logger,
	}
}

type converter struct{}

//tracking written on node js so nodejs timestamp is +3 zero on end of timestamp
func (c *converter) convertNodeTimeStampToTime(t int64) time.Time {
	return time.Unix(t/1000, 0)
}
func (c *converter) convertGrpcInfoAboutMoving(resp []*InfoAboutMoving) []BaseInfoAboutMoving {
	var infoAboutMoving []BaseInfoAboutMoving
	for _, v := range resp {
		infoAboutMoving = append(infoAboutMoving, BaseInfoAboutMoving{Time: c.convertNodeTimeStampToTime(v.GetTime()), Location: v.GetLocation(), OperationName: v.GetOperationName(), Vessel: v.GetVessel()})
	}
	return infoAboutMoving
}
func (c *converter) convertGrpcBlNoResponse(response *TrackingByBillNumberResponse) BillNumberResponse {
	return BillNumberResponse{
		BillNo:           response.GetBillNo(),
		Scac:             response.GetScac().String(),
		InfoAboutMoving:  c.convertGrpcInfoAboutMoving(response.InfoAboutMoving),
		EtaFinalDelivery: c.convertNodeTimeStampToTime(response.GetEtaFinalDelivery()),
	}
}
func (c *converter) convertGrpcContainerNoResponse(response *TrackingByContainerNumberResponse) ContainerNumberResponse {
	return ContainerNumberResponse{
		Container:       response.GetContainer(),
		ContainerSize:   response.GetContainerSize(),
		Scac:            response.GetScac().String(),
		InfoAboutMoving: c.convertGrpcInfoAboutMoving(response.GetInfoAboutMoving()),
	}
}
