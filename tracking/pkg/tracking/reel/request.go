package reel

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang_tracking/pkg/tracking/util/requests"
)

type IHeadersGeneratorForBillRequest interface {
	Generate(billNo string) map[string]string
}

type HeadersGeneratorForBillRequest struct {
	generator requests.IUserAgentGenerator
}

func NewHeadersGeneratorForInfoAboutMovingRequest(generator requests.IUserAgentGenerator) IHeadersGeneratorForBillRequest {
	return &HeadersGeneratorForBillRequest{generator: generator}
}
func (h *HeadersGeneratorForBillRequest) Generate(billNo string) map[string]string {
	return map[string]string{
		"authority":          "tracking.reelshipping.com",
		"accept":             "*/*",
		"accept-language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"content-type":       "application/x-www-form-urlencoded; charset=UTF-8",
		"origin":             "https://tracking.reelshipping.com",
		"referer":            fmt.Sprintf("https://tracking.reelshipping.com/tracking/?BlNo=%s&keyCode=&Id1=", billNo),
		"sec-ch-ua":          "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"macOS\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"x-requested-with":   "XMLHttpRequest",
		"User-Agent":         h.generator.Generate(),
	}
}

type IUrlGeneratorForBillRequest interface {
	Generate(billNo string) string
}

type UrlGeneratorForBillRequest struct {
}

func NewUrlGeneratorForBillRequest() IUrlGeneratorForBillRequest {
	return &UrlGeneratorForBillRequest{}
}

func (u *UrlGeneratorForBillRequest) Generate(billNo string) string {
	return fmt.Sprintf(`https://tracking.reelshipping.com/Tracking//TrackingView.asp?BlNo=%s&keyCode=&Id1=`, billNo)
}

type BillRequest struct {
	request          requests.IHttp
	headersGenerator IHeadersGeneratorForBillRequest
	urlGenerator     IUrlGeneratorForBillRequest
}

func NewBillRequest(request requests.IHttp, urlGenerator IUrlGeneratorForBillRequest, headersGenerator IHeadersGeneratorForBillRequest) *BillRequest {
	return &BillRequest{
		request:          request,
		headersGenerator: headersGenerator,
		urlGenerator:     urlGenerator,
	}
}
func (i *BillRequest) Send(ctx context.Context, billNo string) (*goquery.Document, error) {
	url := i.urlGenerator.Generate(billNo)
	headers := i.headersGenerator.Generate(billNo)
	response, err := i.request.Url(url).Method("POST").Headers(headers).Do(ctx)
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

type IUrlGeneratorForContainerInfo interface {
	Generate(billNo, containerNo string) string
}

type UrlGeneratorForContainerInfo struct {
}

func NewUrlGeneratorForContainerInfo() *UrlGeneratorForContainerInfo {
	return &UrlGeneratorForContainerInfo{}
}

func (u *UrlGeneratorForContainerInfo) Generate(billNo, containerNo string) string {
	return fmt.Sprintf("https://tracking.reelshipping.com/Tracking//ContainerStatus.asp?button=View&ContainerId=%s&BookingId=%s&ViewContainer=CONTAINER1", containerNo, billNo)
}

type ContainerInfoRequest struct {
	request          requests.IHttp
	headersGenerator IHeadersGeneratorForBillRequest
	urlGenerator     IUrlGeneratorForContainerInfo
}

func NewContainerInfoRequest(request requests.IHttp, headersGenerator IHeadersGeneratorForBillRequest, urlGenerator IUrlGeneratorForContainerInfo) *ContainerInfoRequest {
	return &ContainerInfoRequest{request: request, headersGenerator: headersGenerator, urlGenerator: urlGenerator}
}

func (c *ContainerInfoRequest) Send(ctx context.Context, billNo, containerNo string) (*goquery.Document, error) {
	url := c.urlGenerator.Generate(billNo, containerNo)
	headers := c.headersGenerator.Generate(billNo)
	response, err := c.request.Url(url).Method("GET").Headers(headers).Do(ctx)
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
