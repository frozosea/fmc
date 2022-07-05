package tracking

import "time"

type Track struct {
	Number  string `protobuf:"number" json:"number" form:"number" validate:"min=1,max=25,regexp=[a-zA-Z]{4,}\d{3,}`
	Scac    string `protobuf:"scac" json:"scac" form:"scac" validate:"min=1,max=4,regexp=[a-zA-z]{4}`
	Country string `protobuf:"country" json:"country" form:"country" validate:"min=1,max=16,regexp=[a-zA-z]{2}`
}
type BaseInfoAboutMoving struct {
	Time          time.Time `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	OperationName string    `protobuf:"bytes,2,opt,name=operation_name,json=operationName,proto3" json:"operation_name,omitempty"`
	Location      string    `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Vessel        string    `protobuf:"bytes,4,opt,name=vessel,proto3" json:"vessel,omitempty"`
}
type BillNumberResponse struct {
	BillNo           string                `protobuf:"bytes,1,opt,name=billNo,proto3" json:"billNo,omitempty"`
	Scac             string                `protobuf:"varint,3,opt,name=Scac,proto3,enum=tracking.Scac" json:"Scac,omitempty"`
	InfoAboutMoving  []BaseInfoAboutMoving `protobuf:"bytes,4,rep,name=info_about_moving,json=infoAboutMoving,proto3" json:"info_about_moving,omitempty"`
	EtaFinalDelivery time.Time             `protobuf:"varint,5,opt,name=eta_final_delivery,json=etaFinalDelivery,proto3" json:"eta_final_delivery,omitempty"`
}
type ContainerNumberResponse struct {
	Container       string                `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	ContainerSize   string                `protobuf:"bytes,2,opt,name=container_size,json=containerSize,proto3" json:"container_size,omitempty"`
	Scac            string                `protobuf:"varint,3,opt,name=Scac,proto3,enum=tracking.Scac" json:"Scac,omitempty"`
	InfoAboutMoving []BaseInfoAboutMoving `protobuf:"bytes,4,rep,name=info_about_moving,json=infoAboutMoving,proto3" json:"info_about_moving,omitempty"`
}
