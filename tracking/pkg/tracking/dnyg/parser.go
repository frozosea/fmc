package dnyg

import (
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
	"time"
)

type InfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func NewInfoAboutMovingParser(datetime datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{datetime: datetime}
}

func (i *InfoAboutMovingParser) Parse(raw *InfoAboutMovingResponse) []*tracking.Event {
	var infoAboutMoving []*tracking.Event

	for _, rawEvent := range raw.DltResultMovementList {
		//202306092330
		eventTime, err := i.datetime.Strptime(rawEvent.OUTDTD, "%Y%m%d%H%M")
		if err != nil {
			continue
		}
		infoAboutMoving = append(infoAboutMoving, &tracking.Event{
			Time:          eventTime,
			OperationName: strings.ToUpper(strings.Trim(rawEvent.OUTDEK, " ")),
			Location:      strings.ToUpper(strings.Trim(rawEvent.OUTAREA, " ")),
			Vessel:        strings.ToUpper(strings.Trim(rawEvent.VSLNAME, " ")),
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

func (e *EtaParser) parse(response *NumberInfoResponse) (time.Time, error) {
	return e.datetime.Strptime(response.DltResultBlList[0].OUTETA, "%Y%m%d%H%M")
}

type PODParser struct {
}

func NewPODParser() *PODParser {
	return &PODParser{}
}
func (p *PODParser) parse(response *NumberInfoResponse) string {
	split := strings.Split(response.DltResultBlList[0].OUTPOD, "<")
	if len(split) > 1 {
		return strings.Trim(strings.ToUpper(split[0]), " ")
	}
	return ""
}
