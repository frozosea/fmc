package huxn

import (
	"context"
	"golang_tracking/pkg/tracking"
	"time"
)

type BillTracker struct {
	etaParser             IETAProvider
	request               *ContainerTrackingRequest
	infoAboutMovingParser *InfoAboutMovingParser
	containerSizeParser   *ContainerSizeParser
}

func NewBillTracker(cfg *tracking.BaseConstructorArgumentsForTracker, etaProvider IETAProvider) *BillTracker {
	return &BillTracker{
		etaParser:             etaProvider,
		request:               NewContainerTrackingRequest(cfg.Request, cfg.UserAgentGenerator),
		infoAboutMovingParser: NewInfoAboutMovingParser(cfg.Datetime),
		containerSizeParser:   NewContainerSizeParser(),
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	var response = &tracking.BillNumberTrackingResponse{
		Number:          number,
		Eta:             time.Time{},
		Scac:            "HUXN",
		InfoAboutMoving: nil,
	}

	data, err := b.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}
	if len(data.ListDynamics) == 0 {
		return nil, tracking.NewNotThisLineException()
	}
	infoAboutMoving, _ := b.infoAboutMovingParser.Parse(data)

	if !checkNumberArrived(infoAboutMoving) {
		etaEvents, err := b.etaParser.GetETA(ctx, data)
		if err == nil {
			infoAboutMoving = append(infoAboutMoving, etaEvents...)
			response.Eta = etaEvents[len(etaEvents)-1].Time
			if len(infoAboutMoving) > 0 {
				if infoAboutMoving[len(infoAboutMoving)-1].Time.Unix() < etaEvents[len(etaEvents)-1].Time.Unix() {
					infoAboutMoving = append(infoAboutMoving, etaEvents...)
				}
			}
		}
	}

	response.InfoAboutMoving = infoAboutMoving

	return response, nil
}
