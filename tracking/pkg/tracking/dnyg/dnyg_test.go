package dnyg

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"os"
	"testing"
	"time"
)

func getInfoAboutMoving() *InfoAboutMovingResponse {
	bytes, err := os.ReadFile("test_data/info_about_moving.json")
	if err != nil {
		panic(err)
	}

	var rawInfoAboutMoving *InfoAboutMovingResponse

	if err := json.Unmarshal(bytes, &rawInfoAboutMoving); err != nil {
		panic(err)
	}
	return rawInfoAboutMoving
}

func getNumberInfo() *NumberInfoResponse {
	bytes, err := os.ReadFile("test_data/container_info.json")
	if err != nil {
		panic(err)
	}

	var numberInfo *NumberInfoResponse

	if err := json.Unmarshal(bytes, &numberInfo); err != nil {
		panic(err)
	}
	return numberInfo
}

var httpMockup = requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
	if r.RUrl == "https://ebiz.pcsline.co.kr/trk/trkE0710R01N" {
		return os.ReadFile("test_data/container_info.json")
	} else if r.RUrl == "https://ebiz.pcsline.co.kr/trk/trkE0710R03" {
		return os.ReadFile("test_data/info_about_moving.json")
	}
	return nil, errors.New("invalid url")
})

var baseCfg = &tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMockup,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

func TestInfoAboutMoving(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	rawInfoAboutMoving := getInfoAboutMoving()

	expectedData := []*tracking.Event{
		{
			Time:          time.Date(2023, 05, 26, 17, 07, 0, 0, time.UTC),
			OperationName: "SHIPPER FULL = EXPORT EMPTY GATE OUT",
			Location:      "NAGOYA, JAPAN",
			Vessel:        "",
		},
		{
			Time:          time.Date(2023, 06, 01, 15, 47, 0, 0, time.UTC),
			OperationName: "LOADING FULL = EXPORT FULL IN TERMINAL",
			Location:      "NAGOYA, JAPAN",
			Vessel:        "",
		},
		{
			Time:          time.Date(2023, 06, 04, 21, 25, 0, 0, time.UTC),
			OperationName: "VESSEL FULL LOADING",
			Location:      "NAGOYA, JAPAN",
			Vessel:        "DONGJIN VENUS",
		},
		{
			Time:          time.Date(2023, 06, 06, 9, 00, 0, 0, time.UTC),
			OperationName: "DISCHARGING FULL = IMPORT FULL TERMINAL",
			Location:      "BUSAN, KOREA",
			Vessel:        "DONGJIN VENUS",
		},
		{
			Time:          time.Date(2023, 06, 8, 3, 30, 0, 0, time.UTC),
			OperationName: "VESSEL FULL LOADING",
			Location:      "BUSAN, KOREA",
			Vessel:        "XIANG REN",
		},
		{
			Time:          time.Date(2023, 06, 9, 23, 30, 0, 0, time.UTC),
			OperationName: "DISCHARGING FULL = IMPORT FULL TERMINAL",
			Location:      "VLADIVOSTOK, RUSSIA",
			Vessel:        "XIANG REN",
		},
	}

	infoAboutMovingParser := NewInfoAboutMovingParser(datetime.NewDatetime())
	infoAboutMoving := infoAboutMovingParser.Parse(rawInfoAboutMoving)
	assert.Equal(t, expectedData, infoAboutMoving)

}

func TestEtaParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	etaParser := NewEtaParser(datetime.NewDatetime())
	//2023-06-10 22:00
	expectedEta := time.Date(2023, 06, 10, 22, 00, 00, 00, time.UTC)

	containerInfo := getNumberInfo()

	eta, err := etaParser.parse(containerInfo)

	assert.NoError(t, err)
	assert.Equal(t, expectedEta, eta)

}

func TestPodParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	podParser := NewPODParser()
	//VLADIVOSTOK, RUSSIA <br> 2023-06-10 22:00
	expectedPod := "VLADIVOSTOK, RUSSIA"

	containerInfo := getNumberInfo()

	pod := podParser.parse(containerInfo)

	assert.Equal(t, expectedPod, pod)
}

