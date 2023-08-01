package huxn

import (
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
)

type InfoAboutMovingParser struct {
	dt datetime.IDatetime
}

func NewInfoAboutMovingParser(dt datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{dt: dt}
}

func (c *InfoAboutMovingParser) Parse(t *TrackingResponse) ([]*tracking.Event, error) {
	var events []*tracking.Event

	if len(t.ListDynamics) == 0 {
		return nil, tracking.NewNotThisLineException()
	}

	for _, rawEvent := range ReverseTrackingEvents(t.ListDynamics) {
		eventTime, err := c.dt.Strptime(rawEvent.DYNTIME, "%Y/%m/%d %H:%M:%S")
		if err != nil {
			continue
		}

		events = append(events, &tracking.Event{
			Time:          eventTime,
			OperationName: strings.ToUpper(rawEvent.DYNTYPE),
			Location:      strings.ToUpper(rawEvent.PLACENAME),
			Vessel:        strings.ToUpper(rawEvent.VESSELVOYAGE),
		})

	}

	return events, nil
}

type ContainerSizeParser struct {
}

func NewContainerSizeParser() *ContainerSizeParser {
	return &ContainerSizeParser{}
}

func (c *ContainerSizeParser) Parse(t *TrackingResponse) string {
	if len(t.ListDynamics) == 0 {
		return ""
	}

	return strings.ToUpper(t.ListDynamics[0].TGCODE)
}
