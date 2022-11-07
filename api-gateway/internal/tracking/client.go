package tracking

import (
	"context"
	"fmc-gateway/pkg/logging"
	"fmt"
	"google.golang.org/grpc"
)

type Converter struct{}

func (c *Converter) convertGrpcInfoAboutMoving(resp []*InfoAboutMoving) []BaseInfoAboutMoving {
	var infoAboutMoving []BaseInfoAboutMoving
	for _, i := range resp {
		infoAboutMoving = append(infoAboutMoving, BaseInfoAboutMoving{Time: i.GetTime(), Location: i.GetLocation(), OperationName: i.GetOperationName(), Vessel: i.GetVessel()})
	}
	return infoAboutMoving
}
func (c *Converter) ConvertGrpcBlNoResponse(response *TrackingByBillNumberResponse) *BillNumberResponse {
	return &BillNumberResponse{
		BillNo:           response.GetBillNo(),
		Scac:             response.GetScac(),
		InfoAboutMoving:  c.convertGrpcInfoAboutMoving(response.InfoAboutMoving),
		EtaFinalDelivery: response.GetEtaFinalDelivery(),
	}
}
func (c *Converter) ConvertGrpcContainerNoResponse(response *TrackingByContainerNumberResponse) ContainerNumberResponse {
	return ContainerNumberResponse{
		Container:       response.GetContainer(),
		ContainerSize:   response.GetContainerSize(),
		Scac:            response.GetScac(),
		InfoAboutMoving: c.convertGrpcInfoAboutMoving(response.GetInfoAboutMoving()),
	}
}

type IClient interface {
	TrackByBillNumber(ctx context.Context, track *Track, _ string) (*BillNumberResponse, error)
	TrackByContainerNumber(ctx context.Context, track Track, ip string) (ContainerNumberResponse, error)
}
type Client struct {
	conn              *grpc.ClientConn
	billNoClient      trackingByBillNumberClient
	containerNoClient trackingByContainerNumberClient
	logger            logging.ILogger
	Converter
	util
}

func (c *Client) TrackByBillNumber(ctx context.Context, track *Track, _ string) (*BillNumberResponse, error) {
	request := Request{
		Number:  track.Number,
		Scac:    track.Scac,
		Country: Country(Country_value["RU"]),
	}
	response, err := c.billNoClient.TrackByBillNumber(ctx, &request)
	if err != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`trackingByBillNumber error: %s`, err.Error()))
		return new(BillNumberResponse), err
	}
	return c.ConvertGrpcBlNoResponse(response), nil
}

func (c *Client) TrackByContainerNumber(ctx context.Context, track Track, ip string) (ContainerNumberResponse, error) {
	country := c.getCountry(ip)
	var request Request
	if country == "RU" {
		request = Request{
			Number:  track.Number,
			Scac:    track.Scac,
			Country: Country(Country_value["RU"]),
		}
	} else {
		request = Request{
			Number:  track.Number,
			Scac:    track.Scac,
			Country: Country(Country_value["OTHER"]),
		}
	}
	response, err := c.containerNoClient.TrackByContainerNumber(ctx, &request)
	if err != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`trackingByContainerNumber error: %s`, err.Error()))
		return c.ConvertGrpcContainerNoResponse(response), err
	}
	return c.ConvertGrpcContainerNoResponse(response), nil
}
func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{
		conn:              conn,
		billNoClient:      trackingByBillNumberClient{cc: conn},
		containerNoClient: trackingByContainerNumberClient{cc: conn},
		logger:            logger,
	}
}
