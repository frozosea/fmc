package reel

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type BillTracker struct {
	billRequest                    *BillRequest
	containerRequest               *ContainerInfoRequest
	billInfoParser                 *billMainInfoParser
	containerInfoAboutMovingParser *infoAboutMovingParser
}

func NewBillTracker(args *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	headersGenerator := NewHeadersGeneratorForInfoAboutMovingRequest(args.UserAgentGenerator)
	return &BillTracker{
		billRequest:                    NewBillRequest(args.Request, NewUrlGeneratorForBillRequest(), headersGenerator),
		containerRequest:               NewContainerInfoRequest(args.Request, headersGenerator, NewUrlGeneratorForContainerInfo()),
		billInfoParser:                 newBillMainInfoParser(args.Datetime),
		containerInfoAboutMovingParser: newInfoAboutMovingParser(args.Datetime)}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	billDoc, err := b.billRequest.Send(ctx, number)
	if err != nil {
		return nil, err
	}

	billInfo, err := b.billInfoParser.Get(billDoc)
	if err != nil {
		return nil, err
	}

	var infoAboutMoving []*tracking.Event
	infoAboutMovingDoc, err := b.containerRequest.Send(ctx, number, billInfo.containerStatus.Number)
	if err == nil {
		containerInfoAboutMoving, err := b.containerInfoAboutMovingParser.Get(infoAboutMovingDoc)
		if err == nil {
			infoAboutMoving = containerInfoAboutMoving
		}
	} else {
		infoAboutMoving = append(infoAboutMoving, billInfo.lastEvent)

	}

	return &tracking.BillNumberTrackingResponse{
		Number:          number,
		Eta:             billInfo.ETD,
		Scac:            "REEL",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
