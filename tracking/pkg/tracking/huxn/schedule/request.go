package huaxin_schedule

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
	"sync"
	"time"
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

type Request struct {
	headersGenerator *headersGenerator
	request          requests.IHttp
}

func NewRequest(request requests.IHttp, uaGenerator requests.IUserAgentGenerator) *Request {
	return &Request{request: request, headersGenerator: newHeadersGenerator(uaGenerator)}
}

func (c *Request) GetWholeWorldSchedule(ctx context.Context, portLoadUnlocode string, portDiscUnlocodes []string, etd time.Time) ([]*Schedule, error) {
	var schedules []*Schedule

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, unlocode := range portDiscUnlocodes {
		wg.Add(1)

		go func(ctx context.Context, unlocode string) {
			defer wg.Done()
			response, err := c.Send(ctx, portLoadUnlocode, unlocode, etd)
			if err != nil || len(response.ListSchedules) == 0 {
				return
			}
			defer mu.Unlock()
			mu.Lock()
			schedules = append(schedules, response.ListSchedules...)
		}(ctx, unlocode)

		wg.Wait()
	}

	return schedules, nil
}

func (c *Request) Send(ctx context.Context, portLoadUnlocode, portDiscUnlocode string, etd time.Time) (*ServerResponse, error) {
	const url = "http://dc.hxlines.com:8099/HX_WeChat/SearchSchedules"
	headers := c.headersGenerator.Generate()
	form := map[string]string{
		"strPortLoad": portLoadUnlocode,
		"strPortDisc": portDiscUnlocode,
		"strEtd":      etd.Format("01/02/2006"),
		"strWeeks":    "2",
		"strDirect":   "false",
	}

	response, err := c.request.Url(url).Method("POST").Headers(headers).Form(form).Do(ctx)
	if err != nil || response.Body == nil {
		return nil, err
	}

	if response.Status > 200 {
		return nil, errors.New(fmt.Sprintf("get schedule HuaXin line status code: %d", response.Status))
	}

	var t *ServerResponse

	if err := json.Unmarshal(response.Body, &t); err != nil {
		return nil, err
	}
	return t, nil
}
