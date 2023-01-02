package zhgu

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type BillTracker struct {
	apiRequest            *ApiRequest
	bookingApiRequest     *BookingApiRequest
	etaParser             *EtaParser
	infoAboutMovingParser *InfoAboutMovingParser
}

func NewBillTracker(args *tracking.BaseConstructorArgumentsForTracker) *BillTracker {
	return &BillTracker{
		apiRequest:            NewApiRequest(args.Request, args.UserAgentGenerator),
		bookingApiRequest:     NewBookingApiRequest(args.Request, args.UserAgentGenerator),
		etaParser:             NewEtaParser(args.Datetime),
		infoAboutMovingParser: NewInfoAboutMovingParser(args.Datetime),
	}
}

func (b *BillTracker) Track(ctx context.Context, number string) (*tracking.BillNumberTrackingResponse, error) {
	book, err := b.bookingApiRequest.Send(ctx, number)
	if err != nil {
		return nil, err
	}

	if len(book.Object) == 0 {
		return nil, tracking.NewNotThisLineException()
	}

	response, err := b.apiRequest.Send(ctx, number)
	if err != nil {
		return nil, err
	}

	infoAboutMoving, err := b.infoAboutMovingParser.Get(response)
	if err != nil {
		return nil, err
	}

	eta, err := b.etaParser.Get(response)
	if err != nil {
		return nil, err
	}

	return &tracking.BillNumberTrackingResponse{
		Number:          number,
		Eta:             eta,
		Scac:            "ZHGU",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
