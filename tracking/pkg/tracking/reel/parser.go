package reel

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
	"time"
	"unicode"
)

func parseDate(doc *goquery.Document, dt datetime.IDatetime, selector string, useAfterYearPattern bool) (time.Time, error) {
	selection := doc.Find(selector)
	date := selection.Text()
	if date == "" {
		return time.Now(), errors.New("no ETD")
	}
	s := strings.Split(strings.ToLower(date), "-")
	r := []rune(strings.ToLower(s[1]))
	r[0] = unicode.ToUpper(r[0])
	s[1] = string(r)
	if useAfterYearPattern {
		return dt.Strptime(strings.Join(s, "-"), "%d-%b-%Y %H:%M:%S")
	}
	return dt.Strptime(strings.Join(s, "-"), "%d-%b-%Y")
}

func parserSelector(doc *goquery.Document, selector, errorMessage string) (string, error) {
	selection := doc.Find(selector)
	text := selection.Text()
	if text == "" {
		return "", errors.New(errorMessage)
	}
	return text, nil
}

type PodParser struct {
}

func NewPodParser() *PodParser {
	return &PodParser{}
}

func (p *PodParser) Get(doc *goquery.Document) (string, error) {
	return parserSelector(doc, `#frm > table:nth-child(2) > tbody:nth-child(2) > tr:nth-child(3) > td:nth-child(5)`, `no pod`)
}

type ETDParser struct {
	datetime datetime.IDatetime
}

func NewETDParser(datetime datetime.IDatetime) *ETDParser {
	return &ETDParser{datetime: datetime}
}
func (e *ETDParser) Get(doc *goquery.Document) (time.Time, error) {
	return parseDate(doc, e.datetime, `#frm > table:nth-child(2) > tbody:nth-child(2) > tr:nth-child(5) > td:nth-child(2)`, true)
}

type ContainerNumberParser struct {
}

func NewContainerNumberParser() *ContainerNumberParser {
	return &ContainerNumberParser{}
}

func (c *ContainerNumberParser) Get(doc *goquery.Document) (string, error) {
	selector := doc.Find(`#BottomFocus > tbody > tr.Mainlabel2 > td:nth-child(3) > a`)
	val, exists := selector.Attr("onclick")
	if !exists || val == "" {
		return "", errors.New("no container number")
	}
	s := strings.Split(val, `'`)
	if len(s) > 1 {
		return s[1], nil
	}
	return "", errors.New("no container number")
}

type containerStatus struct {
	Number    string
	Type      string
	EventDate time.Time
	Status    string
	Location  string
}

type ContainerStatusParser struct {
	dt           datetime.IDatetime
	numberParser *ContainerNumberParser
}

func NewContainerStatusParser(dt datetime.IDatetime) *ContainerStatusParser {
	return &ContainerStatusParser{dt: dt, numberParser: NewContainerNumberParser()}
}

func (c *ContainerStatusParser) Get(doc *goquery.Document) (*containerStatus, error) {
	number, err := c.numberParser.Get(doc)
	if err != nil {
		return nil, err
	}

	containerType, _ := parserSelector(doc, `#BottomFocus > tbody > tr.Mainlabel2 > td:nth-child(5)`, `no container type`)

	eventDate, _ := parseDate(doc, c.dt, `#BottomFocus > tbody > tr.Mainlabel2 > td:nth-child(6)`, false)

	operationName, _ := parserSelector(doc, `#BottomFocus > tbody > tr.Mainlabel2 > td:nth-child(7)`, `no operation name`)

	location, _ := parserSelector(doc, `#BottomFocus > tbody > tr.Mainlabel2 > td:nth-child(8)`, `no location`)

	return &containerStatus{
		Number:    number,
		Type:      containerType,
		EventDate: eventDate,
		Status:    operationName,
		Location:  location,
	}, nil
}

type billInfoAboutMovingParser struct {
}

