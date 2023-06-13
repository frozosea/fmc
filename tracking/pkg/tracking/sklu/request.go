package sklu

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang_tracking/pkg/tracking/util/requests"
	"time"
)

type IHeadersGeneratorForApiRequest interface {
	Generate() map[string]string
}

type HeadersGeneratorForApiRequest struct {
	generator requests.IUserAgentGenerator
}

func NewHeadersGeneratorForApiRequest(generator requests.IUserAgentGenerator) *HeadersGeneratorForApiRequest {
	return &HeadersGeneratorForApiRequest{generator: generator}
}
func (h *HeadersGeneratorForApiRequest) Generate() map[string]string {
	return map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh;q=0.5",
		"Connection":       "keep-alive",
		"Host":             "ebiz.sinokor.co.kr",
		"Referer":          "http://ebiz.sinokor.co.kr/Tracking",
		"X-Requested-With": "XMLHttpRequest",
	}
}

type IApiUrlGenerator interface {
	Generate(number string, year int) string
}

type UrlGeneratorForApiRequest struct {
}

func NewUrlGeneratorForApiRequest() *UrlGeneratorForApiRequest {
	return &UrlGeneratorForApiRequest{}
}

func (u *UrlGeneratorForApiRequest) Generate(number string, year int) string {
	return fmt.Sprintf("http://ebiz.sinokor.co.kr/Tracking/GetBLList?cntrno=%s&year=%d", number, year)
}

type ApiRequest struct {
	request          requests.IHttp
	urlGenerator     IApiUrlGenerator
	headersGenerator IHeadersGeneratorForApiRequest
}

func NewApiRequest(request requests.IHttp, urlGenerator IApiUrlGenerator, headersGenerator IHeadersGeneratorForApiRequest) *ApiRequest {
	return &ApiRequest{
		request:          request,
		urlGenerator:     urlGenerator,
		headersGenerator: headersGenerator,
	}
}
func (a *ApiRequest) send(ctx context.Context, billNo, containerNo string, year int) (*ApiResponse, error) {
	url := a.urlGenerator.Generate(containerNo, year)
	headers := a.headersGenerator.Generate()

	response, err := a.request.Url(url).Method("GET").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 210 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s []*ApiResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	if billNo != "" {
		for _, v := range s {
			if v.BKNO == billNo {
				return v, nil
			}
		}
	}
	if len(s) != 0 {
		return s[0], nil
	}
	return nil, errors.New("no len")
}
func (a *ApiRequest) Send(ctx context.Context, billNo, containerNo string) (*ApiResponse, error) {
	now := time.Now()
	s, err := a.send(ctx, billNo, containerNo, now.Year())
	if err != nil {
		if now.Month() > 4 {
			return nil, err
		}
		s, err = a.send(ctx, billNo, containerNo, now.Year()-1)
		if err != nil {
			return nil, err
		}
		return s, nil
	}
	return s, nil
}

type IHeadersGeneratorForInfoAboutMovingRequest interface {
	Generate(billNo, containerNo string) map[string]string
}

type HeadersGeneratorForInfoAboutMovingRequest struct {
	generator requests.IUserAgentGenerator
}

func NewHeadersGeneratorForInfoAboutMovingRequest(generator requests.IUserAgentGenerator) *HeadersGeneratorForInfoAboutMovingRequest {
	return &HeadersGeneratorForInfoAboutMovingRequest{generator: generator}
}
func (h *HeadersGeneratorForInfoAboutMovingRequest) Generate(billNo, containerNo string) map[string]string {
	return map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Encoding":           "gzip, deflate",
		"Accept-Language":           "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh;q=0.5",
		"Cache-Control":             "max-age=0",
		"Connection":                "keep-alive",
		"Host":                      "ebiz.sinokor.co.kr",
		"Referer":                   fmt.Sprintf("http://ebiz.sinokor.co.kr/Tracking?blno=%s&cntrno=%s", billNo, containerNo),
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                h.generator.Generate(),
	}
}

type IUrlGeneratorForInfoAboutMovingRequest interface {
	Generate(billNo, containerNo string) string
}

type UrlGeneratorForInfoAboutMovingRequest struct {
}

func NewUrlGeneratorForInfoAboutMovingRequest() *UrlGeneratorForInfoAboutMovingRequest {
	return &UrlGeneratorForInfoAboutMovingRequest{}
}

func (u *UrlGeneratorForInfoAboutMovingRequest) Generate(billNo, containerNo string) string {
	if containerNo != "" {
		return fmt.Sprintf(`http://ebiz.sinokor.co.kr/Tracking?blno=%s&cntrno=%s`, billNo, containerNo)
	}
	return fmt.Sprintf(`http://ebiz.sinokor.co.kr/Tracking?blno=%s&cntrno=`, billNo)
}

type InfoAboutMovingRequest struct {
	request          requests.IHttp
	headersGenerator IHeadersGeneratorForInfoAboutMovingRequest
	urlGenerator     IUrlGeneratorForInfoAboutMovingRequest
}

func NewInfoAboutMovingRequest(request requests.IHttp, urlGenerator IUrlGeneratorForInfoAboutMovingRequest, headersGenerator IHeadersGeneratorForInfoAboutMovingRequest) *InfoAboutMovingRequest {
	return &InfoAboutMovingRequest{
		request:          request,
		headersGenerator: headersGenerator,
		urlGenerator:     urlGenerator,
	}
}
func (i *InfoAboutMovingRequest) Send(ctx context.Context, billNo, containerNo string) (*goquery.Document, error) {
	url := i.urlGenerator.Generate(billNo, containerNo)
	headers := i.headersGenerator.Generate(billNo, containerNo)
	response, err := i.request.Url(url).Method("GET").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 250 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
	if err != nil {
		return nil, err
	}
	return doc, nil
}
