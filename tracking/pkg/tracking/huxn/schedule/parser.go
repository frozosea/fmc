package huaxin_schedule

import (
	"errors"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/huxn"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
	"time"
)

type Parser struct {
	dt datetime.IDatetime
}

func NewParser(dt datetime.IDatetime) *Parser {
	return &Parser{dt: dt}
}
func (s *Parser) getLastVoyage(data *huxn.TrackingResponse) string {
	events := data.ListDynamics
	if len(events) == 0 {
		return ""
	}

	lastEventVesselAndVoyage := events[0].VESSELVOYAGE
	if lastEventVesselAndVoyage == "" {
		return ""
	}

	split := strings.Split(lastEventVesselAndVoyage, "/")
	if len(split) == 1 {
		return ""
	}

	rawVoyage := split[1]

	return strings.Trim(rawVoyage, " ")
}

func (s *Parser) getLastPortUnlocode(response *huxn.TrackingResponse) (string, error) {
	if len(response.ListDynamics) == 0 {
		return "", errors.New("no len")
	}

	rawPort := response.ListDynamics[0].PORTFULLNAME

	split := strings.Split(rawPort, "(")
	if len(split) == 1 {
		return "", errors.New("no len")
	}

	unlocode := strings.Replace(split[len(split)-1], ")", "", 1)
	return strings.ToUpper(unlocode), nil
}
func (s *Parser) getETD(response *huxn.TrackingResponse) (time.Time, error) {
	if len(response.ListDynamics) == 0 {
		return time.Time{}, errors.New("no len")
	}

	rawTime := response.ListDynamics[0].DYNTIME
	return s.dt.Strptime(rawTime, "%Y/%m/%d %H:%M:%S")
}

func (s *Parser) GetDataForScheduleRequest(response *huxn.TrackingResponse) (*DataForScheduleRequest, error) {
	voyage := s.getLastVoyage(response)
	if voyage == "" {
		return nil, errors.New("no voyage")
	}
	portUnlocode, err := s.getLastPortUnlocode(response)
	if err != nil {
		return nil, err
	}
	etd, err := s.getETD(response)
	if err != nil {
		return nil, err
	}
	return &DataForScheduleRequest{
		lastVoyage:       voyage,
		lastPortUnlocode: portUnlocode,
		etd:              etd,
	}, nil
}

type ServerResponseScheduleParser struct {
	dt datetime.IDatetime
}

func NewScheduleParser(dt datetime.IDatetime) *ServerResponseScheduleParser {
	return &ServerResponseScheduleParser{dt: dt}
}

func (s *ServerResponseScheduleParser) GetETA(voyage string, data []*Schedule) ([]*tracking.Event, error) {
	if len(data) == 0 {
		return nil, errors.New("no len")
	}

	var events []*tracking.Event

	for _, item := range data {
		if item.VOYAGE == voyage {
			t, err := s.dt.Strptime(item.DISCETA, "%Y-%m-%d")
			if err != nil {
				return nil, err
			}
			events = append(events, &tracking.Event{
				Time:          t,
				OperationName: "ETA",
				Location:      item.DISCPIERNAME,
				Vessel:        item.VESSEL,
			})
		}
	}
	return events, nil
}
