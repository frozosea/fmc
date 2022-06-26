package client

import (
	pb "fmc-newest/internal/proto"
)

//TODO create all clients
type Freight struct {
	pb.FreightServiceClient
}

//func (f *Freight) GetFreights(ctx context.Context, in *pb.GetFreightRequest, opts ...grpc.CallOption) (*pb.GetFreightsResponseList, error) {
//
//}
//
//func (f *Freight) AddFreight(ctx context.Context, in *pb.AddFreightRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
//
//}
