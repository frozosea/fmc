package huxn

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type IETAProvider interface {
	GetETA(ctx context.Context, data *TrackingResponse) ([]*tracking.Event, error)
}

type ContainerTracker struct {
	etaParser             IETAProvider
	request               *ContainerTrackingRequest
	infoAboutMovingParser *InfoAboutMovingParser
	containerSizeParser   *ContainerSizeParser
}

func NewContainerTracker(cfg *tracking.BaseConstructorArgumentsForTracker, etaProvider IETAProvider) *ContainerTracker {
	return &ContainerTracker{
		etaParser:             etaProvider,
		request:               NewContainerTrackingRequest(cfg.Request, cfg.UserAgentGenerator),
		infoAboutMovingParser: NewInfoAboutMovingParser(cfg.Datetime),
		containerSizeParser:   NewContainerSizeParser(),
	}
}
func (c *ContainerTracker) trackOldSystem(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	data, err := c.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}
	if len(data.ListDynamics) == 0 {
		return nil, tracking.NewNotThisLineException()
	}
	infoAboutMoving, _ := c.infoAboutMovingParser.Parse(data)

	if !checkNumberArrived(infoAboutMoving) {
		etaEvents, err := c.etaParser.GetETA(ctx, data)
		if err == nil {
			infoAboutMoving = append(infoAboutMoving, etaEvents...)
			if len(infoAboutMoving) > 0 {
				if infoAboutMoving[len(infoAboutMoving)-1].Time.Unix() < etaEvents[len(etaEvents)-1].Time.Unix() {
					infoAboutMoving = append(infoAboutMoving, etaEvents...)
				}
			}
		}
	}

	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            c.containerSizeParser.Parse(data),
		Scac:            "HUXN",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
func (c *ContainerTracker) trackNewSystem(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	return nil, tracking.NewNotThisLineException()
}
func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	data, err := c.trackOldSystem(ctx, number)
	if err == nil {
		return data, nil
	} else {
		return c.trackNewSystem(ctx, number)
	}
}
