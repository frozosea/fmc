package freight

import (
	"context"
	"fmc-newest/pkg/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type freightRepoMoch struct {
}

func (s *freightRepoMoch) Get(_ context.Context, freight domain.GetFreight) ([]domain.BaseFreight, error) {
	var array []domain.BaseFreight
	for i := 0; i < int(freight.Limit); i++ {
		array = append(array, domain.BaseFreight{FromCity: domain.City{Id: 1, Name: "fromCity", Unlocode: "RUVVO"}, ToCity: domain.City{Id: 1, Name: "toCity", Unlocode: "RUVFP"}, Container: domain.Container{Id: 1, Type: "40DC"}, UsdPrice: i * 1000, Line: domain.Line{LineId: 1, Scac: "FESO", LineName: "Feso shipping line", LineImage: "http://localhost/img/1"}, Contact: domain.Contact{Url: "http://localhost", AgentName: "his name", PhoneNumber: "1654783200000"}})
	}
	return array, nil
}
func (s *freightRepoMoch) Add(ctx context.Context, freight domain.AddFreight) error {
	return nil
}

// city contact container
type cityRepoMoch struct {
}

func (r *cityRepoMoch) Add(_ context.Context, city domain.AddCity) error {
	return nil
}
func (r *cityRepoMoch) GetAll(ctx context.Context) ([]*domain.City, error) {
	var cities []*domain.City
	return cities, nil
}

type contactRepoMoch struct {
}

func (r *contactRepoMoch) Add(ctx context.Context, contact domain.Contact) error {
	return nil
}

func (r *contactRepoMoch) GetAll(ctx context.Context) ([]*domain.Contact, error) {
	var contacts []*domain.Contact
	return contacts, nil
}

type containerRepoMoch struct {
}

func (r *containerRepoMoch) Add(ctx context.Context, containerType string) error {
	return nil
}
func (r *containerRepoMoch) GetAll(ctx context.Context) ([]*domain.Container, error) {
	var containers []*domain.Container
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

var service = NewController(&freightRepoMoch{}, &cityRepoMoch{}, &contactRepoMoch{}, &containerRepoMoch{}, loggerMoch{})

func TestGetBestFreights(t *testing.T) {
	ctx := context.Background()
	const limit = 20
	var getFreightRequestStruct = NewGetFreight("fromCity", "toCity", "40DC", limit)
	result, err := service.freightRepository.Get(ctx, getFreightRequestStruct)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, len(result), limit)
	for index, value := range result {
		assert.Equal(t, value.UsdPrice, index*1000)
	}
}
