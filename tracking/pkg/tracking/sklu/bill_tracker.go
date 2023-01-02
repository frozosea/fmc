package sklu

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type BillTracker struct {
	ApiRequest             *ApiRequest
	InfoAboutMovingRequest *InfoAboutMovingRequest
	ApiParser              *ApiParser
	InfoAboutMovingParser  *InfoAboutMovingParser
	ContainerNumberParser  *ContainerNumberParser
}

func NewBillTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	return &BillTracker{
		ApiRequest:             NewApiRequest(cfg.Request, NewUrlGeneratorForApiRequest(), NewHeadersGeneratorForApiRequest(cfg.UserAgentGenerator)),
		InfoAboutMovingRequest: NewInfoAboutMovingRequest(cfg.Request, NewUrlGeneratorForInfoAboutMovingRequest(), NewHeadersGeneratorForInfoAboutMovingRequest(cfg.UserAgentGenerator)),
		ApiParser:              NewApiParser(cfg.Datetime),
		InfoAboutMovingParser:  NewInfoAboutMovingParser(cfg.Datetime),
		ContainerNumberParser:  NewContainerNumberParser(),
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	docForParseContainerNumber, err := b.InfoAboutMovingRequest.Send(ctx, number, "")
	if err != nil {
		panic(err)
		return nil, err
	}

	containerNumber, err := b.ContainerNumberParser.Get(docForParseContainerNumber)
	if err != nil {
		return nil, err
	}

	apiResponse, err := b.ApiRequest.Send(ctx, number, containerNumber)
	if err != nil {
		return nil, err
	}

	containerInfo := b.ApiParser.Get(apiResponse)

	docForInfoAboutMoving, err := b.InfoAboutMovingRequest.Send(ctx, number, containerNumber)
	if err != nil {
		return nil, err
	}

	infoAboutMoving, err := b.InfoAboutMovingParser.Get(docForInfoAboutMoving, containerNumber)
	if err != nil {
		return nil, err
	}

	return &tracking.BillNumberTrackingResponse{
		Number:          number,
		Eta:             containerInfo.Eta,
		Scac:            "SKLU",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
