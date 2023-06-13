package dnyg

import (
	"context"
	"errors"
	"golang_tracking/pkg/tracking"
)

type BillTracker struct {
	numberInfoRequest      INumberInfoRequest
	infoAboutMovingRequest IInfoAboutMovingRequest
	infoAboutMovingParser  *InfoAboutMovingParser
	etaParser              *EtaParser
	podParser              *PODParser
}

func NewBillTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	return &BillTracker{
		numberInfoRequest:      NewNumberInfoRequest(cfg.Request, cfg.UserAgentGenerator),
		infoAboutMovingRequest: NewInfoAboutMovingRequest(cfg.Request, cfg.UserAgentGenerator),
		infoAboutMovingParser:  NewInfoAboutMovingParser(cfg.Datetime),
		etaParser:              NewEtaParser(cfg.Datetime),
		podParser:              NewPODParser(),
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	containerNumberInfo, err := b.numberInfoRequest.Send(ctx, number, false)
	if err != nil {
		return nil, err
	}
	if len(containerNumberInfo.DltResultBlList) > 0 {
		infoAboutMovingRawResponse, err := b.infoAboutMovingRequest.Send(ctx, containerNumberInfo.DltResultBlList[0].OUTBKN, containerNumberInfo.DltResultBlList[0].OUTBNO, containerNumberInfo.DltResultBlList[0].OUTCNT)
		if err != nil {
			return nil, err
		}

		infoAboutMoving := b.infoAboutMovingParser.Parse(infoAboutMovingRawResponse)

		eta, err := b.etaParser.parse(containerNumberInfo)

		return &tracking.BillNumberTrackingResponse{
			Number:          number,
			Eta:             eta,
			Scac:            "DNYG",
			InfoAboutMoving: infoAboutMoving,
		}, nil
	} else {
		return nil, errors.New("invalid data")
	}
}
