package zhgu

import (
	"errors"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
	"time"
)

type EtaParser struct {
	datetime datetime.IDatetime
}

func NewEtaParser(datetime datetime.IDatetime) *EtaParser {
	return &EtaParser{datetime: datetime}
}

func (e *EtaParser) Get(response *ApiResponseSchema) (time.Time, error) {
	eta := response.Object[0].Eta
	if eta == "" {
		return time.Time{}, NewGetEtaException()
	}
	return e.datetime.Strptime(eta, "%Y-%m-%d")
}

func parserWrap(datetime datetime.IDatetime, response *ApiResponseSchema, value, operationName string) (*tracking.Event, error) {
	if value == "" {
		return nil, NewGetEtdException()
	}
	etd, err := datetime.Strptime(value, "%Y-%m-%d")
	if err != nil {
		return nil, NewGetEtdException()
	}
	return &tracking.Event{
		Time:          etd,
		OperationName: operationName,
		Location:      strings.ToTitle(response.Object[0].PortFromName),
		Vessel:        strings.ToUpper(response.Object[0].VesselName),
	}, nil
}

type InfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func NewInfoAboutMovingParser(datetime datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{datetime: datetime}
}

func (i *InfoAboutMovingParser) Get(response *ApiResponseSchema) ([]*tracking.Event, error) {
	var infoAboutMoving []*tracking.Event

	etd, err := parserWrap(i.datetime, response, response.Object[0].Etd, "ETD")
	if err == nil {
		infoAboutMoving = append(infoAboutMoving, etd)
	}

	atd, err := parserWrap(i.datetime, response, response.Object[0].Atd, "ATD")
	if err == nil {
		infoAboutMoving = append(infoAboutMoving, atd)
	}

	ata, err := parserWrap(i.datetime, response, response.Object[0].Ata, "ATA")
	if err == nil {
		infoAboutMoving = append(infoAboutMoving, ata)
	}
	if len(infoAboutMoving) == 0 {
		return nil, errors.New("no len")
	}
	return infoAboutMoving, nil
}
