package feso

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type BillTracker struct {
	*ContainerTracker
	etaParser *EtaParser
}

func NewBillTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	return &BillTracker{ContainerTracker: NewContainerTracker(cfg), etaParser: NewEtaParser(cfg.Datetime)}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	response, err := b.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}
	infoAboutMoving := b.infoAboutMovingParser.get(response.LastEvents)
	eta, etaIndexInArray, err := b.etaParser.Get(infoAboutMoving)
	if err != nil {
		return nil, err
	}
	return &tracking.BillNumberTrackingResponse{
		Number:          number,
		Eta:             eta,
		Scac:            "FESO",
		InfoAboutMoving: append(infoAboutMoving[:etaIndexInArray], infoAboutMoving[etaIndexInArray+1:]...),
	}, nil
}