func newBillInfoAboutMovingParser() *billInfoAboutMovingParser {
	return &billInfoAboutMovingParser{}
}

func (b *billInfoAboutMovingParser) Get(s *containerStatus) (*tracking.Event, error) {
	if s.Status == "" {
		return nil, errors.New("no operation name")
	}

	if s.Location == "" {
		return nil, errors.New("no location")
	}

	return &tracking.Event{
		Time:          s.EventDate,
		OperationName: strings.ToUpper(s.Status),
		Location:      strings.ToUpper(s.Location),
		Vessel:        "",
	}, nil

}

type billMainInfo struct {
	POD             string
	ETD             time.Time
	lastEvent       *tracking.Event
	containerStatus *containerStatus
}

type billMainInfoParser struct {
	podParser                 *PodParser
	etdParser                 *ETDParser
	billInfoAboutMovingParser *billInfoAboutMovingParser
	containerStatusParser     *ContainerStatusParser
}

func newBillMainInfoParser(dt datetime.IDatetime) *billMainInfoParser {
	return &billMainInfoParser{
		podParser:                 NewPodParser(),
		etdParser:                 NewETDParser(dt),
		containerStatusParser:     NewContainerStatusParser(dt),
		billInfoAboutMovingParser: newBillInfoAboutMovingParser(),
	}
}

func (b *billMainInfoParser) Get(doc *goquery.Document) (*billMainInfo, error) {
	pod, _ := b.podParser.Get(doc)

	etd, _ := b.etdParser.Get(doc)

	containerStatusInfo, err := b.containerStatusParser.Get(doc)
	if err == nil {
		lastEvent, err := b.billInfoAboutMovingParser.Get(containerStatusInfo)
		if err != nil {
			return nil, err
		}
		return &billMainInfo{
			POD:             pod,
			ETD:             etd,
			containerStatus: containerStatusInfo,
			lastEvent:       lastEvent,
		}, nil
	}

	return &billMainInfo{
		POD:             pod,
		ETD:             etd,
		containerStatus: nil,
		lastEvent:       nil,
	}, tracking.NewNotThisLineException()

}

type infoAboutMovingParser struct {
	dt datetime.IDatetime
}

func newInfoAboutMovingParser(dt datetime.IDatetime) *infoAboutMovingParser {
	return &infoAboutMovingParser{dt: dt}
}
func (i *infoAboutMovingParser) getLoopLen(doc *goquery.Document) int {
	selection := doc.Find(`#row > div > div > div:nth-child(2) > table > tbody`)
	return selection.Children().Length()
}
func (i *infoAboutMovingParser) Get(doc *goquery.Document) ([]*tracking.Event, error) {
	var events []*tracking.Event

	for index := i.getLoopLen(doc); index >= 2; index-- {
		eventDate, err := parseDate(doc, i.dt, fmt.Sprintf(`#row > div > div > div:nth-child(2) > table > tbody > tr:nth-child(%d) > td:nth-child(3)`, index), true)
		if err != nil {
			continue
		}

		operationName, err := parserSelector(doc, fmt.Sprintf(`#row > div > div > div:nth-child(2) > table > tbody > tr:nth-child(%d) > td:nth-child(2)`, index), `no operation name`)
		if err != nil {
			continue
		}

		location, _ := parserSelector(doc, fmt.Sprintf(`#row > div > div > div:nth-child(2) > table > tbody > tr:nth-child(%d) > td:nth-child(4)`, index), `no location`)

		vessel, _ := parserSelector(doc, fmt.Sprintf(`#row > div > div > div:nth-child(2) > table > tbody > tr:nth-child(%d) > td:nth-child(6)`, index), `no vessel`)

		events = append(events, &tracking.Event{Time: eventDate, OperationName: strings.ToUpper(operationName), Location: strings.ToUpper(location), Vessel: strings.ToUpper(vessel)})
	}

	return events, nil
}
