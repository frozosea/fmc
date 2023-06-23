package sklu

import (
	"context"
	"golang_tracking/pkg/tracking"
	"strings"
)

type ContainerTracker struct {
	ApiRequest             *ApiRequest
	InfoAboutMovingRequest *InfoAboutMovingRequest
	ApiParser              *ApiParser
	InfoAboutMovingParser  *InfoAboutMovingParser
	UnlocodesRepo          IRepository
}

func NewContainerTracker(cfg *tracking.BaseConstructorArgumentsForTracker, repo IRepository) *ContainerTracker {
	return &ContainerTracker{
		ApiRequest:             NewApiRequest(cfg.Request, NewUrlGeneratorForApiRequest(), NewHeadersGeneratorForApiRequest(cfg.UserAgentGenerator)),
		InfoAboutMovingRequest: NewInfoAboutMovingRequest(cfg.Request, NewUrlGeneratorForInfoAboutMovingRequest(), NewHeadersGeneratorForInfoAboutMovingRequest(cfg.UserAgentGenerator)),
		ApiParser:              NewApiParser(cfg.Datetime),
		InfoAboutMovingParser:  NewInfoAboutMovingParser(cfg.Datetime),
		UnlocodesRepo:          repo,
	}
}

func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	apiResponse, err := c.ApiRequest.Send(ctx, "", number)
	if err != nil {
		return nil, err
	}

	containerInfo := c.ApiParser.Get(apiResponse)

	infoAboutMovingDoc, err := c.InfoAboutMovingRequest.Send(ctx, containerInfo.BillNo, number)
	if err != nil {
		return nil, err
	}

	infoAboutMoving, err := c.InfoAboutMovingParser.Get(infoAboutMovingDoc, number)
	
	podFullName, err := c.UnlocodesRepo.GetFullName(ctx, containerInfo.Unlocode)
	if err != nil {
		podFullName = containerInfo.Unlocode
	}

	infoAboutMoving = append(infoAboutMoving, &tracking.Event{
		Time:          containerInfo.Eta,
		OperationName: "ETA",
		Location:      strings.ToUpper(podFullName),
		Vessel:        "",
	})

	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            strings.ToUpper(containerInfo.ContainerSize),
		Scac:            "SKLU",
		InfoAboutMoving: infoAboutMoving,
	}, nil
}
