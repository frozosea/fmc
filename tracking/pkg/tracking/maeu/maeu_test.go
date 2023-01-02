package maeu

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"os"
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

const containerNumber = "MSKU6874333"

func getResponse(t *testing.T) *ApiResponse {
	if !testing.Short() {
		t.Skip()
	}

	resp, err := httpMockUp.Do(ctx)
	assert.NoError(t, err)
	var s *ApiResponse
	err = json.Unmarshal(resp.Body, &s)
	assert.NoError(t, err)
	return s
}

func testContainerSize(t *testing.T, size string) {
	const expectedContainerSize = "40DRY"
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
	expectedEta := time.Date(2022, 6, 11, 13, 54, 0, 0, time.UTC)

	assert.Equal(t, expectedEta, event.Time)
	assert.Equal(t, "SPARTANBURG", event.Location)
	assert.Equal(t, "ETA", event.OperationName)
	assert.Equal(t, "", event.Vessel)
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

func TestPodParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	parser := NewPodParser()
	pod := parser.get(getResponse(t))
	assert.Equal(t, "SPARTANBURG", pod)

}

func testInfoAboutMoving(t *testing.T, infoAboutMoving []*tracking.Event) {
	var expectedInfoAboutMoving = []*tracking.Event{
		{
			Time:          time.Date(2022, 3, 29, 16, 42, 0, 0, time.UTC),
			OperationName: "GATE-OUT-EMPTY",
			Location:      "WIN WIN CONTAINER DEPOT, THAILAND",
			Vessel:        "MSC SVEVA",
		},
		{
			Time:          time.Date(2022, 3, 30, 10, 2, 0, 0, time.UTC),
			OperationName: "GATE-IN",
			Location:      "LAEM CHABANG TERMINAL PORT D1, THAILAND",
			Vessel:        "MSC SVEVA",
		},
		{
			Time:          time.Date(2022, 04, 14, 0, 15, 0, 0, time.UTC),
			OperationName: "LOAD",
			Location:      "LAEM CHABANG TERMINAL PORT D1, THAILAND",
			Vessel:        "MSC SVEVA",
		},
		{
			Time:          time.Date(2022, 04, 25, 12, 49, 0, 0, time.UTC),
			OperationName: "DISCHARG",
			Location:      "YANGSHAN SGH GUANDONG TERMINAL, CHINA",
			Vessel:        "MSC SVEVA",
		},
		{
			Time:          time.Date(2022, 05, 02, 1, 11, 0, 0, time.UTC),
			OperationName: "LOAD",
			Location:      "YANGSHAN SGH GUANDONG TERMINAL, CHINA",
			Vessel:        "ZIM WILMINGTON",
		},
		{
			Time:          time.Date(2022, 06, 11, 13, 54, 0, 0, time.UTC),
			OperationName: "DISCHARG",
			Location:      "CHARLESTON WANDO WELCH TERMINAL N59, UNITED STATES",
			Vessel:        "ZIM WILMINGTON",
		},
		{
			Time:          time.Date(2022, 06, 14, 10, 0, 0, 0, time.UTC),
			OperationName: "GATE-OUT",
			Location:      "CHARLESTON WANDO WELCH TERMINAL N59, UNITED STATES",
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

func TestInfoAboutMoving(t *testing.T) {
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

	assert.Equal(t, "MAEU", response.Scac)
	assert.Equal(t, containerNumber, response.Number)

	infoAboutMoving := response.InfoAboutMoving
	lastIndex := len(response.InfoAboutMoving) - 1

	testEta(t, infoAboutMoving[lastIndex])
	testContainerSize(t, response.Size)
	testInfoAboutMoving(t, infoAboutMoving[:lastIndex])
}
