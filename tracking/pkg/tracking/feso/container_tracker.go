package feso

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type ContainerTracker struct {
	request               *Request
	infoAboutMovingParser *InfoAboutMovingParser
	containerSizeParser   *ContainerSizeParser
}

func NewContainerTracker(baseRequestConfig *tracking.BaseConstructorArgumentsForTracker) *ContainerTracker {
	return &ContainerTracker{
		request:               NewFesoRequest(baseRequestConfig.Request, baseRequestConfig.UserAgentGenerator),
		infoAboutMovingParser: NewInfoAboutMovingParser(baseRequestConfig.Datetime),
		containerSizeParser:   NewContainerSizeParser(),
	}
}

func (f *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	response, err := f.request.Send(ctx, number)
	if err != nil {
		return nil, err
	}
	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            f.containerSizeParser.Get(response),
		Scac:            "FESO",
		InfoAboutMoving: f.infoAboutMovingParser.get(response.LastEvents),
	}, nil
}
