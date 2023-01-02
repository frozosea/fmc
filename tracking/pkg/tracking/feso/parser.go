package feso

import (
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
	"time"
)

type ContainerSizeParser struct {
}

func NewContainerSizeParser() *ContainerSizeParser {
	return &ContainerSizeParser{}
}

func (c *ContainerSizeParser) Get(response *ResponseSchema) string {
	return response.ContainerCTCode
}

type InfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func NewInfoAboutMovingParser(datetime datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{datetime: datetime}
}
func (f *InfoAboutMovingParser) parseTime(event *fesoLastEventSchema) time.Time {
	var eventTime time.Time
	if event.Time != "" {
		//2022-06-11T04:51:35
		t, err := f.datetime.Strptime(event.Time, "%Y-%m-%dT%H:%M:%S")
		if err != nil {
			eventTime = time.Time{}
		} else {
			eventTime = t
		}
	}
	return eventTime
}
func (f *InfoAboutMovingParser) get(events []*fesoLastEventSchema) []*tracking.Event {
	var infoAboutMoving []*tracking.Event
	for _, rawEvent := range events {
		infoAboutMoving = append(infoAboutMoving, &tracking.Event{
			Time:          f.parseTime(rawEvent),
			OperationName: strings.ToTitle(strings.ToLower(strings.Trim(rawEvent.OperationNameLatin, " "))),
			Location:      strings.ToUpper(strings.Trim(rawEvent.LocNameLatin, " ")),
			Vessel:        strings.ToUpper(rawEvent.Vessel),
		})
	}
	return infoAboutMoving
}

type EtaParser struct {
	datetime datetime.IDatetime
}

func NewEtaParser(datetime datetime.IDatetime) *EtaParser {
	return &EtaParser{datetime: datetime}
}

func (e *EtaParser) Get(infoAboutMoving []*tracking.Event) (time.Time, int, error) {
	for index, v := range infoAboutMoving {
		if strings.ToUpper(v.OperationName) == "ETA" {
			return v.Time, index, nil
		}
	}
	return time.Time{}, -1, NewGetEtaException()
}
