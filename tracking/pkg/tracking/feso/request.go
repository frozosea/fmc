package feso

import (
	"context"
	"encoding/json"
	"errors"
	"golang_tracking/pkg/tracking/util/requests"
	"strings"
)

type requestBodyGenerator struct {
}

func newRequestBodyGenerator() *requestBodyGenerator {
	return &requestBodyGenerator{}
}

func (b *requestBodyGenerator) Gen(numbers []string) ([]byte, error) {
	var r struct {
		Codes    []string    `json:"codes"`
		Email    interface{} `json:"email"`
		ForDate  interface{} `json:"forDate"`
		FromFile bool        `json:"fromFile"`
	}
	r.Codes = numbers
	r.Email = nil
	r.ForDate = nil
	r.FromFile = false
	return json.Marshal(r)
}

type headersGenerator struct {
	uAgentGen requests.IUserAgentGenerator
}

func newHeadersGenerator(uAgentGen requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{uAgentGen: uAgentGen}
}

func (h *headersGenerator) Generate() map[string]string {
	return map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"content-type":       "application/json",
		"sec-ch-ua":          "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"macOS\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "cross-site",
		"Referer":            "https://www.fesco.ru/",
		"Referrer-Policy":    "strict-origin-when-cross-origin",
		"User-Agent":         h.uAgentGen.Generate(),
	}
}

type Request struct {
	bodyGen *requestBodyGenerator
	hGen    *headersGenerator
	request requests.IHttp
}

func NewFesoRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *Request {
	return &Request{bodyGen: newRequestBodyGenerator(), hGen: newHeadersGenerator(generator), request: request}
}
func (f *Request) Send(ctx context.Context, number string) (*ResponseSchema, error) {
	const fesoApiUrl = "https://tracking.fesco.com/api/v1/tracking/Get"
	body, err := f.bodyGen.Gen([]string{number})
	if err != nil {
		return nil, err
	}

	response, err := f.request.Url(fesoApiUrl).Method("POST").Body(body).Headers(f.hGen.Generate()).Do(ctx)
	if err != nil {
		return nil, err
	}

	if response.Status > 300 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var rawBody struct {
		RequestKey string
		Containers []string
		Missing    []string
	}
	if err := json.Unmarshal(response.Body, &rawBody); err != nil {
		return nil, err
	}

	if len(rawBody.Containers) == 0 {
		return nil, errors.New("no len")
	}

	for index, item := range rawBody.Containers {
		rawBody.Containers[index] = strings.Replace(item, `\`, ``, -1)
	}
	var b *ResponseSchema
	if err := json.Unmarshal([]byte(rawBody.Containers[0]), &b); err != nil {
		return nil, err
	}
	if len(b.LastEvents) == 0 {
		return nil, errors.New("no len")
	}
	return b, nil
}
