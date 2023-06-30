package sklu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"os"
	"testing"
	"time"
)

type unlocodeRepoMockUp struct {
}

func newUnlocodeRepoMockUp() *unlocodeRepoMockUp {
	return &unlocodeRepoMockUp{}
}

func (u unlocodeRepoMockUp) GetFullName(_ context.Context, _ string) (string, error) {
	return "", nil
}

func (u unlocodeRepoMockUp) Add(_ context.Context, _, _ string) error {
	return nil
}

var httpMockUp = requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
	if fmt.Sprintf("http://ebiz.sinokor.co.kr/Tracking/GetBLList?cntrno=%s&year=%d", "SKLU1327134", time.Now().Year()) == r.RUrl {
		return os.ReadFile("test_data/exampleApiResponse.json")
	} else if r.RUrl == fmt.Sprintf("http://ebiz.sinokor.co.kr/Tracking/GetBLList?cntrno=%s&year=%d", containerNumber, time.Now().Year()) {
		return os.ReadFile("test_data/exampleApiResponse.json")
	} else if r.RUrl == fmt.Sprintf(`http://ebiz.sinokor.co.kr/Tracking?blno=%s&cntrno=`, billNumber) {
		return os.ReadFile("test_data/exampleInfoAboutMoving.txt")
	} else if r.RUrl == fmt.Sprintf(`http://ebiz.sinokor.co.kr/Tracking?blno=%s&cntrno=%s`, billNumber, containerNumber) {
		return os.ReadFile("test_data/exampleInfoAboutMoving.txt")
	} else if r.RUrl == fmt.Sprintf("http://ebiz.sinokor.co.kr/Home/chkExistsBooking?bkno=%s", billNumber) {
		return json.Marshal(&CheckBookingNumberExistsResponse{STATUS: "Y", MSG: ""})
	} else {
		return os.ReadFile("test_data/exampleInfoAboutMoving.txt")
	}
})
var ctx = context.Background()

const containerNumber = "TEMU2094051"
const billNumber = "SNKO101220501450"

var baseCfg = &tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMockUp,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

func getApiResponse(t *testing.T) *ApiResponse {
	r, err := os.ReadFile("test_data/exampleApiResponse.json")
	assert.NoError(t, err)
	var s []*ApiResponse
	if err := json.Unmarshal(r, &s); err != nil {
		return nil
	}
	return s[0]
}

func testContainerSize(t *testing.T, size string) {
	const expectedContainerSize = "20'X1"
	assert.Equal(t, expectedContainerSize, size)
}
func testEta(t *testing.T, eta time.Time) {
	expectedEta := time.Date(2022, 06, 17, 8, 0, 0, 0, time.UTC)
	assert.Equal(t, expectedEta, eta)
}
func TestApiParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	parser := NewApiParser(baseCfg.Datetime)

	d := parser.Get(getApiResponse(t))

	testContainerSize(t, d.ContainerSize)
	assert.Equal(t, d.Eta, time.Date(2022, time.November, 6, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, "RUVYP", d.Unlocode)
}

func getInfoAboutMovingDoc(t *testing.T) *goquery.Document {
	response, err := os.ReadFile("test_data/exampleInfoAboutMoving.txt")
	assert.NoError(t, err)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response))
	if err != nil {
		return nil
	}
	return doc
}

func testInfoAboutMoving(t *testing.T, infoAboutMoving []*tracking.Event) {
	var expectedInfoAboutMoving = []*tracking.Event{
		{
			Time:          time.Date(2022, 05, 30, 10, 22, 0, 0, time.UTC),
			OperationName: "PICKUP (1/1)",
			Location:      "SINOKOR TAM CANG CAT LAI DEPOT",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 05, 31, 01, 48, 0, 0, time.UTC),
			OperationName: "RETURN (1/1)",
			Location:      "CAT LAI",
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 07, 2, 00, 0, 0, time.UTC),
			OperationName: "DEPARTURE",
			Location:      "CAT LAI",
			Vessel:        "HEUNG-A HOCHIMINH / 2205N",
		},
		{
			Time:          time.Date(2022, 06, 13, 2, 00, 0, 0, time.UTC),
			OperationName: "ARRIVAL(T/S) (SCHEDULED)",
			Location:      "BPTS",
			Vessel:        "HEUNG-A HOCHIMINH / 2205N",
		},
		{
			Time:          time.Date(2022, 06, 15, 9, 00, 0, 0, time.UTC),
			OperationName: "DEPARTURE(T/S) (SCHEDULED)",
			Location:      "BPTS",
			Vessel:        "HEUNG-A ULSAN / 2256E",
		},
		{
			Time:          time.Date(2022, 06, 17, 8, 00, 0, 0, time.UTC),
			OperationName: "ARRIVAL (SCHEDULED)",
			Location:      "HOSOSHIMA TERMINAL(SHIRAHMA #14)",
			Vessel:        "HEUNG-A ULSAN / 2256E",
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

	infoAboutMoving, err := parser.Get(getInfoAboutMovingDoc(t), containerNumber)
	assert.NoError(t, err)
	testInfoAboutMoving(t, infoAboutMoving)
}

func TestContainerTracker(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	tracker := NewContainerTracker(baseCfg, newUnlocodeRepoMockUp())
	response, err := tracker.Track(ctx, containerNumber)
	assert.NoError(t, err)

	assert.Equal(t, response.Number, containerNumber)
	assert.Equal(t, "SKLU", response.Scac)
	testContainerSize(t, response.Size)
	testInfoAboutMoving(t, response.InfoAboutMoving[:len(response.InfoAboutMoving)-1])
	assert.Equal(t, response.InfoAboutMoving[len(response.InfoAboutMoving)-1].Time, time.Date(2022, time.November, 6, 0, 0, 0, 0, time.UTC))
}

func TestBillNumberTracker(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	tracker := NewBillTracker(baseCfg)

	response, err := tracker.Track(ctx, billNumber)
	assert.NoError(t, err)

	assert.Equal(t, billNumber, response.Number)
	assert.Equal(t, "SKLU", response.Scac)

	testEta(t, response.Eta)
	for _, item := range response.InfoAboutMoving {
		if item.Vessel == containerNumber {
			item.Vessel = ""
		}
	}
	testInfoAboutMoving(t, response.InfoAboutMoving)
}
