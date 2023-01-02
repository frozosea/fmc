package sitc

import (
	"context"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/sitc/captcha_resolver"
)

type BillTracker struct {
	request               IBillRequest
	infoAboutMovingParser *billNumberInfoAboutMovingParser
	etaParser             *etaParser
	containerNumberParser *containerNumberParser
	captchaSolver         captcha_resolver.ICaptcha
}

func NewBillTracker(b *tracking.BaseConstructorArgumentsForTracker, request IBillRequest, captchaSolver captcha_resolver.ICaptcha) *BillTracker {
	return &BillTracker{
		request:               request,
		infoAboutMovingParser: newBillNumberInfoAboutMovingParser(b.Datetime),
		etaParser:             newEtaParser(b.Datetime),
		containerNumberParser: newContainerNumberParser(),
		captchaSolver:         captchaSolver,
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	randomString, solvedCaptcha, err := b.captchaSolver.Resolve(ctx)
	if err != nil {
		return nil, err
	}
	billNumberResponse, err := b.request.GetBillNumberInfo(ctx, number, string(randomString), string(solvedCaptcha))
	if err != nil {
		return nil, err
	}
	eta, err := b.etaParser.get(billNumberResponse)
	if err != nil {
		return nil, err
	}
	containerNumber := b.containerNumberParser.get(billNumberResponse)
	containerNumberInfoResponse, err := b.request.GetContainerInfo(ctx, number, containerNumber)
	if err != nil {
		return nil, err
	}
	infoAboutMoving := b.infoAboutMovingParser.get(containerNumberInfoResponse)
	return &tracking.BillNumberTrackingResponse{
		Number:          number,
		Eta:             eta,
		Scac:            "SITC",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
