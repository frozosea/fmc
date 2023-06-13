package akkn

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type BillTracker struct {
	request               *Request
	infoAboutMovingParser *InfoAboutMovingParser
	etaParser             *EtaParser
}

func NewBillTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	return &BillTracker{
		request:               NewRequest(cfg.Request, cfg.UserAgentGenerator),
		infoAboutMovingParser: NewInfoAboutMovingParser(cfg.Datetime),
		etaParser:             NewEtaParser(cfg.Datetime),
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	doc, err := b.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}

	infoAboutMoving := b.infoAboutMovingParser.Parse(doc)
	eta, _ := b.etaParser.Parse(doc)

	return &tracking.BillNumberTrackingResponse{
		Number:          number,
		Eta:             eta,
		Scac:            "AKKN",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
