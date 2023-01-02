package maeu

import (
	"fmt"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
	"time"
)

type PodParser struct {
}

func NewPodParser() *PodParser {
	return &PodParser{}
}

func (p *PodParser) get(response *ApiResponse) string {
	if response.Destination.Terminal != "" {
		return strings.ToUpper(response.Destination.Terminal)
	} else {
		return strings.ToUpper(response.Destination.City)
	}
}

type ContainerSizeParser struct {
}

func NewContainerSizeParser() *ContainerSizeParser {
	return &ContainerSizeParser{}
}

func (c *ContainerSizeParser) get(response *ApiResponse) string {
	return strings.ToUpper(fmt.Sprintf(`%s%s`,
		response.Containers[0].ContainerSize,
		response.Containers[0].ContainerType,
	))
}

type EtaParser struct {
	datetime  datetime.IDatetime
	porParser *PodParser
}

func NewEtaParser(datetime datetime.IDatetime) *EtaParser {
	return &EtaParser{datetime: datetime, porParser: NewPodParser()}
}

func (e *EtaParser) get(response *ApiResponse) (*tracking.Event, error) {
	stringEta := response.Containers[0].EtaFinalDelivery
	if stringEta != "" {
		stringEta = strings.Split(stringEta, ".")[0]
		eta, err := e.datetime.Strptime(stringEta, "%Y-%m-%dT%H:%M:%S")
		if err != nil {
			return nil, NewGetEtaException()
		}
		return &tracking.Event{
			Time:          eta,
			OperationName: "ETA",
			Location:      e.porParser.get(response),
			Vessel:        "",
		}, nil
	} else {
		return nil, NewGetEtaException()
	}
}

type InfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func NewInfoAboutMovingParser(datetime datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{datetime: datetime}
}

func (i *InfoAboutMovingParser) getTimeFromEvent(childEvent *Event) (time.Time, error) {
	var rawTime string
	if childEvent.ActualTime != "" {
		rawTime = childEvent.ActualTime
	} else {
		rawTime = childEvent.ExpectedTime
	}
	return i.datetime.Strptime(strings.Split(rawTime, ".")[0], "%Y-%m-%dT%H:%M:%S")
}
func (i *InfoAboutMovingParser) getLocation(parentEvent *Location) string {
	if parentEvent.Terminal != "" {
		if parentEvent.Country != "" {
			return strings.ToUpper(fmt.Sprintf(`%s, %s`, parentEvent.Terminal, parentEvent.Country))
		}
		return strings.ToUpper(parentEvent.Terminal)
	} else {
		if parentEvent.Country != "" {
			return strings.ToUpper(fmt.Sprintf(`%s, %s`, parentEvent.City, parentEvent.Country))
		}
		return strings.ToUpper(parentEvent.City)
	}
}
func (i *InfoAboutMovingParser) get(response *ApiResponse) []*tracking.Event {
	var infoAboutMoving []*tracking.Event
	for _, location := range response.Containers[0].Locations {
		for _, event := range location.Events {
			eventTime, err := i.getTimeFromEvent(event)
			if err != nil {
				continue
			}
			infoAboutMoving = append(infoAboutMoving, &tracking.Event{
				Time:          eventTime,
				OperationName: strings.ToTitle(strings.ToLower(event.Activity)),
				Location:      i.getLocation(location),
				Vessel:        strings.ToUpper(event.VesselName),
			})
		}
	}
	return infoAboutMoving
}
