// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: freight.proto

package proto

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

// FreightServiceClient is the client API for FreightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreightServiceClient interface {
	GetFreights(ctx context.Context, in *GetFreightRequest, opts ...grpc.CallOption) (*GetFreightsResponseList, error)
	AddFreight(ctx context.Context, in *AddFreightRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type freightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreightServiceClient(cc grpc.ClientConnInterface) FreightServiceClient {
	return &freightServiceClient{cc}
}

func (c *freightServiceClient) GetFreights(ctx context.Context, in *GetFreightRequest, opts ...grpc.CallOption) (*GetFreightsResponseList, error) {
	out := new(GetFreightsResponseList)
	err := c.cc.Invoke(ctx, "/proto_freight.FreightService/GetFreights", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) AddFreight(ctx context.Context, in *AddFreightRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto_freight.FreightService/AddFreight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreightServiceServer is the server API for FreightService service.
// All implementations must embed UnimplementedFreightServiceServer
// for forward compatibility
type FreightServiceServer interface {
	GetFreights(context.Context, *GetFreightRequest) (*GetFreightsResponseList, error)
	AddFreight(context.Context, *AddFreightRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedFreightServiceServer()
}

// UnimplementedFreightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFreightServiceServer struct {
}

func (UnimplementedFreightServiceServer) GetFreights(context.Context, *GetFreightRequest) (*GetFreightsResponseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreights not implemented")
}
func (UnimplementedFreightServiceServer) AddFreight(context.Context, *AddFreightRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFreight not implemented")
}
func (UnimplementedFreightServiceServer) mustEmbedUnimplementedFreightServiceServer() {}

// UnsafeFreightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreightServiceServer will
// result in compilation errors.
type UnsafeFreightServiceServer interface {
	mustEmbedUnimplementedFreightServiceServer()
}

func RegisterFreightServiceServer(s grpc.ServiceRegistrar, srv FreightServiceServer) {
	s.RegisterService(&FreightService_ServiceDesc, srv)
}

func _FreightService_GetFreights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFreightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).GetFreights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_freight.FreightService/GetFreights",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).GetFreights(ctx, req.(*GetFreightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_AddFreight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFreightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).AddFreight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_freight.FreightService/AddFreight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).AddFreight(ctx, req.(*AddFreightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FreightService_ServiceDesc is the grpc.ServiceDesc for FreightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_freight.FreightService",
	HandlerType: (*FreightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFreights",
			Handler:    _FreightService_GetFreights_Handler,
		},
		{
			MethodName: "AddFreight",
			Handler:    _FreightService_AddFreight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "freight.proto",
}

// CityServiceClient is the client API for CityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CityServiceClient interface {
	AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllCities(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCitiesResponse, error)
}

type cityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCityServiceClient(cc grpc.ClientConnInterface) CityServiceClient {
	return &cityServiceClient{cc}
}

func (c *cityServiceClient) AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto_freight.CityService/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) GetAllCities(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCitiesResponse, error) {
	out := new(GetAllCitiesResponse)
	err := c.cc.Invoke(ctx, "/proto_freight.CityService/GetAllCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityServiceServer is the server API for CityService service.
// All implementations must embed UnimplementedCityServiceServer
// for forward compatibility
type CityServiceServer interface {
	AddCity(context.Context, *AddCityRequest) (*emptypb.Empty, error)
	GetAllCities(context.Context, *emptypb.Empty) (*GetAllCitiesResponse, error)
	mustEmbedUnimplementedCityServiceServer()
}

// UnimplementedCityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCityServiceServer struct {
}

func (UnimplementedCityServiceServer) AddCity(context.Context, *AddCityRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedCityServiceServer) GetAllCities(context.Context, *emptypb.Empty) (*GetAllCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCities not implemented")
}
func (UnimplementedCityServiceServer) mustEmbedUnimplementedCityServiceServer() {}

// UnsafeCityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CityServiceServer will
// result in compilation errors.
type UnsafeCityServiceServer interface {
	mustEmbedUnimplementedCityServiceServer()
}

func RegisterCityServiceServer(s grpc.ServiceRegistrar, srv CityServiceServer) {
	s.RegisterService(&CityService_ServiceDesc, srv)
}

func _CityService_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_freight.CityService/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).AddCity(ctx, req.(*AddCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_GetAllCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetAllCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_freight.CityService/GetAllCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetAllCities(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CityService_ServiceDesc is the grpc.ServiceDesc for CityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_freight.CityService",
	HandlerType: (*CityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCity",
			Handler:    _CityService_AddCity_Handler,
		},
		{
			MethodName: "GetAllCities",
			Handler:    _CityService_GetAllCities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "freight.proto",
}

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactServiceClient interface {
	AddContact(ctx context.Context, in *AddContactRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllContacts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllContactsResponse, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) AddContact(ctx context.Context, in *AddContactRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto_freight.ContactService/AddContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) GetAllContacts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllContactsResponse, error) {
	out := new(GetAllContactsResponse)
	err := c.cc.Invoke(ctx, "/proto_freight.ContactService/GetAllContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
// All implementations must embed UnimplementedContactServiceServer
// for forward compatibility
type ContactServiceServer interface {
	AddContact(context.Context, *AddContactRequest) (*emptypb.Empty, error)
	GetAllContacts(context.Context, *emptypb.Empty) (*GetAllContactsResponse, error)
	mustEmbedUnimplementedContactServiceServer()
}

// UnimplementedContactServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (UnimplementedContactServiceServer) AddContact(context.Context, *AddContactRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContact not implemented")
}
func (UnimplementedContactServiceServer) GetAllContacts(context.Context, *emptypb.Empty) (*GetAllContactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllContacts not implemented")
}
func (UnimplementedContactServiceServer) mustEmbedUnimplementedContactServiceServer() {}

// UnsafeContactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactServiceServer will
// result in compilation errors.
type UnsafeContactServiceServer interface {
	mustEmbedUnimplementedContactServiceServer()
}

func RegisterContactServiceServer(s grpc.ServiceRegistrar, srv ContactServiceServer) {
	s.RegisterService(&ContactService_ServiceDesc, srv)
}

func _ContactService_AddContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).AddContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_freight.ContactService/AddContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).AddContact(ctx, req.(*AddContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_GetAllContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetAllContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_freight.ContactService/GetAllContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetAllContacts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ContactService_ServiceDesc is the grpc.ServiceDesc for ContactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_freight.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddContact",
			Handler:    _ContactService_AddContact_Handler,
		},
		{
			MethodName: "GetAllContacts",
			Handler:    _ContactService_GetAllContacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "freight.proto",
}

// LineServiceClient is the client API for LineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LineServiceClient interface {
	AddLine(ctx context.Context, opts ...grpc.CallOption) (LineService_AddLineClient, error)
	GetAllLines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllLinesResponse, error)
}

type lineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLineServiceClient(cc grpc.ClientConnInterface) LineServiceClient {
	return &lineServiceClient{cc}
}

func (c *lineServiceClient) AddLine(ctx context.Context, opts ...grpc.CallOption) (LineService_AddLineClient, error) {
	stream, err := c.cc.NewStream(ctx, &LineService_ServiceDesc.Streams[0], "/proto_freight.LineService/AddLine", opts...)
	if err != nil {
		return nil, err
	}
	x := &lineServiceAddLineClient{stream}
	return x, nil
}

type LineService_AddLineClient interface {
	Send(*AddLineRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type lineServiceAddLineClient struct {
	grpc.ClientStream
}

func (x *lineServiceAddLineClient) Send(m *AddLineRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *lineServiceAddLineClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lineServiceClient) GetAllLines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllLinesResponse, error) {
	out := new(GetAllLinesResponse)
	err := c.cc.Invoke(ctx, "/proto_freight.LineService/GetAllLines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LineServiceServer is the server API for LineService service.
// All implementations must embed UnimplementedLineServiceServer
// for forward compatibility
type LineServiceServer interface {
	AddLine(LineService_AddLineServer) error
	GetAllLines(context.Context, *emptypb.Empty) (*GetAllLinesResponse, error)
	mustEmbedUnimplementedLineServiceServer()
}

// UnimplementedLineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLineServiceServer struct {
}

func (UnimplementedLineServiceServer) AddLine(LineService_AddLineServer) error {
	return status.Errorf(codes.Unimplemented, "method AddLine not implemented")
}
func (UnimplementedLineServiceServer) GetAllLines(context.Context, *emptypb.Empty) (*GetAllLinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLines not implemented")
}
func (UnimplementedLineServiceServer) mustEmbedUnimplementedLineServiceServer() {}

// UnsafeLineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LineServiceServer will
// result in compilation errors.
type UnsafeLineServiceServer interface {
	mustEmbedUnimplementedLineServiceServer()
}

func RegisterLineServiceServer(s grpc.ServiceRegistrar, srv LineServiceServer) {
	s.RegisterService(&LineService_ServiceDesc, srv)
}

func _LineService_AddLine_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LineServiceServer).AddLine(&lineServiceAddLineServer{stream})
}

type LineService_AddLineServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*AddLineRequest, error)
	grpc.ServerStream
}

type lineServiceAddLineServer struct {
	grpc.ServerStream
}

func (x *lineServiceAddLineServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *lineServiceAddLineServer) Recv() (*AddLineRequest, error) {
	m := new(AddLineRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LineService_GetAllLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LineServiceServer).GetAllLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_freight.LineService/GetAllLines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LineServiceServer).GetAllLines(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// LineService_ServiceDesc is the grpc.ServiceDesc for LineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_freight.LineService",
	HandlerType: (*LineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllLines",
			Handler:    _LineService_GetAllLines_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddLine",
			Handler:       _LineService_AddLine_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "freight.proto",
}
