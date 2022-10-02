package archive

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "schedule-tracking/pkg/proto"
	"schedule-tracking/pkg/tracking"
)

type converter struct {
}

func newConverter() *converter {
	return &converter{}
}

func (c *converter) convertInfoAboutMoving(r []tracking.BaseInfoAboutMoving) []*pb.InfoAboutMoving {
	var ar []*pb.InfoAboutMoving
	for _, v := range r {
		ar = append(ar, &pb.InfoAboutMoving{
			Time:          v.Time.Unix(),
			OperationName: v.OperationName,
			Location:      v.Location,
			Vessel:        v.Vessel,
		})
	}
	return ar
}

func (c *converter) convertBillResponse(r *tracking.BillNumberResponse) *pb.TrackingByBillNumberResponse {
	return &pb.TrackingByBillNumberResponse{
		BillNo:           r.BillNo,
		Scac:             r.Scac,
		InfoAboutMoving:  c.convertInfoAboutMoving(r.InfoAboutMoving),
		EtaFinalDelivery: r.EtaFinalDelivery.Unix(),
	}
}
func (c *converter) convertContainerResponse(r *tracking.ContainerNumberResponse) *pb.TrackingByContainerNumberResponse {
	return &pb.TrackingByContainerNumberResponse{
		Container:       r.Container,
		ContainerSize:   r.ContainerSize,
		Scac:            r.Scac,
		InfoAboutMoving: c.convertInfoAboutMoving(r.InfoAboutMoving),
	}
}
func (c *converter) convertContainersResponse(r []*tracking.ContainerNumberResponse) []*pb.TrackingByContainerNumberResponse {
	var ar []*pb.TrackingByContainerNumberResponse
	for _, v := range r {
		ar = append(ar, c.convertContainerResponse(v))
	}
	return ar
}
func (c *converter) convertBillsResponse(r []*tracking.BillNumberResponse) []*pb.TrackingByBillNumberResponse {
	var ar []*pb.TrackingByBillNumberResponse
	for _, v := range r {
		ar = append(ar, c.convertBillResponse(v))
	}
	return ar
}

type Grpc struct {
	service *Service
	pb.UnimplementedArchiveServer
	converter *converter
}

func NewGrpc(service *Service) *Grpc {
	return &Grpc{service: service, UnimplementedArchiveServer: pb.UnimplementedArchiveServer{}, converter: newConverter()}
}

func (g *Grpc) GetAllBillsContainers(ctx context.Context, r *pb.GetAllBillsContainerRequest) (*pb.GetAllBillsContainerResponse, error) {
	result, err := g.service.GetAll(ctx, int(r.GetUserId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetAllBillsContainerResponse{
		Bills:      g.converter.convertBillsResponse(result.bills),
		Containers: g.converter.convertContainersResponse(result.containers),
	}, nil
}
