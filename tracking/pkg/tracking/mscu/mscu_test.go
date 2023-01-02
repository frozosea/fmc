package mscu

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
var ctx = context.Background()
var baseCfg = tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMockUp,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

const containerNumber = "MEDU3170580"

func getResponse(t *testing.T) *ApiResponse {
	resp, err := httpMockUp.Do(ctx)
	assert.NoError(t, err)
	var s *ApiResponse
	err = json.Unmarshal(resp.Body, &s)
	assert.NoError(t, err)
	return s
}

func testContainerSize(t *testing.T, size string) {
	const expectedContainerSize = "20' DRY VAN"
	assert.Equal(t, expectedContainerSize, size)
}

func TestContainerSizeParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	parser := NewContainerSizeParser()

	size := parser.get(getResponse(t))
	testContainerSize(t, size)
}

func testEta(t *testing.T, event *tracking.Event) {
	expectedEta := time.Date(2022, 6, 12, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, expectedEta, event.Time)
	assert.Equal(t, event.OperationName, "ETA")
	assert.Equal(t, event.Location, "")
	assert.Equal(t, event.Vessel, "")
}

func TestEtaParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	parser := NewEtaParser(baseCfg.Datetime)

	event, err := parser.get(getResponse(t))
	assert.NoError(t, err)
	testEta(t, event)
}

func testInfoAboutMoving(t *testing.T, infoAboutMoving []*tracking.Event) {
	expectedInfoAboutMoving := []*tracking.Event{
		{
			Time:          time.Date(2022, 6, 9, 0, 0, 0, 0, time.UTC),
			OperationName: strings.ToUpper("Empty to Shipper"),
			Location:      "CHONGQING, CN",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 10, 0, 0, 0, 0, time.UTC),
			OperationName: strings.ToUpper("Export at barge yard"),
			Location:      "CHONGQING, CN",
			Vessel:        "",
		},
	}
	assert.Equal(t, len(expectedInfoAboutMoving), len(infoAboutMoving))

	for index, item := range infoAboutMoving {
		assert.Equal(t, expectedInfoAboutMoving[index].Time, item.Time)
		assert.Equal(t, expectedInfoAboutMoving[index].OperationName, item.OperationName)
		assert.Equal(t, expectedInfoAboutMoving[index].Location, item.Location)
		assert.Equal(t, expectedInfoAboutMoving[index].Vessel, item.Vessel)
	}
}

func TestInfoAboutMovingParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	parser := NewInfoAboutMovingParser(baseCfg.Datetime)

	infoAboutMoving := parser.get(getResponse(t))
	testInfoAboutMoving(t, infoAboutMoving)
}

func TestContainerTrackerWithMockUp(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	tracker := NewContainerTracker(&baseCfg)

	response, err := tracker.Track(ctx, containerNumber)
	assert.NoError(t, err)

	assert.Equal(t, containerNumber, response.Number)
	assert.Equal(t, "MSCU", response.Scac)

	testContainerSize(t, response.Size)

	infoAboutMoving := response.InfoAboutMoving
	lastIndex := len(infoAboutMoving) - 1

	testEta(t, infoAboutMoving[lastIndex])
	testInfoAboutMoving(t, infoAboutMoving[:lastIndex])
}
