package tracking

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/frozosea/fmc-pb/tracking"
	"golang_tracking/pkg/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ContainerTrackingGrpc struct {
	service *ContainerTrackingService
	logger  logging.ILogger
	*pb.UnimplementedTrackingByContainerNumberServer
}

func NewContainerTrackingGrpc(service *ContainerTrackingService, logger logging.ILogger) *ContainerTrackingGrpc {
	return &ContainerTrackingGrpc{
		service: service,
		logger:  logger,
		UnimplementedTrackingByContainerNumberServer: &pb.UnimplementedTrackingByContainerNumberServer{},
	}
}

func (c *ContainerTrackingGrpc) TrackByContainerNumber(ctx context.Context, r *pb.Request) (*pb.TrackingByContainerNumberResponse, error) {
	go c.logger.InfoLog(fmt.Sprintf(`tracking container: %s; scac: %s`, r.GetNumber(), r.GetScac()))
	response, err := c.service.Track(ctx, r.GetScac(), r.GetNumber())
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`tracking container: %s; scac: %s error: %s`, r.GetNumber(), r.GetScac(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	//go func() {
	//	j, err := json.Marshal(response)
	//	if err != nil {
	//		return
	//	}
	//	c.logger.InfoLog(fmt.Sprintf(`tracking container: %s; scac: %s result: %v"`, r.GetNumber(), r.GetScac(), j))
	//}()
	return response.ToGRPC(), nil
}

type BillNumberTrackingGrpc struct {
	service *BillTrackingService
	logger  logging.ILogger
	*pb.UnimplementedTrackingByBillNumberServer
}

func NewBillNumberTrackingGrpc(service *BillTrackingService, logger logging.ILogger) *BillNumberTrackingGrpc {
	return &BillNumberTrackingGrpc{
		service:                                 service,
		logger:                                  logger,
		UnimplementedTrackingByBillNumberServer: &pb.UnimplementedTrackingByBillNumberServer{},
	}
}
func (b *BillNumberTrackingGrpc) TrackByBillNumber(ctx context.Context, r *pb.Request) (*pb.TrackingByBillNumberResponse, error) {
	go b.logger.InfoLog(fmt.Sprintf(`tracking bill: %s; scac: %s`, r.GetNumber(), r.GetScac()))
	response, err := b.service.Track(ctx, r.GetScac(), r.GetNumber())
	if err != nil {
		go b.logger.ExceptionLog(fmt.Sprintf(`tracking bill: %s; scac: %s error: %s`, r.GetNumber(), r.GetScac(), err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}
	go func() {
		j, err := json.Marshal(response)
		if err != nil {
			return
		}
		b.logger.InfoLog(fmt.Sprintf(`tracking bill: %s; scac: %s result: %v"`, r.GetNumber(), r.GetScac(), j))
	}()
	return response.ToGRPC(), nil
}
