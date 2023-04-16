package reel

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"os"
	"strings"
	"testing"
	"time"
)

func getDoc(path string) *goquery.Document {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		panic(err)
	}
	return doc
}

var billDoc = getDoc(`test_data/bill_page.txt`)
var containerDoc = getDoc(`test_data/container_page.txt`)

func TestPodParser(t *testing.T) {

	podParser := NewPodParser()

	pod, err := podParser.Get(billDoc)
	assert.NoError(t, err)
	assert.Equal(t, "VLADIVOSTOK", pod)
}

func TestETDParser(t *testing.T) {
	etdParser := NewETDParser(datetime.NewDatetime())

	etd, err := etdParser.Get(billDoc)
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC), etd)
}

func TestContainerNumberParser(t *testing.T) {
	containerNumberParser := NewContainerNumberParser()
	number, err := containerNumberParser.Get(billDoc)
	assert.NoError(t, err)
	assert.Equal(t, `CNTR40575`, number)
}

func TestContainerStatusParser(t *testing.T) {
	containerStatusParser := NewContainerStatusParser(datetime.NewDatetime())

	data, err := containerStatusParser.Get(billDoc)
	assert.NoError(t, err)

	assert.Equal(t, &containerStatus{
		Number:    "CNTR40575",
		Type:      "HC40 - DRY",
		EventDate: time.Date(2023, 03, 27, 0, 0, 0, 0, time.UTC),
		Status:    "RECD MTY FROM CONSIGNEE",
		Location:  "VLADIVOSTOK",
	}, data)

}

func TestBillInfoAboutMoving(t *testing.T) {
	containerStatusParser := NewContainerStatusParser(datetime.NewDatetime())

	data, err := containerStatusParser.Get(billDoc)
	assert.NoError(t, err)

	parser := newBillInfoAboutMovingParser()
	event, err := parser.Get(data)
	assert.NoError(t, err)

	assert.Equal(t, &tracking.Event{
		Time:          time.Date(2023, 03, 27, 0, 0, 0, 0, time.UTC),
		OperationName: "RECD MTY FROM CONSIGNEE",
		Location:      "VLADIVOSTOK",
		Vessel:        "",
	}, event)

}

func TestInfoAboutMovingParser(t *testing.T) {
	parser := newInfoAboutMovingParser(datetime.NewDatetime())

	events, err := parser.Get(containerDoc)
	assert.NoError(t, err)

	expectedData := []*tracking.Event{
		{
			Time:          time.Date(2023, 03, 01, 13, 0, 13, 0, time.UTC),
			OperationName: "PICKUP BY SHIPPER",
			Location:      "NINGBO,CHINA(CNNGB)",
			Vessel:        "VESSEL6 / VOYAGE6",
		},
		{
			Time:          time.Date(2023, 03, 03, 15, 0, 51, 0, time.UTC),
			OperationName: "EXPORT AT PORT",
			Location:      "NINGBO,CHINA(SHIPPER PREMISES)",
			Vessel:        "VESSEL5 / VOYAGE5",
		},
		{
			Time:          time.Date(2023, 03, 05, 18, 0, 13, 0, time.UTC),
			OperationName: "LOADED FULL",
			Location:      "NINGBO,CHINA(CNNGB)",
			Vessel:        "VESSEL4 / VOYAGE4",
		},
		{
			Time:          time.Date(2023, 03, 13, 7, 0, 0, 0, time.UTC),
			OperationName: "IMPORT DISCHARGE FULL",
			Location:      "VLADIVOSTOK,RUSSIA(FROM / TO VESSEL)",
			Vessel:        "VESSEL3 / VOYAGE3",
		},
		{
			Time:          time.Date(2023, 03, 27, 9, 0, 1, 0, time.UTC),
			OperationName: "PICKUP BY CONSIGNEE",
			Location:      "VLADIVOSTOK,RUSSIA(RUVVO)",
			Vessel:        "VESSEL2 / VOYAGE2",
		},
		{
			Time:          time.Date(2023, 03, 27, 11, 0, 32, 0, time.UTC),
			OperationName: "RECD MTY FROM CONSIGNEE",
			Location:      "VLADIVOSTOK,RUSSIA(CONSIGNEE PREMISES)",
			Vessel:        "VESSEL1 / VOYAGE1",
		},
	}

	assert.Equal(t, len(expectedData), len(events))
	assert.Equal(t, expectedData, events)
}

func TestBillTracker(t *testing.T) {
	httpMockup := requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
		if strings.Contains(r.RUrl, `TrackingView.asp`) {
			return os.ReadFile("test_data/bill_page.txt")
		}
		return os.ReadFile("test_data/container_page.txt")
	})

	dt := datetime.NewDatetime()
	tracker := NewBillTracker(&tracking.BaseConstructorArgumentsForTracker{
		Request:            httpMockup,
		UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
		Datetime:           dt,
	})

	const billNumber = "RB62CG23000870"

	data, err := tracker.Track(context.Background(), billNumber)
	assert.NoError(t, err)

	expectedInfoAboutMoving := []*tracking.Event{
		{
			Time:          time.Date(2023, 03, 01, 13, 0, 13, 0, time.UTC),
			OperationName: "PICKUP BY SHIPPER",
			Location:      "NINGBO,CHINA(CNNGB)",
			Vessel:        "VESSEL6 / VOYAGE6",
		},
		{
			Time:          time.Date(2023, 03, 03, 15, 0, 51, 0, time.UTC),
			OperationName: "EXPORT AT PORT",
			Location:      "NINGBO,CHINA(SHIPPER PREMISES)",
			Vessel:        "VESSEL5 / VOYAGE5",
		},
		{
			Time:          time.Date(2023, 03, 05, 18, 0, 13, 0, time.UTC),
			OperationName: "LOADED FULL",
			Location:      "NINGBO,CHINA(CNNGB)",
			Vessel:        "VESSEL4 / VOYAGE4",
		},
		{
			Time:          time.Date(2023, 03, 13, 7, 0, 0, 0, time.UTC),
			OperationName: "IMPORT DISCHARGE FULL",
			Location:      "VLADIVOSTOK,RUSSIA(FROM / TO VESSEL)",
			Vessel:        "VESSEL3 / VOYAGE3",
		},
		{
			Time:          time.Date(2023, 03, 27, 9, 0, 1, 0, time.UTC),
			OperationName: "PICKUP BY CONSIGNEE",
			Location:      "VLADIVOSTOK,RUSSIA(RUVVO)",
			Vessel:        "VESSEL2 / VOYAGE2",
		},
		{
			Time:          time.Date(2023, 03, 27, 11, 0, 32, 0, time.UTC),
			OperationName: "RECD MTY FROM CONSIGNEE",
			Location:      "VLADIVOSTOK,RUSSIA(CONSIGNEE PREMISES)",
			Vessel:        "VESSEL1 / VOYAGE1",
		},
	}

	assert.Equal(t, expectedInfoAboutMoving, data.InfoAboutMoving)

	assert.Equal(t, billNumber, data.Number)

	assert.Equal(t, "REEL", data.Scac)

	assert.Equal(t, time.Date(2023, 3, 5, 0, 0, 0, 0, time.UTC), data.Eta)
}
