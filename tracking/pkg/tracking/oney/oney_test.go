package oney

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
	if r.RQuery["f_cmd"] == "122" {
		s := BkgAndCopNosApiResponseSchema{
			BaseApiResponseEntity: &BaseApiResponseEntity{
				TRANSRESULTKEY: "",
				Exception:      "",
				Count:          "",
			},
			List: []*struct {
				MaxRows     int             `json:"maxRows"`
				Models      []interface{}   `json:"models"`
				CopNo       string          `json:"copNo"`
				BkgNo       string          `json:"bkgNo"`
				CntrNo      string          `json:"cntrNo"`
				EnblFlag    string          `json:"enblFlag"`
				HashColumns [][]interface{} `json:"hashColumns"`
				HashFields  []interface{}   `json:"hashFields"`
			}{{
				1,
				[]interface{}{},
				"CSEL2303645419",
				"CSEL2303645419",
				"CntrNo",
				"EnblFlag",
				nil,
				nil,
			}},
		}
		return json.Marshal(s)
	} else if r.RQuery["f_cmd"] == "125" {
		return os.ReadFile("test_data/exampleInfoAboutMovingResponse.json")

	} else if r.RQuery["f_cmd"] == "123" {
		return os.ReadFile("test_data/exampleContainerInfoResponse.json")
	} else if r.RQuery["_search"] == "false" {
		s := &BaseApiResponseEntity{
			TRANSRESULTKEY: "2",
			Exception:      "",
			Count:          "2",
		}
		return json.Marshal(s)
	}
	return os.ReadFile("test_data/exampleInfoAboutMovingResponse.json")
})
var ctx = context.Background()
var baseCfg = tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMockUp,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

const containerNumber = "CSEL2303645419"

func testContainerSize(t *testing.T, size string) {
	const expectedContainerSize = "40'DRY HC."
	assert.Equal(t, expectedContainerSize, size)
}

func TestContainerSizeParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	r, err := os.ReadFile("test_data/exampleContainerInfoResponse.json")
	assert.NoError(t, err)

	var s *ContainerSizeApiResponseSchema
	if err := json.Unmarshal(r, &s); err != nil {
		panic(err)
		return
	}

	parser := NewContainerSizeParser()
	size := parser.get(s)

	testContainerSize(t, size)
}

func testInfoAboutMoving(t *testing.T, infoAboutMoving []*tracking.Event) {
	var expectedInfoAboutMoving = []*tracking.Event{
		{
			Time:          time.Date(2022, 03, 24, 11, 19, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("EMPTY CONTAINER RELEASE TO SHIPPER")),
			Location:      "PUSAN, KOREA REPUBLIC OF",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 04, 05, 10, 40, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Gate In to Outbound Terminal")),
			Location:      "PUSAN, KOREA REPUBLIC OF",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 04, 07, 12, 9, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Loaded on 'HYUNDAI SINGAPORE 126E' at Port of Loading")),
			Location:      "PUSAN, KOREA REPUBLIC OF",
			Vessel:        "HYUNDAI SINGAPORE",
		},
		{
			Time:          time.Date(2022, 04, 07, 21, 00, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("'HYUNDAI SINGAPORE 126E' Departure from Port of Loading")),
			Location:      "PUSAN, KOREA REPUBLIC OF",
			Vessel:        "HYUNDAI SINGAPORE",
		},
		{
			Time:          time.Date(2022, 04, 18, 16, 00, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("'HYUNDAI SINGAPORE 126E' Arrival at Port of Discharging")),
			Location:      "VANCOUVER, BC, CANADA",
			Vessel:        "HYUNDAI SINGAPORE",
		},
		{
			Time:          time.Date(2022, 05, 28, 03, 14, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("'HYUNDAI SINGAPORE 126E' POD Berthing Destination")),
			Location:      "VANCOUVER, BC, CANADA",
			Vessel:        "HYUNDAI SINGAPORE",
		},
		{
			Time:          time.Date(2022, 05, 31, 21, 47, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Unloaded from 'HYUNDAI SINGAPORE 126E' at Port of Discharging")),
			Location:      "VANCOUVER, BC, CANADA",
			Vessel:        "HYUNDAI SINGAPORE",
		},
		{
			Time:          time.Date(2022, 06, 03, 00, 24, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Loaded on rail at inbound rail origin")),
			Location:      "VANCOUVER, BC, CANADA",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 03, 07, 9, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Inbound Rail Departure")),
			Location:      "VANCOUVER, BC, CANADA",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 9, 18, 34, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Inbound Rail Arrival")),
			Location:      "DETROIT, MI, UNITED STATES",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 9, 21, 2, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Unloaded from rail at inbound rail destination")),
			Location:      "DETROIT, MI, UNITED STATES",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 10, 10, 38, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Gate Out from Inbound CY for Delivery to Consignee")),
			Location:      "DETROIT, MI, UNITED STATES",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 10, 12, 24, 0, 0, time.UTC),
			OperationName: strings.ToTitle(strings.ToLower("Empty Container Returned from Customer")),
			Location:      "DETROIT, MI, UNITED STATES",
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
	exampleInfoAboutMovingResponse, err := os.ReadFile("test_data/exampleInfoAboutMovingResponse.json")
	assert.NoError(t, err)
	var s *InfoAboutMovingApiResponseSchema

	if err := json.Unmarshal(exampleInfoAboutMovingResponse, &s); err != nil {
		panic(err)
	}

	parser := NewInfoAboutMovingParser(baseCfg.Datetime)

	infoAboutMoving := parser.get(s)
	testInfoAboutMoving(t, infoAboutMoving)
}

func TestContainerTrackerWithMockUp(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	tracker := NewContainerTracker(&baseCfg)

	response, err := tracker.Track(ctx, containerNumber)
	assert.NoError(t, err)

	assert.Equal(t, response.Number, containerNumber)
	assert.Equal(t, response.Scac, "ONEY")

	testContainerSize(t, response.Size)
	testInfoAboutMoving(t, response.InfoAboutMoving)
}
