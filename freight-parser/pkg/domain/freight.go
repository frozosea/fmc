package domain

import "time"

type Contact struct {
	Id          int
	Url         string
	Email       string
	AgentName   string
	PhoneNumber string
}

type Container struct {
	Id   int
	Type string
}
type City struct {
	Id       int
	Name     string
	Unlocode string
}
type Line struct {
	LineId    int
	Scac      string
	LineName  string
	LineImage string
}

type BaseFreight struct {
	Id       int
	FromCity City
	ToCity   City
	UsdPrice int
	Container
	Line
	FromDate    time.Time
	ExpiresDate time.Time
	Contact     Contact
}

type GetFreight struct {
	FromCity      string
	ToCity        string
	ContainerType string
	Limit         int64
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

type AddCity struct {
	CityFullName string
	Unlocode     string
}
type GetCity struct {
	Id int
	AddCity
}

//const (
//	TWENTY_STANDARD containerType = iota
//	FORTY_STANDARD
//	FORTY_HIGH_CUBE
//	FORTY_FIVE_HIGH_CUBE
//)

//func (s containerType) ConvertToString() string {
//	switch s {
//	case 0:
//		return "20DC"
//	case 1:
//		return "40DC"
//	case 2:
//		return "40HC"
//	case 3:
//		return "45HC"
//	default:
//		return ""
//
//	}
//}
