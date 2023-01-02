package feso

import (
	"context"
	"github.com/pkg/profile"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

var httpMockUp = requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
	return os.ReadFile("test_data/exampleApiResponse.json")
})
var ctx = context.Background()
var requestMockUp = NewFesoRequest(httpMockUp, requests.NewUserAgentGeneratorMockUp())
var baseCfg = tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMockUp,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

func TestContainerSizeParser(t *testing.T) {
	defer profile.Start().Stop()
	if !testing.Short() {
		t.Skip()
	}
	containerSizeParser := NewContainerSizeParser()

	apiResponse, err := requestMockUp.Send(ctx, "")
	assert.NoError(t, err)

	const expectedContainerSize = "20DC"
	actualContainerSize := containerSizeParser.Get(apiResponse)
	assert.Equal(t, expectedContainerSize, actualContainerSize)
}
func testInfoAboutMoving(t *testing.T, infoAboutMoving []*tracking.Event) {
	const expectedInfoAboutMovingLen = 2
	assert.Equal(t, expectedInfoAboutMovingLen, len(infoAboutMoving))

	assert.Equal(t, infoAboutMoving[0].Time, time.Date(2022, 6, 6, 16, 0, 0, 0, time.UTC))
	assert.Equal(t, "MAGISTRAL", infoAboutMoving[0].Location)
	assert.Equal(t, "GATE OUT EMPTY FOR LOADING", infoAboutMoving[0].OperationName)
	assert.Equal(t, "", infoAboutMoving[0].Vessel)

	assert.Equal(t, infoAboutMoving[1].Time, time.Date(2022, 6, 7, 18, 0, 0, 0, time.UTC))
	assert.Equal(t, "ZAPSIBCONT", infoAboutMoving[1].Location)
	assert.Equal(t, "GATE IN EMPTY FROM CONSIGNEE", infoAboutMoving[1].OperationName)
	assert.Equal(t, "", infoAboutMoving[1].Vessel)
}

func TestInfoAboutMovingParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	infoAboutMovingParser := NewInfoAboutMovingParser(datetime.NewDatetime())

	apiResponse, err := requestMockUp.Send(ctx, "")
	assert.NoError(t, err)

	infoAboutMoving := infoAboutMovingParser.get(apiResponse.LastEvents)

	testInfoAboutMoving(t, infoAboutMoving)
	defer pprof.StopCPUProfile()
}

func TestContainerTrackerWithMockUp(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	tracker := NewContainerTracker(&baseCfg)
	response, err := tracker.Track(ctx, "")
	assert.NoError(t, err)

	testInfoAboutMoving(t, response.InfoAboutMoving)

	assert.Equal(t, "20DC", response.Size)
	assert.Equal(t, "FESO", response.Scac)
}
