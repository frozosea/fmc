package oney

import (
	"golang_tracking/pkg/tracking"
	"golang_tracking/pkg/tracking/util/datetime"
	"strings"
)

type CopNo string

type BkgNo string

type CopNoAndBkgNoParser struct {
}

func NewCopNoAndBkgNoParser() *CopNoAndBkgNoParser {
	return &CopNoAndBkgNoParser{}
}
func (c *CopNoAndBkgNoParser) get(apiResp *BkgAndCopNosApiResponseSchema) (CopNo, BkgNo) {
	return CopNo(apiResp.List[0].CopNo), BkgNo(apiResp.List[0].BkgNo)
}

type ContainerSizeParser struct {
}

func NewContainerSizeParser() *ContainerSizeParser {
	return &ContainerSizeParser{}
}
func (c *ContainerSizeParser) get(apiResp *ContainerSizeApiResponseSchema) string {
	return apiResp.List[0].CntrTpszNm
}

type InfoAboutMovingParser struct {
	datetime datetime.IDatetime
}

func NewInfoAboutMovingParser(datetime datetime.IDatetime) *InfoAboutMovingParser {
	return &InfoAboutMovingParser{datetime: datetime}
}

func (i *InfoAboutMovingParser) get(apiResp *InfoAboutMovingApiResponseSchema) []*tracking.Event {
	var infoAboutMoving []*tracking.Event
	for _, item := range apiResp.List {
		//YYYY-MM-DD HH:mm
		eventTime, err := i.datetime.Strptime(item.EventDt, "%Y-%m-%d %H:%M")
		if err != nil {
			continue
		}
		infoAboutMoving = append(infoAboutMoving, &tracking.Event{
			Time:          eventTime,
			OperationName: strings.ToUpper(strings.ToLower(strings.Trim(item.StatusNm, " "))),
			Location:      strings.ToUpper(strings.Trim(item.PlaceNm, " ")),
			Vessel:        strings.ToUpper(strings.Trim(item.VslEngNm, " ")),
		})
	}
	return infoAboutMoving
}
