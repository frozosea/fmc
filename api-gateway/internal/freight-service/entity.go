package freight_service

import (
	pb "github.com/frozosea/fmc-pb/freight"

	"time"
)

type GetFreight struct {
	FromCityId      int64
	ToCityId        int64
	ContainerTypeId int64
	Limit           int64
}

func (g *GetFreight) ToGrpc() *pb.GetFreightRequest {
	return &pb.GetFreightRequest{
		FromCityId:      g.FromCityId,
		ToCityId:        g.ToCityId,
		ContainerTypeId: g.ContainerTypeId,
		Limit:           g.Limit,
	}
}

type UpdateFreight struct {
	Id              int `json:"id"`
	FromCityId      int64
	ToCityId        int64
	ContainerTypeId int
	UsdPrice        int
	FromDate        time.Time
	ExpiresDate     time.Time
	ContactId       int
}

type DeleteFreight struct {
	Id int `json:"id" form:"id"`
}

type Container struct {
	Id   int
	Type string
}

func (c *Container) ToGrpc() *pb.Container {
	return &pb.Container{
		ContainerType:   c.Type,
		ContainerTypeId: int64(c.Id),
	}
}
func (c *Container) FromGrpc(container *pb.Container) *Container {
	*c = Container{
		Id:   int(container.GetContainerTypeId()),
		Type: container.GetContainerType(),
	}
	return c
}

type BaseCompany struct {
	Url         string
	Email       string
	Name        string
	PhoneNumber string
}

type Company struct {
	Id int
	*BaseCompany
}

func (c *Company) ToGrpc() *pb.Company {
	return &pb.Company{
		Id:          int64(c.Id),
		Url:         c.Url,
		PhoneNumber: c.PhoneNumber,
		AgentName:   c.Name,
		Email:       c.Email,
	}
}
func (c *Company) FromGrpc(company *pb.Company) *Company {
	*c = Company{
		Id: int(company.Id),
		BaseCompany: &BaseCompany{
			Url:         company.GetUrl(),
			Email:       company.GetEmail(),
			Name:        company.GetAgentName(),
			PhoneNumber: company.GetPhoneNumber(),
		},
	}
	return c
}

type BaseLocation struct {
	RuFullName string
	EnFullName string
}

type City struct {
	BaseLocation
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
func (c *City) FromGrpc(city *pb.City) *City {
	pbCity := city.GetCity()
	pbCountry := city.GetCountry()
	*c = City{
		BaseLocation: BaseLocation{
			RuFullName: pbCity.GetRuName(),
			EnFullName: pbCity.GetEnName(),
		},
		Id: int(pbCity.GetId()),
		Country: &Country{
			BaseLocation: BaseLocation{
				RuFullName: pbCountry.GetRuName(),
				EnFullName: pbCountry.GetEnName(),
			},
			Id: int(pbCountry.GetId()),
		},
	}
	return c
}

type Country struct {
	BaseLocation
	Id int
}

func (c *Country) ToGrpc() *pb.LocEntity {
	return &pb.LocEntity{
		Id:     int64(c.Id),
		RuName: c.RuFullName,
		EnName: c.EnFullName,
	}
}
func (c *Country) FromGrpc(country *pb.LocEntity) *Country {
	*c = Country{
		BaseLocation: BaseLocation{
			RuFullName: country.GetRuName(),
			EnFullName: country.GetEnName(),
		},
		Id: int(country.GetId()),
	}
	return c
}

type CountryWithId struct {
	BaseLocation
	CountryId int
}

type UpdateCity struct {
	Id int `json:"id"`
	CountryWithId
}

type Freight struct {
	Id        int
	FromCity  *City
	ToCity    *City
	UsdPrice  int
	Container *Container
	Company   *Company
}

func (b *Freight) ToGrpc() *pb.GetFreightResponse {
	return &pb.GetFreightResponse{
		FromCity:      b.FromCity.ToGrpc(),
		ToCity:        b.ToCity.ToGrpc(),
		ContainerType: b.Container.ToGrpc(),
		UsdPrice:      int64(b.UsdPrice),
		Company:       b.Company.ToGrpc(),
	}
}

func (b *Freight) FromGrpc(freight *pb.GetFreightResponse) *Freight {
	*b = Freight{
		Id:        int(freight.GetId()),
		FromCity:  new(City).FromGrpc(freight.GetFromCity()),
		ToCity:    new(City).FromGrpc(freight.GetToCity()),
		UsdPrice:  int(freight.GetUsdPrice()),
		Container: new(Container).FromGrpc(freight.GetContainerType()),
		Company:   new(Company).FromGrpc(freight.GetCompany()),
	}
	return b
}

type BaseResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
