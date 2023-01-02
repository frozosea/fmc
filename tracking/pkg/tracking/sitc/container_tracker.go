package sitc

import (
	"context"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/sitc/login_provider"
)

type ContainerTracker struct {
	request *ContainerTrackingRequest
	parser  *ContainerTrackingInfoAboutMovingParser
}

func NewContainerTracker(b *tracking.BaseConstructorArgumentsForTracker, store *login_provider.Store) *ContainerTracker {
	return &ContainerTracker{
		request: NewContainerTrackingRequest(b.Request, b.UserAgentGenerator, store),
		parser:  NewContainerTrackingInfoAboutMovingParser(b.Datetime),
	}
}

func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	response, err := c.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}
	infoAboutMoving := c.parser.get(response)
	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            "",
		Scac:            "SITC",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
