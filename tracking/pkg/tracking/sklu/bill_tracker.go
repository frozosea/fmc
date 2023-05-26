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
	EtaParser              *EtaParser
}

func NewBillTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	return &BillTracker{
		ApiRequest:             NewApiRequest(cfg.Request, NewUrlGeneratorForApiRequest(), NewHeadersGeneratorForApiRequest(cfg.UserAgentGenerator)),
		InfoAboutMovingRequest: NewInfoAboutMovingRequest(cfg.Request, NewUrlGeneratorForInfoAboutMovingRequest(), NewHeadersGeneratorForInfoAboutMovingRequest(cfg.UserAgentGenerator)),
		ApiParser:              NewApiParser(cfg.Datetime),
		InfoAboutMovingParser:  NewInfoAboutMovingParser(cfg.Datetime),
		ContainerNumberParser:  NewContainerNumberParser(),
		EtaParser:              NewEtaParser(cfg.Datetime),
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	doc, err := b.InfoAboutMovingRequest.Send(ctx, number, "")
	if err != nil {
		return nil, err
	}
	containerNumber, err := b.ContainerNumberParser.Get(doc)
	if err != nil {
		return nil, err
	}
	//
	//apiResponse, err := b.ApiRequest.Send(ctx, number, containerNumber)
	//if err != nil {
	//	return nil, err
	//}
	//
	//containerInfo := b.ApiParser.Get(apiResponse)
	//
	//docForInfoAboutMoving, err := b.InfoAboutMovingRequest.Send(ctx, number, containerNumber)
	//if err != nil {
	//	return nil, err
	//}

	infoAboutMoving, _ := b.InfoAboutMovingParser.Get(doc, containerNumber)
	if err != nil {
		switch err.(type) {
		case *LensNotEqualError:
			infoAboutMoving = []*tracking.Event{}
		default:
			return nil, err
		}
	}

	eta, err := b.EtaParser.GetEta(doc)
	if err != nil {
		return nil, err
	}

	return &tracking.BillNumberTrackingResponse{
		Number:          number,
		Eta:             eta,
		Scac:            "SKLU",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
