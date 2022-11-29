package freight

import (
	"freight_service/internal/city"
	"freight_service/internal/company"
	"freight_service/internal/container"
	pb "github.com/frozosea/fmc-proto/freight"
	"time"
)

type BaseFreight struct {
	Id          int
	FromCity    city.City
	ToCity      city.City
	UsdPrice    int
	Container   container.Container
	FromDate    time.Time
	ExpiresDate time.Time
	Company     *company.Company
}

func (b *BaseFreight) ToGrpc() *pb.GetFreightResponse {
	return &pb.GetFreightResponse{
		Id:            int64(b.Id),
		FromCity:      b.FromCity.ToGrpc(),
		ToCity:        b.ToCity.ToGrpc(),
		ContainerType: b.Container.ToGrpc(),
		UsdPrice:      int64(b.UsdPrice),
		Company:       b.Company.ToGrpc(),
	}
}

type AddFreight struct {
	FromCityId      int64
	ToCityId        int64
	ContainerTypeId int
	UsdPrice        int
	FromDate        time.Time
	ExpiresDate     time.Time
	ContactId       int
}
type GetFreight struct {
	FromCityId      int64
	ToCityId        int64
	ContainerTypeId int64
	Limit           int64
}

type UpdateFreight struct {
	Id int `json:"id"`
	AddFreight
}

type DeleteFreight struct {
	Id int `json:"id" form:"id"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
