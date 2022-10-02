// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: tracking.proto

package tracking

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Country int32

const (
	Country_RU    Country = 0
	Country_OTHER Country = 1
)

// Enum value maps for Country.
var (
	Country_name = map[int32]string{
		0: "RU",
		1: "OTHER",
	}
	Country_value = map[string]int32{
		"RU":    0,
		"OTHER": 1,
	}
)

func (x Country) Enum() *Country {
	p := new(Country)
	*p = x
	return p
}

func (x Country) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Country) Descriptor() protoreflect.EnumDescriptor {
	return file_tracking_proto_enumTypes[0].Descriptor()
}

func (Country) Type() protoreflect.EnumType {
	return &file_tracking_proto_enumTypes[0]
}

func (x Country) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Country.Descriptor instead.
func (Country) EnumDescriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number  string  `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Scac    string  `protobuf:"bytes,2,opt,name=scac,proto3" json:"scac,omitempty"`
	Country Country `protobuf:"varint,3,opt,name=country,proto3,enum=tracking.Country" json:"country,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Request) GetScac() string {
	if x != nil {
		return x.Scac
	}
	return ""
}

func (x *Request) GetCountry() Country {
	if x != nil {
		return x.Country
	}
	return Country_RU
}

type InfoAboutMoving struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time          int64  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	OperationName string `protobuf:"bytes,2,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	Location      string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Vessel        string `protobuf:"bytes,4,opt,name=vessel,proto3" json:"vessel,omitempty"`
}

func (x *InfoAboutMoving) Reset() {
	*x = InfoAboutMoving{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoAboutMoving) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoAboutMoving) ProtoMessage() {}

func (x *InfoAboutMoving) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoAboutMoving.ProtoReflect.Descriptor instead.
func (*InfoAboutMoving) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{1}
}

func (x *InfoAboutMoving) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *InfoAboutMoving) GetOperationName() string {
	if x != nil {
		return x.OperationName
	}
	return ""
}

func (x *InfoAboutMoving) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *InfoAboutMoving) GetVessel() string {
	if x != nil {
		return x.Vessel
	}
	return ""
}

type TrackingByContainerNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Container       string             `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	ContainerSize   string             `protobuf:"bytes,2,opt,name=container_size,json=containerSize,proto3" json:"container_size,omitempty"`
	Scac            string             `protobuf:"bytes,3,opt,name=scac,proto3" json:"scac,omitempty"`
	InfoAboutMoving []*InfoAboutMoving `protobuf:"bytes,4,rep,name=info_about_moving,json=infoAboutMoving,proto3" json:"info_about_moving,omitempty"`
}

func (x *TrackingByContainerNumberResponse) Reset() {
	*x = TrackingByContainerNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackingByContainerNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackingByContainerNumberResponse) ProtoMessage() {}

func (x *TrackingByContainerNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackingByContainerNumberResponse.ProtoReflect.Descriptor instead.
func (*TrackingByContainerNumberResponse) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{2}
}

func (x *TrackingByContainerNumberResponse) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *TrackingByContainerNumberResponse) GetContainerSize() string {
	if x != nil {
		return x.ContainerSize
	}
	return ""
}

func (x *TrackingByContainerNumberResponse) GetScac() string {
	if x != nil {
		return x.Scac
	}
	return ""
}

func (x *TrackingByContainerNumberResponse) GetInfoAboutMoving() []*InfoAboutMoving {
	if x != nil {
		return x.InfoAboutMoving
	}
	return nil
}

type TrackingByBillNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillNo           string             `protobuf:"bytes,1,opt,name=billNo,proto3" json:"billNo,omitempty"`
	Scac             string             `protobuf:"bytes,3,opt,name=scac,proto3" json:"scac,omitempty"`
	InfoAboutMoving  []*InfoAboutMoving `protobuf:"bytes,4,rep,name=info_about_moving,json=infoAboutMoving,proto3" json:"info_about_moving,omitempty"`
	EtaFinalDelivery int64              `protobuf:"varint,5,opt,name=eta_final_delivery,json=etaFinalDelivery,proto3" json:"eta_final_delivery,omitempty"`
}

func (x *TrackingByBillNumberResponse) Reset() {
	*x = TrackingByBillNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackingByBillNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackingByBillNumberResponse) ProtoMessage() {}

