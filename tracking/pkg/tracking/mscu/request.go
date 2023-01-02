package mscu

import (
	"context"
	"encoding/json"
	"golang_tracking/pkg/tracking/util/requests"
)

type headersGenerator struct {
	userAgentGenerator requests.IUserAgentGenerator
}

func newHeadersGenerator(userAgentGenerator requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{userAgentGenerator: userAgentGenerator}
}

func (h *headersGenerator) generate() map[string]string {
	return map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"content-type":       "application/json",
		"sec-ch-ua":          "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Google Chrome\";v=\"101\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"macOS\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"x-requested-with":   "XMLHttpRequest",
		"user-agent":         h.userAgentGenerator.Generate(),
	}
}

type body struct {
	TrackingNumber string `json:"trackingNumber"`
	TrackingMode   int    `json:"trackingMode"`
}

type bodyGenerator struct {
}

func newBodyGenerator() *bodyGenerator {
	return &bodyGenerator{}
}

func (b *bodyGenerator) generate(number string) ([]byte, error) {
	return json.Marshal(&body{
		TrackingNumber: number,
		TrackingMode:   0,
	})
}

type Request struct {
	request          requests.IHttp
	bodyGenerator    *bodyGenerator
	headersGenerator *headersGenerator
}

func NewRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *Request {
	return &Request{request: request, bodyGenerator: newBodyGenerator(), headersGenerator: newHeadersGenerator(generator)}
}

func (r *Request) Send(ctx context.Context, number string) (*ApiResponse, error) {
	const url = "https://www.msc.com/api/feature/tools/TrackingInfo"
	b, err := r.bodyGenerator.generate(number)
	if err != nil {
		return nil, err
	}
	response, err := r.request.Url(url).Method("POST").Headers(r.headersGenerator.generate()).Body(b).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 300 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s *ApiResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	return s, nil
}
