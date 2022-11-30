package city

import pb "github.com/frozosea/fmc-pb/freight"

type BaseEntity struct {
	RuFullName string
	EnFullName string
}

type City struct {
	BaseEntity
	Id      int
	Country *Country
}

func (c *City) ToGrpc() *pb.City {
	return &pb.City{
		City: &pb.LocEntity{
			Id:     int64(c.Id),
			RuName: c.RuFullName,
			EnName: c.EnFullName,
		},
		Country: &pb.LocEntity{
			Id:     int64(c.Country.Id),
			RuName: c.Country.RuFullName,
			EnName: c.Country.EnFullName,
		},
	}
}

type Country struct {
	BaseEntity
	Id int
}

func (c *Country) ToGrpc() *pb.LocEntity {
	return &pb.LocEntity{
		Id:     int64(c.Id),
		RuName: c.RuFullName,
		EnName: c.EnFullName,
	}
}

type CountryWithId struct {
	BaseEntity
	CountryId int
}

type UpdateCity struct {
	Id int `json:"id"`
	CountryWithId
}

type Id struct {
	Id int `json:"id" form:"id"`
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
