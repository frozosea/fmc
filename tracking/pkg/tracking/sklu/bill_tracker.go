package sklu

import (
	"context"
	"errors"
	"golang_tracking/pkg/tracking"
)

type BillTracker struct {
	NumberExistsChecker    ICheckBookingNumberExists
	ApiRequest             *ApiRequest
	InfoAboutMovingRequest *InfoAboutMovingRequest
	ApiParser              *ApiParser
	InfoAboutMovingParser  *InfoAboutMovingParser
	ContainerNumberParser  *ContainerNumberParser
	EtaParser              *EtaParser
}

func NewBillTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	return &BillTracker{
		NumberExistsChecker:    NewCheckBookingNumberExists(cfg.Request, NewCheckBookingNumberExistsUrlGenerator(), NewCheckNumberExistsHeadersGenerator()),
		ApiRequest:             NewApiRequest(cfg.Request, NewUrlGeneratorForApiRequest(), NewHeadersGeneratorForApiRequest(cfg.UserAgentGenerator)),
		InfoAboutMovingRequest: NewInfoAboutMovingRequest(cfg.Request, NewUrlGeneratorForInfoAboutMovingRequest(), NewHeadersGeneratorForInfoAboutMovingRequest(cfg.UserAgentGenerator)),
		ApiParser:              NewApiParser(cfg.Datetime),
		InfoAboutMovingParser:  NewInfoAboutMovingParser(cfg.Datetime),
		ContainerNumberParser:  NewContainerNumberParser(),
		EtaParser:              NewEtaParser(cfg.Datetime),
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	if exists, err := b.NumberExistsChecker.CheckExists(ctx, number); !exists || err != nil {
		return nil, errors.New("doesnt exists")
	}
	doc, err := b.InfoAboutMovingRequest.Send(ctx, number, "")
	if err != nil {
		return nil, err
	}
	containerNumber, _ := b.ContainerNumberParser.Get(doc)

	infoAboutMoving, err := b.InfoAboutMovingParser.Get(doc, containerNumber)
	if err != nil {
		switch err.(type) {
		case *LensNotEqualError:
			doc, err = b.InfoAboutMovingRequest.Send(ctx, number, containerNumber)
			infoAboutMoving, err = b.InfoAboutMovingParser.Get(doc, containerNumber)
			if err != nil {
				infoAboutMoving = []*tracking.Event{}
			}
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
