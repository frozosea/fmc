package maeu

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
		"authority":          "api.maersk.com",
		"accept":             "application/json",
		"accept-language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"origin":             "https://www.maersk.com.cn",
		"referer":            "https://www.maersk.com.cn/",
		"sec-ch-ua":          "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"macOS\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "cross-site",
		"user-agent":         h.userAgentGenerator.Generate(),
	}
}

type urlGenerator struct {
}

func newUrlGenerator() *urlGenerator {
	return &urlGenerator{}
}

func (u *urlGenerator) generate(number string) string {
	return fmt.Sprintf(`https://api.maersk.com/track/%s?operator=MAEU`, number)
}

type Request struct {
	request          requests.IHttp
	headersGenerator *headersGenerator
	urlGenerator     *urlGenerator
}

func NewRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *Request {
	return &Request{
		request:          request,
		headersGenerator: newHeadersGenerator(generator),
		urlGenerator:     newUrlGenerator(),
	}
}

func (r *Request) Send(ctx context.Context, number string) (*ApiResponse, error) {
	url := r.urlGenerator.generate(number)
	headers := r.headersGenerator.generate()
	response, err := r.request.Url(url).Method("GET").Headers(headers).Do(ctx)
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
	if len(s.Containers) != 0 {
		return s, nil
	}
	return nil, errors.New("no len")
}
