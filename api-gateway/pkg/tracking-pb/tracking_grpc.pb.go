// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.6
// source: tracking.proto

package tracking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TrackingByContainerNumberClient is the client API for TrackingByContainerNumber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackingByContainerNumberClient interface {
	TrackByContainerNumber(ctx context.Context, in *Request, opts ...grpc.CallOption) (*TrackingByContainerNumberResponse, error)
}

type trackingByContainerNumberClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackingByContainerNumberClient(cc grpc.ClientConnInterface) TrackingByContainerNumberClient {
	return &trackingByContainerNumberClient{cc}
}

func (c *trackingByContainerNumberClient) TrackByContainerNumber(ctx context.Context, in *Request, opts ...grpc.CallOption) (*TrackingByContainerNumberResponse, error) {
	out := new(TrackingByContainerNumberResponse)
	err := c.cc.Invoke(ctx, "/tracking.TrackingByContainerNumber/TrackByContainerNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackingByContainerNumberServer is the server API for TrackingByContainerNumber service.
// All implementations must embed UnimplementedTrackingByContainerNumberServer
// for forward compatibility
type TrackingByContainerNumberServer interface {
	TrackByContainerNumber(context.Context, *Request) (*TrackingByContainerNumberResponse, error)
	mustEmbedUnimplementedTrackingByContainerNumberServer()
}

// UnimplementedTrackingByContainerNumberServer must be embedded to have forward compatible implementations.
type UnimplementedTrackingByContainerNumberServer struct {
}

func (UnimplementedTrackingByContainerNumberServer) TrackByContainerNumber(context.Context, *Request) (*TrackingByContainerNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackByContainerNumber not implemented")
}
func (UnimplementedTrackingByContainerNumberServer) mustEmbedUnimplementedTrackingByContainerNumberServer() {
}

// UnsafeTrackingByContainerNumberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackingByContainerNumberServer will
// result in compilation errors.
type UnsafeTrackingByContainerNumberServer interface {
	mustEmbedUnimplementedTrackingByContainerNumberServer()
}

func RegisterTrackingByContainerNumberServer(s grpc.ServiceRegistrar, srv TrackingByContainerNumberServer) {
	s.RegisterService(&TrackingByContainerNumber_ServiceDesc, srv)
}

func _TrackingByContainerNumber_TrackByContainerNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingByContainerNumberServer).TrackByContainerNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracking.TrackingByContainerNumber/TrackByContainerNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingByContainerNumberServer).TrackByContainerNumber(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// TrackingByContainerNumber_ServiceDesc is the grpc.ServiceDesc for TrackingByContainerNumber service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrackingByContainerNumber_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracking.TrackingByContainerNumber",
	HandlerType: (*TrackingByContainerNumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrackByContainerNumber",
			Handler:    _TrackingByContainerNumber_TrackByContainerNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracking.proto",
}

// TrackingByBillNumberClient is the client API for TrackingByBillNumber service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrackingByBillNumberClient interface {
	TrackByBillNumber(ctx context.Context, in *Request, opts ...grpc.CallOption) (*TrackingByBillNumberResponse, error)
}

type trackingByBillNumberClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackingByBillNumberClient(cc grpc.ClientConnInterface) TrackingByBillNumberClient {
	return &trackingByBillNumberClient{cc}
}

func (c *trackingByBillNumberClient) TrackByBillNumber(ctx context.Context, in *Request, opts ...grpc.CallOption) (*TrackingByBillNumberResponse, error) {
	out := new(TrackingByBillNumberResponse)
	err := c.cc.Invoke(ctx, "/tracking.TrackingByBillNumber/TrackByBillNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackingByBillNumberServer is the server API for TrackingByBillNumber service.
// All implementations must embed UnimplementedTrackingByBillNumberServer
// for forward compatibility
type TrackingByBillNumberServer interface {
	TrackByBillNumber(context.Context, *Request) (*TrackingByBillNumberResponse, error)
	mustEmbedUnimplementedTrackingByBillNumberServer()
}

// UnimplementedTrackingByBillNumberServer must be embedded to have forward compatible implementations.
type UnimplementedTrackingByBillNumberServer struct {
}

func (UnimplementedTrackingByBillNumberServer) TrackByBillNumber(context.Context, *Request) (*TrackingByBillNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackByBillNumber not implemented")
}
func (UnimplementedTrackingByBillNumberServer) mustEmbedUnimplementedTrackingByBillNumberServer() {}

// UnsafeTrackingByBillNumberServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackingByBillNumberServer will
// result in compilation errors.
type UnsafeTrackingByBillNumberServer interface {
	mustEmbedUnimplementedTrackingByBillNumberServer()
}

func RegisterTrackingByBillNumberServer(s grpc.ServiceRegistrar, srv TrackingByBillNumberServer) {
	s.RegisterService(&TrackingByBillNumber_ServiceDesc, srv)
}

func _TrackingByBillNumber_TrackByBillNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackingByBillNumberServer).TrackByBillNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracking.TrackingByBillNumber/TrackByBillNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackingByBillNumberServer).TrackByBillNumber(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// TrackingByBillNumber_ServiceDesc is the grpc.ServiceDesc for TrackingByBillNumber service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrackingByBillNumber_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracking.TrackingByBillNumber",
	HandlerType: (*TrackingByBillNumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrackByBillNumber",
			Handler:    _TrackingByBillNumber_TrackByBillNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracking.proto",
}

// ScacServiceClient is the client API for ScacService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScacServiceClient interface {
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllScacResponse, error)
}

type scacServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScacServiceClient(cc grpc.ClientConnInterface) ScacServiceClient {
	return &scacServiceClient{cc}
}

func (c *scacServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllScacResponse, error) {
	out := new(GetAllScacResponse)
	err := c.cc.Invoke(ctx, "/tracking.ScacService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScacServiceServer is the server API for ScacService service.
// All implementations must embed UnimplementedScacServiceServer
// for forward compatibility
type ScacServiceServer interface {
	GetAll(context.Context, *emptypb.Empty) (*GetAllScacResponse, error)
	mustEmbedUnimplementedScacServiceServer()
}

// UnimplementedScacServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScacServiceServer struct {
}

func (UnimplementedScacServiceServer) GetAll(context.Context, *emptypb.Empty) (*GetAllScacResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedScacServiceServer) mustEmbedUnimplementedScacServiceServer() {}

// UnsafeScacServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScacServiceServer will
// result in compilation errors.
type UnsafeScacServiceServer interface {
	mustEmbedUnimplementedScacServiceServer()
}

func RegisterScacServiceServer(s grpc.ServiceRegistrar, srv ScacServiceServer) {
	s.RegisterService(&ScacService_ServiceDesc, srv)
}

func _ScacService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScacServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracking.ScacService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScacServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ScacService_ServiceDesc is the grpc.ServiceDesc for ScacService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScacService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracking.ScacService",
	HandlerType: (*ScacServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _ScacService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracking.proto",
}