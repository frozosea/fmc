package oney

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type ContainerTracker struct {
	request               *Request
	copNoAndBkgNoParser   *CopNoAndBkgNoParser
	containerSizeParser   *ContainerSizeParser
	infoAboutMovingParser *InfoAboutMovingParser
}

func NewContainerTracker(b *tracking.BaseConstructorArgumentsForTracker) *ContainerTracker {
	return &ContainerTracker{
		request:               NewRequest(b.Request, b.UserAgentGenerator),
		copNoAndBkgNoParser:   NewCopNoAndBkgNoParser(),
		containerSizeParser:   NewContainerSizeParser(),
		infoAboutMovingParser: NewInfoAboutMovingParser(b.Datetime),
	}
}

func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	if isAccessoryThisLine, err := c.request.CheckContainerAccessoryToThisLine(ctx, number); err != nil || !isAccessoryThisLine {
		return nil, tracking.NewNotThisLineException()
	}
	copAndBillNosApiResponse, err := c.request.SendRequestForGetBkgNoAndCopNo(ctx, number)
	if err != nil {
		return nil, err
	}
	copNo, bkgNo := c.copNoAndBkgNoParser.get(copAndBillNosApiResponse)
	infoAboutMovingApiResponse, err := c.request.SendForInfoAboutMoving(ctx, number, string(bkgNo), string(copNo))
	if err != nil {
		return nil, err
	}
	infoAboutMoving := c.infoAboutMovingParser.get(infoAboutMovingApiResponse)
	containerSizeApiResponse, err := c.request.SendForContainerSize(ctx, number, string(bkgNo), string(copNo))
	if err != nil {
		return nil, err
	}
	containerSize := c.containerSizeParser.get(containerSizeApiResponse)
	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            containerSize,
		Scac:            "ONEY",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
