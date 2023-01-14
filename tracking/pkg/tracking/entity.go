package tracking

import (
	pb "github.com/frozosea/fmc-pb/tracking"
	"time"
)

type Event struct {
	Time          time.Time
	OperationName string
	Location      string
	Vessel        string
}

func (e *Event) ToGRPC() *pb.InfoAboutMoving {
	return &pb.InfoAboutMoving{
		Time:          e.Time.UnixMilli(),
		OperationName: e.OperationName,
		Location:      e.Location,
		Vessel:        e.Vessel,
	}
}

type ContainerTrackingResponse struct {
	Number          string
	Size            string
	Scac            string
	InfoAboutMoving []*Event
}

func (c *ContainerTrackingResponse) ToGRPC() *pb.TrackingByContainerNumberResponse {
	var infoAboutMoving []*pb.InfoAboutMoving
	for _, v := range c.InfoAboutMoving {
		infoAboutMoving = append(infoAboutMoving, v.ToGRPC())
	}
	return &pb.TrackingByContainerNumberResponse{
		Container:       c.Number,
		ContainerSize:   c.Size,
		Scac:            c.Scac,
		InfoAboutMoving: infoAboutMoving,
	}
}

type BillNumberTrackingResponse struct {
	Number          string
	Eta             time.Time
	Scac            string
	InfoAboutMoving []*Event
}

func (b *BillNumberTrackingResponse) ToGRPC() *pb.TrackingByBillNumberResponse {
	var infoAboutMoving []*pb.InfoAboutMoving
	for _, v := range b.InfoAboutMoving {
		infoAboutMoving = append(infoAboutMoving, v.ToGRPC())
	}
	return &pb.TrackingByBillNumberResponse{
		BillNo:           b.Number,
		Scac:             b.Scac,
		InfoAboutMoving:  infoAboutMoving,
		EtaFinalDelivery: b.Eta.UnixMilli(),
	}
}
