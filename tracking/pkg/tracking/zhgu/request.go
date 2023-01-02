package zhgu

import (
	"context"
	"encoding/json"
	"golang_tracking/pkg/tracking/util/requests"
)

type headersGenerator struct {
	generator requests.IUserAgentGenerator
}

func newHeadersGenerator(generator requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{generator: generator}
}

func (h *headersGenerator) Generate() map[string]string {
	return map[string]string{
		"accept":          "application/json, text/plain, */*",
		"accept-language": "en",
		"content-type":    "application/json;charset=UTF-8",
		"Referer":         "http://elines.zhonggu56.com/track",
		"Referrer-Policy": "strict-origin-when-cross-origin",
		"User-Agent":      h.generator.Generate(),
	}
}

type bodyGenerator struct {
}

func newBodyGenerator() *bodyGenerator {
	return &bodyGenerator{}
}

func (b *bodyGenerator) Generate(number string) ([]byte, error) {
	var s struct {
		BlNo string `json:"blNo"`
	}
	s.BlNo = number
	return json.Marshal(s)
}

func requestWrap(ctx context.Context, number string, request requests.IHttp, headersGenerator *headersGenerator, bodyGenerator *bodyGenerator) ([]byte, error) {
	const url = "http://elines.zhonggu56.com/api/booking/getVoyageInfo"
	headers := headersGenerator.Generate()
	body, err := bodyGenerator.Generate(number)
	if err != nil {
		return nil, err
	}

	response, err := request.Url(url).Method("POST").Headers(headers).Body(body).Do(ctx)
	if err != nil {
		return nil, err
	}

	if response.Status > 210 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	return response.Body, nil
}

type ApiRequest struct {
	request          requests.IHttp
	headersGenerator *headersGenerator
	bodyGenerator    *bodyGenerator
}

func NewApiRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *ApiRequest {
	return &ApiRequest{
		request:          request,
		headersGenerator: newHeadersGenerator(generator),
		bodyGenerator:    newBodyGenerator(),
	}
}

func (a *ApiRequest) Send(ctx context.Context, number string) (*ApiResponseSchema, error) {
	r, err := requestWrap(ctx, number, a.request, a.headersGenerator, a.bodyGenerator)
	if err != nil {
		return nil, err
	}

	var s *ApiResponseSchema
	if err := json.Unmarshal(r, &s); err != nil {
		return nil, err
	}
	return s, nil
}

type BookingApiRequest struct {
	request          requests.IHttp
	headersGenerator *headersGenerator
	bodyGenerator    *bodyGenerator
}

func NewBookingApiRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *BookingApiRequest {
	return &BookingApiRequest{
		request:          request,
		headersGenerator: newHeadersGenerator(generator),
		bodyGenerator:    newBodyGenerator(),
	}
}

func (b *BookingApiRequest) Send(ctx context.Context, number string) (*BookingApiResponseSchema, error) {
	r, err := requestWrap(ctx, number, b.request, b.headersGenerator, b.bodyGenerator)
	if err != nil {
		return nil, err
	}

	var s *BookingApiResponseSchema
	if err := json.Unmarshal(r, &s); err != nil {
		return nil, err
	}
	return s, nil
}
