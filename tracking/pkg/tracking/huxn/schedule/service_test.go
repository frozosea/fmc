package huaxin_schedule

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/huxn"
	"golang_tracking/pkg/tracking/huxn/schedule/unlocodesParser"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"os"
	"testing"
	"time"
)

var httpMock = requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
	return os.ReadFile("test_data/schedule.json")
})

func getTrackingResponse() *huxn.ContainerTrackingResponse {
	data, err := os.ReadFile("test_data/tracking_response.json")
	if err != nil {
		panic(err)
	}
	var t *huxn.ContainerTrackingResponse

	if err := json.Unmarshal(data, &t); err != nil {
		panic(err)
	}

	return t
}

func getSchedule() []*Schedule {
	data, err := os.ReadFile("test_data/schedule.json")
	if err != nil {
		panic(err)
	}
	var t *ServerResponse
	if err := json.Unmarshal(data, &t); err != nil {
		panic(err)
	}
	return t.ListSchedules
}

var trackingResponse = getTrackingResponse()
var schedules = getSchedule()

var cfg = tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMock,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

func TestParser(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	expectedData := &DataForScheduleRequest{
		lastVoyage:       "23016N",
		lastPortUnlocode: "CNTAC",
		//2023/5/29 15:53:23
		etd: time.Date(2023, 05, 29, 15, 53, 23, 0, time.UTC),
	}
	p := NewParser(cfg.Datetime)

	data, err := p.GetDataForScheduleRequest(trackingResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, data)
}

func TestScheduleParser(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	expectedData := &tracking.Event{
		Time:          time.Date(2023, 06, 9, 0, 0, 0, 0, time.UTC),
		OperationName: "ETA",
		Location:      "VLADIVOSTOK SOLLERS TERMINAL",
		Vessel:        "CHANG RONG 8",
	}

	p := NewScheduleParser(cfg.Datetime)
	data, err := p.GetETA("23016N", schedules)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, data)
}

func TestService(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	s := NewService(unlocodesParser.NewServiceMockup(), cfg.Request, cfg.UserAgentGenerator, cfg.Datetime)

	expectedData := &tracking.Event{
		Time:          time.Date(2023, 06, 9, 0, 0, 0, 0, time.UTC),
		OperationName: "ETA",
		Location:      "VLADIVOSTOK SOLLERS TERMINAL",
		Vessel:        "CHANG RONG 8",
	}

	data, err := s.GetETA(context.Background(), trackingResponse)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, data)
}
