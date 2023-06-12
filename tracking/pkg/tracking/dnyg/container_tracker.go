package dnyg

import (
	"context"
	"golang_tracking/pkg/tracking"
)

type ContainerTracker struct {
	numberInfoRequest      INumberInfoRequest
	infoAboutMovingRequest IInfoAboutMovingRequest
	infoAboutMovingParser  *InfoAboutMovingParser
	etaParser              *EtaParser
	podParser              *PODParser
}

func NewContainerTracker(cfg *tracking.BaseConstructorArgumentsForTracker) *ContainerTracker {
	return &ContainerTracker{
		numberInfoRequest:      NewNumberInfoRequest(cfg.Request, cfg.UserAgentGenerator),
		infoAboutMovingRequest: NewInfoAboutMovingRequest(cfg.Request, cfg.UserAgentGenerator),
		infoAboutMovingParser:  NewInfoAboutMovingParser(cfg.Datetime),
		etaParser:              NewEtaParser(cfg.Datetime),
		podParser:              NewPODParser(),
	}
}

func (c *ContainerTracker) Track(ctx context.Context, number string) (*tracking.ContainerTrackingResponse, error) {
	containerNumberInfo, err := c.numberInfoRequest.Send(ctx, number, true)
	if err != nil {
		return nil, err
	}

	infoAboutMovingRawResponse, err := c.infoAboutMovingRequest.Send(ctx, containerNumberInfo.DltResultBlList[0].OUTBKN, containerNumberInfo.DltResultBlList[0].OUTBNO, containerNumberInfo.DltResultBlList[0].OUTCNT)
	if err != nil {
		return nil, err
	}

	infoAboutMoving := c.infoAboutMovingParser.Parse(infoAboutMovingRawResponse)

	eta, err := c.etaParser.parse(containerNumberInfo)
	if err == nil {
		infoAboutMoving = append(infoAboutMoving, &tracking.Event{
			Time:          eta,
			OperationName: "ETA",
			Location:      c.podParser.parse(containerNumberInfo),
			Vessel:        "",
		})
	}

	return &tracking.ContainerTrackingResponse{
		Number:          number,
		Size:            containerNumberInfo.DltResultBlList[0].SIZETYPE,
		Scac:            "DNYG",
		InfoAboutMoving: infoAboutMoving,
	}, nil

}
