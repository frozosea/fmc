package akkn

import (
	"bytes"
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
	"os"
	"testing"
	"time"
)

var httpMockup = requests.NewRequestMockUp(200, func(r requests.RequestMockUp) ([]byte, error) {
	return os.ReadFile("test_data/bill.txt")
})
var baseCfg = &tracking.BaseConstructorArgumentsForTracker{
	Request:            httpMockup,
	UserAgentGenerator: requests.NewUserAgentGeneratorMockUp(),
	Datetime:           datetime.NewDatetime(),
}

func getInfoAboutMovingDoc(t *testing.T) *goquery.Document {
	response, err := os.ReadFile("test_data/bill.txt")
	assert.NoError(t, err)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response))
	if err != nil {
		return nil
	}
	return doc
}

func TestInfoAboutMovingParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	doc := getInfoAboutMovingDoc(t)
	parser := NewInfoAboutMovingParser(datetime.NewDatetime())
	expectedData := []*tracking.Event{
		{
			Time:          time.Date(2023, 06, 8, 0, 0, 0, 0, time.UTC),
			OperationName: "LOAD ON VESSEL",
			Location:      "GEBZE YILPORT",
			Vessel:        "MV NARA",
		},
		{
			Time:          time.Date(2023, 06, 13, 0, 0, 0, 0, time.UTC),
			OperationName: "ARRIVE AT PORT OF DISCHARGING",
			Location:      "NOVOROSSIYSK",
			Vessel:        "MV NARA",
		},
	}
	actualData := parser.Parse(doc)
	assert.Equal(t, expectedData, actualData)
}

func TestEtaParser(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	doc := getInfoAboutMovingDoc(t)
	parser := NewEtaParser(datetime.NewDatetime())
	expectedEta := time.Date(2023, 06, 13, 0, 0, 0, 0, time.UTC)
	actualEta, err := parser.Parse(doc)
	assert.NoError(t, err)
	assert.Equal(t, expectedEta, actualEta)
}

func TestBillTracker(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	const billNumber = "AKKDIL23033434"
	tracker := NewBillTracker(baseCfg)

	expectedData := &tracking.BillNumberTrackingResponse{
		Number: billNumber,
		Eta:    time.Date(2023, 06, 13, 0, 0, 0, 0, time.UTC),
		Scac:   "AKKN",
		InfoAboutMoving: []*tracking.Event{
			{
				Time:          time.Date(2023, 06, 8, 0, 0, 0, 0, time.UTC),
				OperationName: "LOAD ON VESSEL",
				Location:      "GEBZE YILPORT",
				Vessel:        "MV NARA",
			},
			{
				Time:          time.Date(2023, 06, 13, 0, 0, 0, 0, time.UTC),
				OperationName: "ARRIVE AT PORT OF DISCHARGING",
				Location:      "NOVOROSSIYSK",
				Vessel:        "MV NARA",
			},
		},
	}

	actualData, err := tracker.Track(context.Background(), billNumber)

	assert.NoError(t, err)

	assert.Equal(t, expectedData, actualData)
}
