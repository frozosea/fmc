package freight

import (
	"context"
	"fmt"
	"freight_service/internal/city"
	"freight_service/internal/company"
	"freight_service/internal/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

type freightRepoMoch struct {
}

func (s *freightRepoMoch) GetAll(_ context.Context) ([]BaseFreight, error) {
	var array []BaseFreight
	for i := 0; i < 20; i++ {
		array = append(array, BaseFreight{FromCity: city.City{Id: 1, BaseEntity: city.BaseEntity{RuFullName: "fromCity", EnFullName: "RUVVO"}}, ToCity: city.City{Id: 1, BaseEntity: city.BaseEntity{RuFullName: "fromCity", EnFullName: "RUVVO"}}, Container: container.Container{Id: 1, Type: "40DC"}, UsdPrice: i * 1000, Company: &company.Company{}})
	}
	return array, nil
}

func (s *freightRepoMoch) Update(_ context.Context, _ int, _ *AddFreight) error {
	return nil
}

func (s *freightRepoMoch) Delete(_ context.Context, _ int) error {
	return nil
}

func (s *freightRepoMoch) Get(_ context.Context, freight GetFreight) ([]BaseFreight, error) {
	var array []BaseFreight
	for i := 0; i < int(freight.Limit); i++ {
		array = append(array, BaseFreight{FromCity: city.City{Id: 1, BaseEntity: city.BaseEntity{RuFullName: "fromCity", EnFullName: "RUVVO"}}, ToCity: city.City{Id: 1, BaseEntity: city.BaseEntity{RuFullName: "fromCity", EnFullName: "RUVVO"}}, Container: container.Container{Id: 1, Type: "40DC"}, UsdPrice: i * 1000, Company: &company.Company{}})
	}
	return array, nil
}
func (s *freightRepoMoch) Add(_ context.Context, _ AddFreight) error {
	return nil
}

// city company container
type cityRepoMockUp struct {
}

func (r *cityRepoMockUp) Add(_ context.Context, _ city.BaseEntity) error {
	return nil
}
func (r *cityRepoMockUp) GetAll(_ context.Context) ([]*city.City, error) {
	var cities []*city.City
	return cities, nil
}

type contactRepoMockUp struct {
}

func (r *contactRepoMockUp) Add(_ context.Context, _ company.BaseCompany) error {
	return nil
}

func (r *contactRepoMockUp) GetAll(_ context.Context) ([]*company.BaseCompany, error) {
	var contacts []*company.BaseCompany
	return contacts, nil
}

type containerRepoMockUp struct {
}

func (r *containerRepoMockUp) Add(_ context.Context, _ string) error {
	return nil
}
func (r *containerRepoMockUp) Get(_ context.Context) ([]*container.Container, error) {
	var containers []*container.Container
	return containers, nil
}

type loggerMoch struct {
}

func (s loggerMoch) InfoLog(logString string) {
	fmt.Println(logString)
}
func (s loggerMoch) ExceptionLog(logString string) {
	fmt.Println(logString)

}
func (s loggerMoch) WarningLog(logString string) {
	fmt.Println(logString)

}
func (s loggerMoch) FatalLog(logString string) {
	fmt.Println(logString)

}

type cacheMoch struct{}

func (c *cacheMoch) Get(_ context.Context, _ string, _ interface{}) error { return nil }
func (c *cacheMoch) Set(_ context.Context, _ string, _ interface{}) error { return nil }
func (c *cacheMoch) Del(_ context.Context, _ string) error                { return nil }

var service = NewService(&freightRepoMoch{}, &loggerMoch{}, &cacheMoch{})

func TestGetBestFreights(t *testing.T) {
	ctx := context.TODO()
	const limit = 20
	var getFreightRequestStruct = GetFreight{
		FromCityId:      0,
		ToCityId:        0,
		ContainerTypeId: 1,
		Limit:           limit,
	}
	result, err := service.repo.Get(ctx, getFreightRequestStruct)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, len(result), limit)
	for index, value := range result {
		assert.Equal(t, value.UsdPrice, index*1000)
	}
}
