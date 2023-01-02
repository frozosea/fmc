package cosu

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type ContainerTracker struct {
	request               IRequest
	podParser             *PodParser
	containerSizeParser   *ContainerSizeParser
	etaParser             *EtaParser
	infoAboutMovingParser *InfoAboutMovingParser
}

func NewContainerTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *ContainerTracker {
	return &ContainerTracker{
		request:               NewRequest(cfg.Request, cfg.UserAgentGenerator),
		podParser:             NewPodParser(),
		containerSizeParser:   NewContainerSizeParser(),
		etaParser:             NewEtaParser(cfg.Datetime),
		infoAboutMovingParser: NewInfoAboutMovingParser(cfg.Datetime),
	}
}

func (c *ContainerTracker) getInfoAboutMoving(ctx context.Context, number string) ([]*tracking.Event, *ApiResponseSchema, error) {
	apiResponse, err := c.request.GetInfoAboutMovingRawResponse(ctx, number)
	if err != nil {
		return nil, nil, err
	}
	return c.infoAboutMovingParser.get(apiResponse), apiResponse, nil
}
func (c *ContainerTracker) getEta(ctx context.Context, number string, apiResponse *ApiResponseSchema) (*tracking.Event, error) {
	rawETAResponse, err := c.request.GetEtaRawResponse(ctx, number)
	if err != nil {
		return nil, err
	}
	eta, err := c.etaParser.get(rawETAResponse, c.podParser.get(apiResponse))
	if err != nil {
		return nil, err
	}
	return eta, nil
}
func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	infoAboutMoving, apiResponse, err := c.getInfoAboutMoving(ctx, number)
	if err != nil {
		return nil, err
	}
	eta, err := c.getEta(ctx, number, apiResponse)
	if err != nil {
		return nil, err
	}
	infoAboutMoving = append(infoAboutMoving, eta)
	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            c.containerSizeParser.get(apiResponse),
		Scac:            "COSU",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
