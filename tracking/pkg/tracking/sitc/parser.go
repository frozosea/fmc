package sitc

import (
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
	"time"
)

type ContainerTrackingInfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func NewContainerTrackingInfoAboutMovingParser(datetime datetime.IDatetime) *ContainerTrackingInfoAboutMovingParser {
	return &ContainerTrackingInfoAboutMovingParser{datetime: datetime}
}

func (c *ContainerTrackingInfoAboutMovingParser) get(response *ContainerApiResponse) []*tracking.Event {
	var infoAboutMoving []*tracking.Event
	for _, item := range response.Data.List {
		//2022-06-04 22:00:00
		eventTime, err := c.datetime.Strptime(item.EventDate, "%Y-%m-%d %H:%M:%S")
		if err != nil {
			continue
		}
		infoAboutMoving = append(infoAboutMoving, &tracking.Event{
			Time:          eventTime,
			OperationName: strings.ToTitle(strings.ToLower(strings.Trim(item.MovementNameEn, " "))),
			Location:      strings.ToUpper(strings.Trim(item.EventPort, " ")),
			Vessel:        strings.ToUpper(strings.Trim(item.VesselCode, " ")),
		})
	}
	return tracking.ReverseSliceWithEvents(infoAboutMoving)
}

type containerNumberParser struct {
}

func newContainerNumberParser() *containerNumberParser {
	return &containerNumberParser{}
}

func (c *containerNumberParser) get(response *BillNumberApiResponse) string {
	if len(response.Data.List3) == 0 {
		return ""
	}
	return response.Data.List3[0].ContainerNo
}

type etaParser struct {
	datetime datetime.IDatetime
}

func newEtaParser(datetime datetime.IDatetime) *etaParser {
	return &etaParser{datetime: datetime}
}

func (e *etaParser) get(response *BillNumberApiResponse) (time.Time, error) {
	lastIndex := len(response.Data.List2) - 1
	return e.datetime.Strptime(response.Data.List2[lastIndex].Eta, "%y-%m-%d %H:%M")
}

type billNumberInfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func newBillNumberInfoAboutMovingParser(datetime datetime.IDatetime) *billNumberInfoAboutMovingParser {
	return &billNumberInfoAboutMovingParser{datetime: datetime}
}

func (b *billNumberInfoAboutMovingParser) get(response *BillNumberInfoAboutContainerApiResponse) []*tracking.Event {
	var infoAboutMoving []*tracking.Event
	for _, item := range response.Data.List {
		//YYYY-MM-DD
		eventTime, err := b.datetime.Strptime(item.Eventdate, "%y-%m-%d")
		if err != nil {
			continue
		}
		infoAboutMoving = append(infoAboutMoving, &tracking.Event{
			Time:          eventTime,
			OperationName: strings.ToTitle(strings.ToLower(strings.Trim(item.Movementnameen, " "))),
			Location:      strings.ToUpper(strings.Trim(item.Portname, " ")),
			Vessel:        "",
		})
	}
	return infoAboutMoving
}
