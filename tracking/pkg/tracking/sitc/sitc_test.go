package sitc

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"golang_tracking/pkg/tracking/util/sitc/captcha_resolver"
	"golang_tracking/pkg/tracking/util/sitc/login_provider"
	"os"
	"strings"
	"testing"
	"time"
)

var httpMockUp = requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
	return os.ReadFile("test_data/containerExampleResponse.json")
})
var ctx = context.Background()
var baseCfg = tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMockUp,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

const containerNumber = `SITU9130070`
const billNumber = `SITDLVK222G951`

type BillRequestMockUp struct {
}

func NewBillRequestMockUp() *BillRequestMockUp {
	return &BillRequestMockUp{}
}

func (b *BillRequestMockUp) GetBillNumberInfo(_ context.Context, _, _, _ string) (*BillNumberApiResponse, error) {
	billNumberExampleResponse, err := os.ReadFile("test_data/billNumberExampleResponse.json")
	if err != nil {
		return nil, err
	}
	var s *BillNumberApiResponse
	if err := json.Unmarshal(billNumberExampleResponse, &s); err != nil {
		return nil, err
	}
	return s, nil
}
func (b *BillRequestMockUp) GetContainerInfo(_ context.Context, _, _ string) (*BillNumberInfoAboutContainerApiResponse, error) {
	containerMovementInfoExampleResponse, err := os.ReadFile("test_data/containerMovementInfoExampleResponse.json")
	if err != nil {
		return nil, err
	}
	var s *BillNumberInfoAboutContainerApiResponse
	if err := json.Unmarshal(containerMovementInfoExampleResponse, &s); err != nil {
		return nil, err
	}
	return s, nil
}

type captchaResolverMockUp struct {
}

func newCaptchaResolverMockUp() *captchaResolverMockUp {
	return &captchaResolverMockUp{}
}

func (c *captchaResolverMockUp) Resolve(_ context.Context) (captcha_resolver.RandomString, captcha_resolver.SolvedCaptcha, error) {
	return "", "", nil
}
func testContainerInfoAboutMoving(t *testing.T, infoAboutMoving []*tracking.Event) {
	var expectedInfoAboutMoving = []*tracking.Event{
		{
			//2022-05-18 18:53:00
			Time:          time.Date(2022, 05, 18, 18, 53, 0, 0, time.UTC),
			OperationName: "EMPTY CONTAINER",
			Location:      "DALIAN",
			Vessel:        "SITC MAKASSAR",
		},
		{
			//2022-05-27 10:59:00
			Time:          time.Date(2022, 05, 27, 10, 59, 0, 0, time.UTC),
			OperationName: "OUTBOUND PICKUP",
			Location:      "DALIAN",
			Vessel:        "SITC CAGAYAN",
		},
		{
			//2022-06-04 22:00:00
			Time:          time.Date(2022, 06, 04, 22, 00, 0, 0, time.UTC),
			OperationName: "LOADED ONTO VESSEL",
			Location:      "DALIAN",
			Vessel:        "SITC CAGAYAN",
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

func TestContainerInfoAboutMoving(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	containerExampleResponse, err := os.ReadFile("test_data/containerExampleResponse.json")
	assert.NoError(t, err)

	var s *ContainerApiResponse
	if err := json.Unmarshal(containerExampleResponse, &s); err != nil {
		panic(err)
		return
	}

	parser := NewContainerTrackingInfoAboutMovingParser(baseCfg.Datetime)

	infoAboutMoving := parser.get(s)
	testContainerInfoAboutMoving(t, infoAboutMoving)
}

func TestContainerTrackerWithMockUp(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	tracker := NewContainerTracker(&baseCfg, login_provider.NewStore("username", "password", "auth", "authkey"))

	response, err := tracker.Track(ctx, containerNumber)
	assert.NoError(t, err)

	assert.Equal(t, response.Scac, "SITC")
	assert.Equal(t, response.Size, "")
	assert.Equal(t, response.Number, containerNumber)
	testContainerInfoAboutMoving(t, response.InfoAboutMoving)
}

func testEtaParser(t *testing.T, eta time.Time) {
	assert.Equal(t, time.Date(2022, time.July, 5, 21, 0, 0, 0, time.UTC), eta)
}

func TestEtaParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	billNumberExampleResponse, err := os.ReadFile("test_data/billNumberExampleResponse.json")
	assert.NoError(t, err)

	var s *BillNumberApiResponse
	if err := json.Unmarshal(billNumberExampleResponse, &s); err != nil {
		panic(err)
		return
	}

	parser := newEtaParser(baseCfg.Datetime)
	response, err := parser.get(s)
	assert.NoError(t, err)
	testEtaParser(t, response)
}

func testBillNumberInfoAboutMovingParser(t *testing.T, infoAboutMoving []*tracking.Event) {

	var expectedInfoAboutMoving = []*tracking.Event{
		{
			Time:          time.Date(2022, 05, 27, 0, 0, 0, 0, time.UTC),
			OperationName: strings.ToTitle("outbound pickup"),
			Location:      strings.ToUpper("dalian"),
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 04, 0, 0, 0, 0, time.UTC),
			OperationName: strings.ToTitle("loaded onto vessel"),
			Location:      strings.ToUpper("dalian"),
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 13, 0, 0, 0, 0, time.UTC),
			OperationName: strings.ToTitle("inbound in cy"),
			Location:      strings.ToUpper("shanghai"),
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 06, 13, 0, 0, 0, 0, time.UTC),
			OperationName: strings.ToTitle("transshipment between vessel"),
			Location:      strings.ToUpper("shanghai"),
			Vessel:        "",
		},
		{
			Time:          time.Date(2022, 07, 03, 0, 0, 0, 0, time.UTC),
			OperationName: strings.ToTitle("loaded onto vessel"),
			Location:      strings.ToUpper("shanghai"),
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

func TestBillNumberInfoAboutMovingParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	containerMovementInfoExampleResponse, err := os.ReadFile("test_data/containerMovementInfoExampleResponse.json")
	assert.NoError(t, err)

	var s *BillNumberInfoAboutContainerApiResponse
	if err := json.Unmarshal(containerMovementInfoExampleResponse, &s); err != nil {
		panic(err)
		return
	}

	parser := newBillNumberInfoAboutMovingParser(baseCfg.Datetime)

	infoAboutMoving := parser.get(s)

	testBillNumberInfoAboutMovingParser(t, infoAboutMoving)
}

func TestBillNumberTracker_TrackByBillNumber(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	tracker := NewBillTracker(&baseCfg, NewBillRequestMockUp(), newCaptchaResolverMockUp())

	response, err := tracker.Track(ctx, billNumber)
	assert.NoError(t, err)

	assert.Equal(t, billNumber, response.Number)
	assert.Equal(t, "SITC", response.Scac)
	testEtaParser(t, response.Eta)
	testBillNumberInfoAboutMovingParser(t, response.InfoAboutMoving)
}
