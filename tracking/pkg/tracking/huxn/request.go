package huxn

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
)

type headersGenerator struct {
	uaGenerator requests.IUserAgentGenerator
}

func newHeadersGenerator(uaGenerator requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{uaGenerator: uaGenerator}
}

func (h *headersGenerator) Generate() map[string]string {
	return map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Language":  "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"Connection":       "keep-alive",
		"Content-Type":     "application/x-www-form-urlencoded; charset=UTF-8",
		"Origin":           "http://dc.hxlines.com:8099",
		"Referer":          "http://dc.hxlines.com:8099/HX_WeChat/HX_Dynamics",
		"User-Agent":       h.uaGenerator.Generate(),
		"X-Requested-With": "XMLHttpRequest",
	}
}

type ContainerTrackingRequest struct {
	headersGenerator *headersGenerator
	request          requests.IHttp
}

func NewContainerTrackingRequest(request requests.IHttp, uaGenerator requests.IUserAgentGenerator) *ContainerTrackingRequest {
	return &ContainerTrackingRequest{request: request, headersGenerator: newHeadersGenerator(uaGenerator)}
}

func (c *ContainerTrackingRequest) Send(ctx context.Context, number string) (*TrackingResponse, error) {
	const url = "http://dc.hxlines.com:8099/HX_WeChat/SearchDynamics"
	headers := c.headersGenerator.Generate()
	form := map[string]string{
		"strNumbers": number,
	}

	response, err := c.request.Url(url).Method("POST").Headers(headers).Form(form).Do(ctx)
	if err != nil {
		return nil, err
	}

	if response.Status > 200 {
		return nil, errors.New(fmt.Sprintf("get container info HuaXin line status code: %d", response.Status))
	}

	var t *TrackingResponse

	if err := json.Unmarshal(response.Body, &t); err != nil {
		return nil, err
	}
	return t, nil
}

type ScheduleRequest struct {
	headersGenerator *headersGenerator
	request          requests.IHttp
}
