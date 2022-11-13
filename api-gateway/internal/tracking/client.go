package tracking

import (
	"context"
	"fmc-gateway/pkg/logging"
	pb "fmc-gateway/pkg/tracking-pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Converter struct{}

func (c *Converter) convertGrpcInfoAboutMoving(resp []*pb.InfoAboutMoving) []BaseInfoAboutMoving {
	var infoAboutMoving []BaseInfoAboutMoving
	for _, i := range resp {
		infoAboutMoving = append(infoAboutMoving, BaseInfoAboutMoving{Time: i.GetTime(), Location: i.GetLocation(), OperationName: i.GetOperationName(), Vessel: i.GetVessel()})
	}
	return infoAboutMoving
}
func (c *Converter) ConvertGrpcBlNoResponse(response *pb.TrackingByBillNumberResponse) *BillNumberResponse {
	return &BillNumberResponse{
		BillNo:           response.GetBillNo(),
		Scac:             response.GetScac(),
		InfoAboutMoving:  c.convertGrpcInfoAboutMoving(response.InfoAboutMoving),
		EtaFinalDelivery: response.GetEtaFinalDelivery(),
	}
}
func (c *Converter) ConvertGrpcContainerNoResponse(response *pb.TrackingByContainerNumberResponse) ContainerNumberResponse {
	return ContainerNumberResponse{
		Container:       response.GetContainer(),
		ContainerSize:   response.GetContainerSize(),
		Scac:            response.GetScac(),
		InfoAboutMoving: c.convertGrpcInfoAboutMoving(response.GetInfoAboutMoving()),
	}
}
func (c *Converter) ConvertScac(response *pb.GetAllScacResponse) []*Scac {
	var ar []*Scac
	for _, v := range response.GetAllScac() {
		ar = append(ar, &Scac{
			ScacCode: v.GetScac(),
			FullName: v.GetFullname(),
		})
	}
	return ar
}

type IClient interface {
	TrackByBillNumber(ctx context.Context, track *Track, _ string) (*BillNumberResponse, error)
	TrackByContainerNumber(ctx context.Context, track Track, ip string) (ContainerNumberResponse, error)
	GetContainerScac(ctx context.Context) ([]*Scac, error)
	GetBillScac(ctx context.Context) ([]*Scac, error)
}
type Client struct {
	conn              *grpc.ClientConn
	billNoClient      pb.TrackingByBillNumberClient
	containerNoClient pb.TrackingByContainerNumberClient
	scacClient        pb.ScacServiceClient
	logger            logging.ILogger
	Converter
	util
}

func (c *Client) TrackByBillNumber(ctx context.Context, track *Track, _ string) (*BillNumberResponse, error) {
	request := pb.Request{
		Number:  track.Number,
		Scac:    track.Scac,
		Country: pb.Country(pb.Country_value["RU"]),
	}
	response, err := c.billNoClient.TrackByBillNumber(ctx, &request)
	if err != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`trackingByBillNumber error: %s`, err.Error()))
		return new(BillNumberResponse), err
	}
	return c.ConvertGrpcBlNoResponse(response), nil
}

func (c *Client) TrackByContainerNumber(ctx context.Context, track Track, ip string) (ContainerNumberResponse, error) {
	request := pb.Request{
		Number:  track.Number,
		Scac:    track.Scac,
		Country: pb.Country(pb.Country_value["OTHER"]),
	}
	response, err := c.containerNoClient.TrackByContainerNumber(ctx, &request)
	if err != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`trackingByContainerNumber error: %s`, err.Error()))
		return c.ConvertGrpcContainerNoResponse(response), err
	}
	return c.ConvertGrpcContainerNoResponse(response), nil
}
func (c *Client) GetContainerScac(ctx context.Context) ([]*Scac, error) {
	data, err := c.scacClient.GetContainerScac(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return c.Converter.ConvertScac(data), nil
}

func (c *Client) GetBillScac(ctx context.Context) ([]*Scac, error) {
	data, err := c.scacClient.GetBillScac(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return c.Converter.ConvertScac(data), nil
}
func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{
		conn:              conn,
		billNoClient:      pb.NewTrackingByBillNumberClient(conn),
		containerNoClient: pb.NewTrackingByContainerNumberClient(conn),
		scacClient:        pb.NewScacServiceClient(conn),
		logger:            logger,
	}
}
