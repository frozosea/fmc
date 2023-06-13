package sklu

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
	"strings"
)

type ICheckBookingNumberExistsUrlGenerator interface {
	GenerateUrl(number string) string
}

type CheckBookingNumberExistsUrlGenerator struct {
}

func NewCheckBookingNumberExistsUrlGenerator() *CheckBookingNumberExistsUrlGenerator {
	return &CheckBookingNumberExistsUrlGenerator{}
}

func (c *CheckBookingNumberExistsUrlGenerator) GenerateUrl(number string) string {
	return fmt.Sprintf("http://ebiz.sinokor.co.kr/Home/chkExistsBooking?bkno=%s", number)
}

type ICheckNumberExistsHeadersGenerator interface {
	GenerateHeaders(number string) map[string]string
}

type CheckNumberExistsHeadersGenerator struct {
}

func NewCheckNumberExistsHeadersGenerator() *CheckNumberExistsHeadersGenerator {
	return &CheckNumberExistsHeadersGenerator{}
}

func (c *CheckNumberExistsHeadersGenerator) GenerateHeaders(number string) map[string]string {
	return map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Language":  "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"Connection":       "keep-alive",
		"Referer":          fmt.Sprintf("http://ebiz.sinokor.co.kr/Tracking?blno=%s&cntrno=", number),
		"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
	}
}

type ICheckBookingNumberExists interface {
	CheckExists(ctx context.Context, number string) (bool, error)
}
type CheckBookingNumberExistsResponse struct {
	STATUS string `json:"STATUS"`
	MSG    string `json:"MSG"`
}
type CheckBookingNumberExists struct {
	request          requests.IHttp
	urlGenerator     ICheckBookingNumberExistsUrlGenerator
	headersGenerator ICheckNumberExistsHeadersGenerator
}

func NewCheckBookingNumberExists(request requests.IHttp, urlGenerator ICheckBookingNumberExistsUrlGenerator, headersGenerator ICheckNumberExistsHeadersGenerator) *CheckBookingNumberExists {
	return &CheckBookingNumberExists{request: request, urlGenerator: urlGenerator, headersGenerator: headersGenerator}
}

func (c *CheckBookingNumberExists) CheckExists(ctx context.Context, number string) (bool, error) {
	url := c.urlGenerator.GenerateUrl(number)
	headers := c.headersGenerator.GenerateHeaders(number)

	response, err := c.request.Method("GET").Url(url).Headers(headers).Do(ctx)
	if err != nil {
		return false, err
	}
	var r *CheckBookingNumberExistsResponse
	if err := json.Unmarshal(response.Body, &r); err != nil {
		return false, err
	}
	if strings.EqualFold(r.STATUS, "N") || strings.EqualFold(r.STATUS, "Can't find Booking that you input.") {
		return false, nil
	}
	return true, nil

}
