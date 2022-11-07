package tracking

type Track struct {
	Number string `protobuf:"number" json:"number" form:"number" validate:"min=7,max=28,regexp=[a-zA-Z]{2,}\d{5,}"`
	Scac   string `protobuf:"scac" json:"scac" form:"scac" validate:"min=4,max=4,regexp=[a-zA-z]{4}"`
}
type BaseInfoAboutMoving struct {
	Time          int64  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	OperationName string `protobuf:"bytes,2,opt,name=operation_name,json=operationName,proto3" json:"operation_name"`
	Location      string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Vessel        string `protobuf:"bytes,4,opt,name=vessel,proto3" json:"vessel,omitempty"`
}
type BillNumberResponse struct {
	BillNo           string                `protobuf:"bytes,1,opt,name=billNo,proto3" json:"number"`
	Scac             string                `protobuf:"varint,3,opt,name=Scac,proto3,enum=tracking.Scac" json:"scac"`
	InfoAboutMoving  []BaseInfoAboutMoving `protobuf:"bytes,4,rep,name=info_about_moving,json=infoAboutMoving,proto3" json:"infoAboutMoving"`
	EtaFinalDelivery int64                 `protobuf:"varint,5,opt,name=eta_final_delivery,json=etaFinalDelivery,proto3" json:"eta"`
}
type ContainerNumberResponse struct {
	Container       string                `protobuf:"bytes,1,opt,name=container,proto3" json:"number"`
	ContainerSize   string                `protobuf:"bytes,2,opt,name=container_size,json=containerSize,proto3" json:"containerSize"`
	Scac            string                `protobuf:"varint,3,opt,name=Scac,proto3,enum=tracking.Scac" json:"scac"`
	InfoAboutMoving []BaseInfoAboutMoving `protobuf:"bytes,4,rep,name=info_about_moving,json=infoAboutMoving,proto3" json:"infoAboutMoving"`
}