func (x *TrackingByBillNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackingByBillNumberResponse.ProtoReflect.Descriptor instead.
func (*TrackingByBillNumberResponse) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{3}
}

func (x *TrackingByBillNumberResponse) GetBillNo() string {
	if x != nil {
		return x.BillNo
	}
	return ""
}

func (x *TrackingByBillNumberResponse) GetScac() string {
	if x != nil {
		return x.Scac
	}
	return ""
}

func (x *TrackingByBillNumberResponse) GetInfoAboutMoving() []*InfoAboutMoving {
	if x != nil {
		return x.InfoAboutMoving
	}
	return nil
}

func (x *TrackingByBillNumberResponse) GetEtaFinalDelivery() int64 {
	if x != nil {
		return x.EtaFinalDelivery
	}
	return 0
}

var File_tracking_proto protoreflect.FileDescriptor

var file_tracking_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x62, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x63, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x63, 0x61,
	0x63, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x80,
	0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x22, 0xc3, 0x01, 0x0a, 0x21, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x63, 0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x63, 0x61, 0x63,
	0x12, 0x45, 0x0a, 0x11, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x6d,
	0x6f, 0x76, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x62, 0x6f, 0x75, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x69, 0x6e, 0x66, 0x6f, 0x41, 0x62, 0x6f, 0x75,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x22, 0xbf, 0x01, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x69, 0x6c, 0x6c,
	0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x69, 0x6c, 0x6c, 0x4e, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x63, 0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x63, 0x61, 0x63, 0x12, 0x45, 0x0a, 0x11, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x62, 0x6f,
	0x75, 0x74, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x41,
	0x62, 0x6f, 0x75, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x69, 0x6e, 0x66, 0x6f,
	0x41, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x12, 0x65,
	0x74, 0x61, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x65, 0x74, 0x61, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2a, 0x1c, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x55, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x32, 0x77, 0x0a, 0x19, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x5a, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x42, 0x79, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x11,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x32, 0x68, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x42, 0x69,
	0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x42, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x11, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x42, 0x69, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x2e, 0x2e,
	0x2f, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracking_proto_rawDescOnce sync.Once
	file_tracking_proto_rawDescData = file_tracking_proto_rawDesc
)

func file_tracking_proto_rawDescGZIP() []byte {
	file_tracking_proto_rawDescOnce.Do(func() {
		file_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracking_proto_rawDescData)
	})
	return file_tracking_proto_rawDescData
}

var file_tracking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tracking_proto_goTypes = []interface{}{
	(Country)(0),            // 0: tracking.Country
	(*Request)(nil),         // 1: tracking.Request
	(*InfoAboutMoving)(nil), // 2: tracking.InfoAboutMoving
	(*TrackingByContainerNumberResponse)(nil), // 3: tracking.TrackingByContainerNumberResponse
	(*TrackingByBillNumberResponse)(nil),      // 4: tracking.TrackingByBillNumberResponse
}
var file_tracking_proto_depIdxs = []int32{
	0, // 0: tracking.Request.country:type_name -> tracking.Country
	2, // 1: tracking.TrackingByContainerNumberResponse.info_about_moving:type_name -> tracking.InfoAboutMoving
	2, // 2: tracking.TrackingByBillNumberResponse.info_about_moving:type_name -> tracking.InfoAboutMoving
	1, // 3: tracking.TrackingByContainerNumber.TrackByContainerNumber:input_type -> tracking.Request
	1, // 4: tracking.TrackingByBillNumber.TrackByBillNumber:input_type -> tracking.Request
	3, // 5: tracking.TrackingByContainerNumber.TrackByContainerNumber:output_type -> tracking.TrackingByContainerNumberResponse
	4, // 6: tracking.TrackingByBillNumber.TrackByBillNumber:output_type -> tracking.TrackingByBillNumberResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tracking_proto_init() }
func file_tracking_proto_init() {
	if File_tracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tracking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoAboutMoving); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tracking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackingByContainerNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tracking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackingByBillNumberResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tracking_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_tracking_proto_goTypes,
		DependencyIndexes: file_tracking_proto_depIdxs,
		EnumInfos:         file_tracking_proto_enumTypes,
		MessageInfos:      file_tracking_proto_msgTypes,
	}.Build()
	File_tracking_proto = out.File
	file_tracking_proto_rawDesc = nil
	file_tracking_proto_goTypes = nil
	file_tracking_proto_depIdxs = nil
}
