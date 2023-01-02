package cosu

import (
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

func (p *PodParser) get(apiResponse *ApiResponseSchema) string {
	return apiResponse.Data.Content.Containers[0].Container.Pod
}

type EtaParser struct {
	datetime datetime.IDatetime
}

func NewEtaParser(datetime datetime.IDatetime) *EtaParser {
	return &EtaParser{datetime: datetime}
}
func (e *EtaParser) get(etaRawResp *EtaApiResponseSchema, POD string) (*tracking.Event, error) {
	stringEta := etaRawResp.Data.Content
	if stringEta == "" {
		return nil, NewGetEtaError()
	}
	eta, err := e.datetime.Strptime(stringEta, "%Y-%m-%d %H:%M")
	if err != nil {
		return nil, err
	}
	return &tracking.Event{
		Time:          eta,
		OperationName: "ETA",
		Location:      POD,
		Vessel:        "",
	}, nil
}

type ContainerSizeParser struct {
}

func NewContainerSizeParser() *ContainerSizeParser {
	return &ContainerSizeParser{}
}

func (c *ContainerSizeParser) get(apiResponse *ApiResponseSchema) string {
	return strings.ToUpper(apiResponse.Data.Content.Containers[0].Container.ContainerType)
}

type InfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func NewInfoAboutMovingParser(datetime datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{datetime: datetime}
}

func (i *InfoAboutMovingParser) get(apiResp *ApiResponseSchema) []*tracking.Event {
	var infoAboutMoving []*tracking.Event
	containerHistory := apiResp.Data.Content.Containers[0].ContainerHistorys
	for _, historyEvent := range containerHistory {
		var eventTime time.Time
		if historyEvent.TimeOfIssue != "" {
			//YYYY-MM-DD HH:mm
			t, err := i.datetime.Strptime(historyEvent.TimeOfIssue, "%Y-%m-%d %H:%M")
			if err != nil {
				eventTime = time.Time{}
			} else {
				eventTime = t
			}
		}
		infoAboutMoving = append(infoAboutMoving, &tracking.Event{
			Time:          eventTime,
			OperationName: strings.ToUpper(historyEvent.ContainerNumberStatus),
			Location:      strings.ToUpper(historyEvent.Location),
			Vessel:        strings.ToUpper(historyEvent.Transportation),
		})
	}
	return infoAboutMoving
}
