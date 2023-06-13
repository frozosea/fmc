package akkn

import (
	"github.com/PuerkitoBio/goquery"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"time"
)

func parseTime(doc *goquery.Document, selector string, dt datetime.IDatetime) (time.Time, error) {
	return dt.Strptime(doc.Find(selector).Text(), "%d/%m/%Y")
}

type InfoAboutMovingParser struct {
	dt datetime.IDatetime
}

func NewInfoAboutMovingParser(dt datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{dt: dt}
}

func (i *InfoAboutMovingParser) parseLoadOnVesselEvent(doc *goquery.Document) (*tracking.Event, error) {
	loc := doc.Find("#GridView1 > tbody > tr:nth-child(2) > td:nth-child(7)").Text()
	date, err := parseTime(doc, "#GridView1 > tbody > tr:nth-child(2) > td:nth-child(8)", i.dt)
	if err != nil {
		return nil, err
	}
	vessel := doc.Find("#GridView1 > tbody > tr:nth-child(2) > td:nth-child(1)").Text()
	return &tracking.Event{
		Time:          date,
		OperationName: "LOAD ON VESSEL",
		Location:      loc,
		Vessel:        vessel,
	}, nil
}
func (i *InfoAboutMovingParser) parseArriveEvent(doc *goquery.Document) (*tracking.Event, error) {
	loc := doc.Find("#GridView1 > tbody > tr:nth-child(2) > td:nth-child(9)").Text()
	date, err := parseTime(doc, "#GridView1 > tbody > tr:nth-child(2) > td:nth-child(10)", i.dt)
	if err != nil {
		return nil, err
	}
	vessel := doc.Find("#GridView1 > tbody > tr:nth-child(2) > td:nth-child(1)").Text()
	return &tracking.Event{
		Time:          date,
		OperationName: "ARRIVE AT PORT OF DISCHARGING",
		Location:      loc,
		Vessel:        vessel,
	}, nil
}
func (i *InfoAboutMovingParser) Parse(doc *goquery.Document) []*tracking.Event {
	var infoAboutMoving []*tracking.Event

	firstEvent, err := i.parseLoadOnVesselEvent(doc)
	if err == nil {
		infoAboutMoving = append(infoAboutMoving, firstEvent)
	}

	arriveEvent, err := i.parseArriveEvent(doc)
	if err == nil {
		infoAboutMoving = append(infoAboutMoving, arriveEvent)
	}

	return infoAboutMoving
}

type EtaParser struct {
	dt datetime.IDatetime
}

func NewEtaParser(dt datetime.IDatetime) *EtaParser {
	return &EtaParser{dt: dt}
}

func (e *EtaParser) Parse(doc *goquery.Document) (time.Time, error) {
	return parseTime(doc, "#GridView1 > tbody > tr:nth-child(2) > td:nth-child(10)", e.dt)
}
