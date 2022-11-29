package freight_service

import (
	"context"
	pb "github.com/frozosea/fmc-proto/freight"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IClient interface {
	GetFreights(ctx context.Context, get *GetFreight) ([]*Freight, error)
	GetCompanies(ctx context.Context) ([]*Company, error)
	GetContainers(ctx context.Context) ([]*Container, error)
	GetCities(ctx context.Context) ([]*City, error)
}

type Client struct {
	freightsCli   pb.FreightServiceClient
	companiesCli  pb.CompanyServiceClient
	containersCli pb.ContainersServiceClient
	citiesCli     pb.CityServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{freightsCli: pb.NewFreightServiceClient(conn), companiesCli: pb.NewCompanyServiceClient(conn), containersCli: pb.NewContainersServiceClient(conn), citiesCli: pb.NewCityServiceClient(conn)}
}

func (c *Client) GetFreights(ctx context.Context, get *GetFreight) ([]*Freight, error) {
	result, err := c.freightsCli.GetFreights(ctx, get.ToGrpc())
	if err != nil {
		return nil, err
	}
	var ar []*Freight
	for _, v := range result.GetMultiResponse() {
		ar = append(ar, new(Freight).FromGrpc(v))
	}
	return ar, nil
}

func (c *Client) GetCompanies(ctx context.Context) ([]*Company, error) {
	result, err := c.companiesCli.GetAllCompanies(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	var ar []*Company
	for _, v := range result.GetContact() {
		ar = append(ar, new(Company).FromGrpc(v))
	}
	return ar, nil
}

func (c *Client) GetContainers(ctx context.Context) ([]*Container, error) {
	result, err := c.containersCli.GetAllContainers(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	var ar []*Container
	for _, v := range result.GetContainers() {
		ar = append(ar, new(Container).FromGrpc(v))
	}
	return ar, nil
}

func (c *Client) GetCities(ctx context.Context) ([]*City, error) {
	result, err := c.citiesCli.GetAllCities(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	var ar []*City
	for _, v := range result.GetCities() {
		ar = append(ar, new(City).FromGrpc(v))
	}
	return ar, nil
}
