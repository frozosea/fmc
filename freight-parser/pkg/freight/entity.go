package freight

import (
	"fmc-newest/pkg/city"
	"fmc-newest/pkg/contact"
	"fmc-newest/pkg/line"
	"time"
)

type Container struct {
	Id   int
	Type string
}

type BaseFreight struct {
	Id       int
	FromCity city.City
	ToCity   city.City
	UsdPrice int
	Container
	line.Line
	FromDate    time.Time
	ExpiresDate time.Time
	Contact     contact.BaseContact
}

type AddFreight struct {
	FromCityId      int64
	ToCityId        int64
	ContainerTypeId int
	UsdPrice        int
	LineId          int64
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
