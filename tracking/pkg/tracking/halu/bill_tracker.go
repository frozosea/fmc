package halu

import (
	"context"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/sklu"
)

type BillTracker struct {
	*sklu.BillTracker
}

func NewBillTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	return &BillTracker{
		BillTracker: &sklu.BillTracker{
			ApiRequest:             sklu.NewApiRequest(cfg.Request, NewUrlGeneratorForApiRequest(), NewHeadersGeneratorForApiRequest(cfg.UserAgentGenerator)),
			InfoAboutMovingRequest: sklu.NewInfoAboutMovingRequest(cfg.Request, NewUrlGeneratorForInfoAboutMovingRequest(), NewHeadersGeneratorForInfoAboutMovingRequest(cfg.UserAgentGenerator)),
			ApiParser:              sklu.NewApiParser(cfg.Datetime),
			InfoAboutMovingParser:  sklu.NewInfoAboutMovingParser(cfg.Datetime),
			ContainerNumberParser:  sklu.NewContainerNumberParser(),
		},
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	response, err := b.BillTracker.Track(ctx, number)
	if err != nil {
		return nil, err
	}
	response.Scac = "HALU"
	return response, nil
}
