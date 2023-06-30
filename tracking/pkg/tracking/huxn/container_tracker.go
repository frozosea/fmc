package huxn

import (
	"context"
	"errors"
	"golang_tracking/pkg/tracking"
	"strings"
)

type IETAProvider interface {
	GetETA(ctx context.Context, data *ContainerTrackingResponse) ([]*tracking.Event, error)
}

type ContainerTracker struct {
	etaParser             IETAProvider
	request               *ContainerTrackingRequest
	infoAboutMovingParser *ContainerTrackerInfoAboutMovingParser
	containerSizeParser   *ContainerSizeParser
}

func NewContainerTracker(cfg *tracking.BaseConstructorArgumentsForTracker, etaProvider IETAProvider) *ContainerTracker {
	return &ContainerTracker{
		etaParser:             etaProvider,
		request:               NewContainerTrackingRequest(cfg.Request, cfg.UserAgentGenerator),
		infoAboutMovingParser: NewContainerTrackerInfoAboutMovingParser(cfg.Datetime),
		containerSizeParser:   NewContainerSizeParser(),
	}
}
func (c *ContainerTracker) checkContainerArrived(infoAboutMoving []*tracking.Event) bool {
	for _, item := range infoAboutMoving {
		if strings.EqualFold(item.OperationName, "Discharge from vessel full") {
			return true
		}
	}
	return false
}
func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	data, err := c.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}
	if len(data.ListDynamics) == 0 {
		return nil, errors.New("huaxin length is zero")
	}
	infoAboutMoving, _ := c.infoAboutMovingParser.Parse(data)

	if !c.checkContainerArrived(infoAboutMoving) {
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
