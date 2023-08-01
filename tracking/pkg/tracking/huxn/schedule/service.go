package huaxin_schedule

import (
	"context"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/huxn"
	"golang_tracking/pkg/tracking/huxn/schedule/unlocodesParser"
	"golang_tracking/pkg/tracking/util/datetime"
	"golang_tracking/pkg/tracking/util/requests"
)

type Service struct {
	unlocodesProvider unlocodesParser.IService
	request           *Request
	parser            *Parser
	scheduleParser    *ServerResponseScheduleParser
}

func NewService(unlocodesProvider unlocodesParser.IService, ihttp requests.IHttp, uaGenerator requests.IUserAgentGenerator, dt datetime.IDatetime) *Service {
	return &Service{
		unlocodesProvider: unlocodesProvider,
		request:           NewRequest(ihttp, uaGenerator),
		scheduleParser:    NewScheduleParser(dt),
		parser:            NewParser(dt),
	}
}

func (s *Service) GetETA(ctx context.Context, data *huxn.TrackingResponse) ([]*tracking.Event, error) {
	dataForScheduleRequest, err := s.parser.GetDataForScheduleRequest(data)
	if err != nil {
		return nil, err
	}
	allUnlocodes, err := s.unlocodesProvider.GetUnlocodes(ctx, dataForScheduleRequest.lastPortUnlocode)
	if err != nil {
		return nil, err
	}
	dataForScheduleRequest.etd = dataForScheduleRequest.etd.AddDate(0, 0, -7)
	schedulesResponse, err := s.request.GetWholeWorldSchedule(ctx, dataForScheduleRequest.lastPortUnlocode, allUnlocodes, dataForScheduleRequest.etd)
	if err != nil {
		return nil, err
	}
	return s.scheduleParser.GetETA(dataForScheduleRequest.lastVoyage, schedulesResponse)
}