func TestContainerTracker(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	containerTracker := NewContainerTracker(baseCfg)

	const containerNumber = "DYLU5112648"

	expectedData := &tracking.ContainerTrackingResponse{
		Number: containerNumber,
		Size:   "4HDC",
		Scac:   "DNYG",
		InfoAboutMoving: []*tracking.Event{
			{
				Time:          time.Date(2023, 05, 26, 17, 07, 0, 0, time.UTC),
				OperationName: "SHIPPER FULL = EXPORT EMPTY GATE OUT",
				Location:      "NAGOYA, JAPAN",
				Vessel:        "",
			},
			{
				Time:          time.Date(2023, 06, 01, 15, 47, 0, 0, time.UTC),
				OperationName: "LOADING FULL = EXPORT FULL IN TERMINAL",
				Location:      "NAGOYA, JAPAN",
				Vessel:        "",
			},
			{
				Time:          time.Date(2023, 06, 04, 21, 25, 0, 0, time.UTC),
				OperationName: "VESSEL FULL LOADING",
				Location:      "NAGOYA, JAPAN",
				Vessel:        "DONGJIN VENUS",
			},
			{
				Time:          time.Date(2023, 06, 06, 9, 00, 0, 0, time.UTC),
				OperationName: "DISCHARGING FULL = IMPORT FULL TERMINAL",
				Location:      "BUSAN, KOREA",
				Vessel:        "DONGJIN VENUS",
			},
			{
				Time:          time.Date(2023, 06, 8, 3, 30, 0, 0, time.UTC),
				OperationName: "VESSEL FULL LOADING",
				Location:      "BUSAN, KOREA",
				Vessel:        "XIANG REN",
			},
			{
				Time:          time.Date(2023, 06, 9, 23, 30, 0, 0, time.UTC),
				OperationName: "DISCHARGING FULL = IMPORT FULL TERMINAL",
				Location:      "VLADIVOSTOK, RUSSIA",
				Vessel:        "XIANG REN",
			},
			{
				Time:          time.Date(2023, 06, 10, 22, 00, 00, 00, time.UTC),
				OperationName: "ETA",
				Location:      "VLADIVOSTOK, RUSSIA",
				Vessel:        "",
			},
		},
	}

	response, err := containerTracker.Track(context.Background(), containerNumber)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, response)
}

func TestBillNumberTracker(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	billTracker := NewBillTracker(baseCfg)

	const billNumber = "PCSLJNGVV23Q0047"

	expectedData := &tracking.BillNumberTrackingResponse{
		Number: billNumber,
		Eta:    time.Date(2023, 06, 10, 22, 00, 00, 00, time.UTC),
		Scac:   "DNYG",
		InfoAboutMoving: []*tracking.Event{
			{
				Time:          time.Date(2023, 05, 26, 17, 07, 0, 0, time.UTC),
				OperationName: "SHIPPER FULL = EXPORT EMPTY GATE OUT",
				Location:      "NAGOYA, JAPAN",
				Vessel:        "",
			},
			{
				Time:          time.Date(2023, 06, 01, 15, 47, 0, 0, time.UTC),
				OperationName: "LOADING FULL = EXPORT FULL IN TERMINAL",
				Location:      "NAGOYA, JAPAN",
				Vessel:        "",
			},
			{
				Time:          time.Date(2023, 06, 04, 21, 25, 0, 0, time.UTC),
				OperationName: "VESSEL FULL LOADING",
				Location:      "NAGOYA, JAPAN",
				Vessel:        "DONGJIN VENUS",
			},
			{
				Time:          time.Date(2023, 06, 06, 9, 00, 0, 0, time.UTC),
				OperationName: "DISCHARGING FULL = IMPORT FULL TERMINAL",
				Location:      "BUSAN, KOREA",
				Vessel:        "DONGJIN VENUS",
			},
			{
				Time:          time.Date(2023, 06, 8, 3, 30, 0, 0, time.UTC),
				OperationName: "VESSEL FULL LOADING",
				Location:      "BUSAN, KOREA",
				Vessel:        "XIANG REN",
			},
			{
				Time:          time.Date(2023, 06, 9, 23, 30, 0, 0, time.UTC),
				OperationName: "DISCHARGING FULL = IMPORT FULL TERMINAL",
				Location:      "VLADIVOSTOK, RUSSIA",
				Vessel:        "XIANG REN",
			},
		},
	}

	response, err := billTracker.Track(context.Background(), billNumber)
	assert.NoError(t, err)
	assert.Equal(t, expectedData, response)
}
