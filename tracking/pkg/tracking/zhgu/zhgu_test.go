package zhgu

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"os"
	"strings"
	"testing"
	"time"
)

var httpMockUp = requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
	return os.ReadFile("test_data/exampleApiResponse.json")
})
var baseCfg = &tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMockUp,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

const billNo = "ZGSHA0100001921"

func getResponse() (*ApiResponseSchema, error) {
	b, err := os.ReadFile("test_data/exampleApiResponse.json")
	if err != nil {
		return nil, err
	}
	var s *ApiResponseSchema
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func testInfoAboutMoving(t *testing.T, infoAboutMoving []*tracking.Event) {
	expectedInfoAboutMoving := []*tracking.Event{
		{
			Time:          time.Date(2022, 07, 22, 0, 0, 0, 0, time.UTC),
			OperationName: "ETD",
			Location:      strings.ToTitle("SHANGHAI"),
			Vessel:        "ZHONG GU BO HAI",
		},
		{
			Time:          time.Date(2022, 07, 25, 0, 0, 0, 0, time.UTC),
			OperationName: "ATD",
			Location:      strings.ToTitle("SHANGHAI"),
			Vessel:        "ZHONG GU BO HAI",
		},
	}
	for index, v := range infoAboutMoving {
		assert.Equal(t, expectedInfoAboutMoving[index].Time, v.Time)
		assert.Equal(t, expectedInfoAboutMoving[index].OperationName, v.OperationName)
		assert.Equal(t, expectedInfoAboutMoving[index].Location, v.Location)
		assert.Equal(t, expectedInfoAboutMoving[index].Vessel, v.Vessel)
	}
}

func testEta(t *testing.T, eta time.Time) {
	expectedEta := time.Date(2022, 07, 31, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, expectedEta, eta)
}

func TestInfoAboutMovingParser_Get(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	response, err := getResponse()
	assert.NoError(t, err)

	parser := NewInfoAboutMovingParser(baseCfg.Datetime)

	infoAboutMoving, err := parser.Get(response)
	assert.NoError(t, err)
	testInfoAboutMoving(t, infoAboutMoving)
}

func TestEtaParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	response, err := getResponse()
	assert.NoError(t, err)

	parser := NewEtaParser(baseCfg.Datetime)

	eta, err := parser.Get(response)
	assert.NoError(t, err)

	testEta(t, eta)
}

func TestBillTracker(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	tracker := NewBillTracker(baseCfg)

	response, err := tracker.Track(context.Background(), billNo)
	assert.NoError(t, err)
	assert.Equal(t, response.Number, billNo)
	assert.Equal(t, response.Scac, "ZHGU")
	testInfoAboutMoving(t, response.InfoAboutMoving)
	testEta(t, response.Eta)
}
