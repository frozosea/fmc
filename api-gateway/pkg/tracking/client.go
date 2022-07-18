package tracking

import (
	"context"
	"fmc-with-git/internal/logging"
	"fmt"
	"google.golang.org/grpc"
)

type converter struct{}

func (c *converter) convertGrpcInfoAboutMoving(resp []*InfoAboutMoving) []BaseInfoAboutMoving {
	var infoAboutMoving []BaseInfoAboutMoving
	for _, i := range resp {
		infoAboutMoving = append(infoAboutMoving, BaseInfoAboutMoving{Time: i.GetTime(), Location: i.GetLocation(), OperationName: i.GetOperationName(), Vessel: i.GetVessel()})
	}
	return infoAboutMoving
}
func (c *converter) convertGrpcBlNoResponse(response *TrackingByBillNumberResponse) *BillNumberResponse {
	return &BillNumberResponse{
		BillNo:           response.GetBillNo(),
		Scac:             response.GetScac().String(),
		InfoAboutMoving:  c.convertGrpcInfoAboutMoving(response.InfoAboutMoving),
		EtaFinalDelivery: response.GetEtaFinalDelivery(),
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

type Client struct {
	conn              *grpc.ClientConn
	billNoClient      trackingByBillNumberClient
	containerNoClient trackingByContainerNumberClient
	logger            logging.ILogger
	converter
	util
}

func (c *Client) TrackByBillNumber(ctx context.Context, track *Track, _ string) (*BillNumberResponse, error) {
	request := Request{
		Number:  track.Number,
		Scac:    Scac(Scac_value[track.Scac]),
		Country: Country(Country_value["RU"]),
	}
	//}
	fmt.Println(Scac(Scac_value[track.Scac]))
	response, err := c.billNoClient.TrackByBillNumber(ctx, &request)
	if err != nil {
		//c.logger.ExceptionLog(fmt.Sprintf(`trackingByBillNumber error: %s`, err.Error()))
		return new(BillNumberResponse), err
	}
	return c.convertGrpcBlNoResponse(response), nil
}

func (c *Client) TrackByContainerNumber(ctx context.Context, track Track, ip string) (ContainerNumberResponse, error) {
	country := c.getCountry(ip)
	var request Request
	if country == "RU" {
		request = Request{
			Number:  track.Number,
			Scac:    Scac(Scac_value[track.Scac]),
			Country: Country(Country_value["RU"]),
		}
	} else {
		request = Request{
			Number:  track.Number,
			Scac:    Scac(Scac_value[track.Scac]),
			Country: Country(Country_value["OTHER"]),
		}
	}
	response, err := c.containerNoClient.TrackByContainerNumber(ctx, &request)
	if err != nil {
		fmt.Println(err.Error())
		c.logger.ExceptionLog(fmt.Sprintf(`trackingByContainerNumber error: %s`, err.Error()))
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
