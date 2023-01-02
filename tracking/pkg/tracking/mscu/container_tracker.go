package mscu

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type ContainerTracker struct {
	request               *Request
	etaParser             *EtaParser
	containerSizeParser   *ContainerSizeParser
	infoAboutMovingParser *InfoAboutMovingParser
}

func NewContainerTracker(b *tracking.BaseConstructorArgumentsForTracker) *ContainerTracker {
	return &ContainerTracker{
		request:               NewRequest(b.Request, b.UserAgentGenerator),
		infoAboutMovingParser: NewInfoAboutMovingParser(b.Datetime),
		etaParser:             NewEtaParser(b.Datetime),
		containerSizeParser:   NewContainerSizeParser(),
	}
}

func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	response, err := c.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}
	infoAboutMoving := c.infoAboutMovingParser.get(response)
	eta, err := c.etaParser.get(response)
	if err == nil {
		infoAboutMoving = append(infoAboutMoving, eta)
	}
	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            c.containerSizeParser.get(response),
		Scac:            "MSCU",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
