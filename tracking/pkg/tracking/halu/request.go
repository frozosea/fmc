package halu

import (
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
)

type UrlGeneratorForApiRequest struct {
}

func NewUrlGeneratorForApiRequest() *UrlGeneratorForApiRequest {
	return &UrlGeneratorForApiRequest{}
}

func (u *UrlGeneratorForApiRequest) Generate(number string, year int) string {
	return fmt.Sprintf("http://ebiz.heung-a.com/Tracking/GetBLList?cntrno=%s&year=%d", number, year)
}

type HeadersGeneratorForApiRequest struct {
	generator requests.IUserAgentGenerator
}

func NewHeadersGeneratorForApiRequest(generator requests.IUserAgentGenerator) *HeadersGeneratorForApiRequest {
	return &HeadersGeneratorForApiRequest{generator: generator}
}

func (h *HeadersGeneratorForApiRequest) Generate() map[string]string {
	return map[string]string{
		"accept":           "application/json, text/javascript, */*; q=0.01",
		"accept-language":  "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"x-requested-with": "XMLHttpRequest",
		"Referer":          "http://ebiz.heung-a.com/Tracking",
		"Referrer-Policy":  "strict-origin-when-cross-origin",
		"User-Agent":       h.generator.Generate(),
	}
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

type HeadersGeneratorForInfoAboutMovingRequest struct {
	generator requests.IUserAgentGenerator
}

func NewHeadersGeneratorForInfoAboutMovingRequest(generator requests.IUserAgentGenerator) *HeadersGeneratorForInfoAboutMovingRequest {
	return &HeadersGeneratorForInfoAboutMovingRequest{generator: generator}
}
func (h *HeadersGeneratorForInfoAboutMovingRequest) Generate(billNo, containerNo string) map[string]string {
	return map[string]string{
		"Accept":           "application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding":  "gzip, deflate",
		"Accept-Language":  "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh;q=0.5",
		"Connection":       "keep-alive",
		"x-requested-with": "XMLHttpRequest",
		"Referer":          fmt.Sprintf("http://ebiz.heung-a.com/Tracking?blno=%s&cntrno=%s", billNo, containerNo),
		"Referrer-Policy":  "strict-origin-when-cross-origin",
		"User-Agent":       h.generator.Generate(),
	}
}
