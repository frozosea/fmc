package maeu

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type ContainerTracker struct {
	request               *Request
	etaParser             *EtaParser
	infoAboutMovingParser *InfoAboutMovingParser
	containerSizeParser   *ContainerSizeParser
}

func NewContainerTracker(b *tracking.BaseConstructorArgumentsForTracker) *ContainerTracker {
	return &ContainerTracker{request: NewRequest(b.Request, b.UserAgentGenerator), etaParser: NewEtaParser(b.Datetime), containerSizeParser: NewContainerSizeParser(), infoAboutMovingParser: NewInfoAboutMovingParser(b.Datetime)}
}

func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	apiResponse, err := c.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}
	infoAboutMoving := c.infoAboutMovingParser.get(apiResponse)
	eta, err := c.etaParser.get(apiResponse)
	if err == nil {
		infoAboutMoving = append(infoAboutMoving, eta)
	}
	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            c.containerSizeParser.get(apiResponse),
		Scac:            "MAEU",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
