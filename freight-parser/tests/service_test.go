package tests

import (
	freightPackage "fmc-newest/pkg/freight"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type repositoryMoch struct {
}

func (s repositoryMoch) GetFrieght(freight freightPackage.GetFreight) ([]freightPackage.BaseFreight, error) {
	var array []freightPackage.BaseFreight
	for i := 0; i < freight.Limit; i++ {
		array = append(array, freightPackage.BaseFreight{FromCity: "fromCity", ToCity: "toCity", ContainerType: "40DC", UsdPrice: i * 1000, Line: "FESO", LineImage: "lineImage", Contacts: freightPackage.Contact{Url: "http://localhost", AgentName: "his name", PhoneNumber: "1654783200000"}})
	}
	return array, nil
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

var service = freightPackage.NewService(repositoryMoch{}, loggerMoch{})

func TestGetBestFreights(t *testing.T) {
	const limit = 20
	var getFreightRequestStruct = freightPackage.NewGetFreight("fromCity", "toCity", 2, limit)
	result, err := service.GetBestFreights(getFreightRequestStruct)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Equal(t, len(result), limit)
	for index, value := range result {
		assert.Equal(t, value.UsdPrice, index*1000)
	}
}
func TestGetAllFreights(t *testing.T) {
	const limit = 20
	var getFreightRequestStruct = freightPackage.NewGetFreight("fromCityGetAllFreights", "toCityGetAllFreights", 3, limit)
	result, err := service.GetBestFreights(getFreightRequestStruct)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, len(result), limit)
	for index, value := range result {
		assert.Equal(t, value.UsdPrice, index*1000)
	}
}
