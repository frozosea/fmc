package freight

import (
	"context"
	"fmc-newest/pkg/city"
	"fmc-newest/pkg/contact"
	"fmc-newest/pkg/line"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type freightRepoMoch struct {
}

func (s *freightRepoMoch) Get(_ context.Context, freight GetFreight) ([]BaseFreight, error) {
	var array []BaseFreight
	for i := 0; i < int(freight.Limit); i++ {
		array = append(array, BaseFreight{FromCity: city.City{Id: 1, BaseCity: city.BaseCity{FullName: "fromCity", Unlocode: "RUVVO"}}, ToCity: city.City{Id: 1, BaseCity: city.BaseCity{FullName: "fromCity", Unlocode: "RUVVO"}}, Container: Container{Id: 1, Type: "40DC"}, UsdPrice: i * 1000, Line: line.Line{Id: 1, BaseLine: line.BaseLine{Scac: "FESO", FullName: "Feso shipping line"}, ImageUrl: "http://localhost/img/1"}, Contact: contact.BaseContact{Url: "http://localhost", AgentName: "his name", PhoneNumber: "1654783200000"}})
	}
	return array, nil
}
func (s *freightRepoMoch) Add(ctx context.Context, freight AddFreight) error {
	return nil
}

// city contact container
type cityRepoMoch struct {
}

func (r *cityRepoMoch) Add(_ context.Context, city city.BaseCity) error {
	return nil
}
func (r *cityRepoMoch) GetAll(ctx context.Context) ([]*city.City, error) {
	var cities []*city.City
	return cities, nil
}

type contactRepoMoch struct {
}

func (r *contactRepoMoch) Add(ctx context.Context, contact contact.BaseContact) error {
	return nil
}

func (r *contactRepoMoch) GetAll(ctx context.Context) ([]*contact.BaseContact, error) {
	var contacts []*contact.BaseContact
	return contacts, nil
}

type containerRepoMoch struct {
}

func (r *containerRepoMoch) Add(ctx context.Context, containerType string) error {
	return nil
}
func (r *containerRepoMoch) Get(ctx context.Context) ([]*Container, error) {
	var containers []*Container
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

var service = NewController(&freightRepoMoch{}, &loggerMoch{})

func TestGetBestFreights(t *testing.T) {
	ctx := context.TODO()
	const limit = 20
	var getFreightRequestStruct = GetFreight{
		FromCityId:    0,
		ToCityId:      0,
		ContainerType: "40DC",
		Limit:         limit,
	}
	result, err := service.freightRepository.Get(ctx, getFreightRequestStruct)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, len(result), limit)
	for index, value := range result {
		assert.Equal(t, value.UsdPrice, index*1000)
	}
}
