package halu

import (
	"context"
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/sklu"
)

type ContainerTracker struct {
	*sklu.ContainerTracker
}

func NewContainerTracker(cfg *tracking.BaseConstructorArgumentsForTracker, repo sklu.IRepository) *ContainerTracker {
	return &ContainerTracker{
		ContainerTracker: &sklu.ContainerTracker{
			ApiRequest:             sklu.NewApiRequest(cfg.Request, NewUrlGeneratorForApiRequest(), NewHeadersGeneratorForApiRequest(cfg.UserAgentGenerator)),
			InfoAboutMovingRequest: sklu.NewInfoAboutMovingRequest(cfg.Request, NewUrlGeneratorForInfoAboutMovingRequest(), NewHeadersGeneratorForInfoAboutMovingRequest(cfg.UserAgentGenerator)),
			ApiParser:              sklu.NewApiParser(cfg.Datetime),
			InfoAboutMovingParser:  sklu.NewInfoAboutMovingParser(cfg.Datetime),
			UnlocodesRepo:          repo,
		},
	}
}

func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	response, err := c.ContainerTracker.Track(ctx, number)
	if err != nil {
		return nil, err
	}
	response.Scac = "HALU"
	return response, nil
}
