package mscu

import (
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
)

type EtaParser struct {
	datetime datetime.IDatetime
}

func NewEtaParser(datetime datetime.IDatetime) *EtaParser {
	return &EtaParser{datetime: datetime}
}

func (e *EtaParser) get(response *ApiResponse) (*tracking.Event, error) {
	stringEta := response.Data.BillOfLadings[0].GeneralTrackingInfo.FinalPodEtaDate
	if stringEta == "" {
		return nil, NewGetEtaException()
	}
	eta, err := e.datetime.Strptime(stringEta, "%d/%m/%Y")
	if err != nil {
		return nil, NewGetEtaException()
	}
	return &tracking.Event{
		Time:          eta,
		OperationName: "ETA",
		Location:      "",
		Vessel:        "",
	}, nil
}

type ContainerSizeParser struct {
}

func NewContainerSizeParser() *ContainerSizeParser {
	return &ContainerSizeParser{}
}

func (c *ContainerSizeParser) get(response *ApiResponse) string {
	return response.Data.BillOfLadings[0].ContainersInfo[0].ContainerType
}

type InfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func NewInfoAboutMovingParser(datetime datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{datetime: datetime}
}

func (i *InfoAboutMovingParser) get(response *ApiResponse) []*tracking.Event {
	var infoAboutMoving []*tracking.Event
	for _, event := range response.Data.BillOfLadings[0].ContainersInfo[0].Events {
		time, err := i.datetime.Strptime(event.Date, "%d/%m/%Y")
		if err != nil {
			continue
		}
		infoAboutMoving = append(infoAboutMoving, &tracking.Event{
			Time:          time,
			OperationName: strings.ToTitle(strings.ToLower(strings.Trim(event.Description, " "))),
			Location:      strings.ToUpper(strings.Trim(event.Location, " ")),
			Vessel:        "",
		})
	}
	return tracking.ReverseSliceWithEvents(infoAboutMoving)
}
