package login_provider

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
)

type headersGenerator struct {
	userAgentGenerator requests.IUserAgentGenerator
}

func newHeadersGenerator(userAgentGenerator requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{userAgentGenerator: userAgentGenerator}
}

func (h *headersGenerator) generate(basicAuth string) map[string]string {
	return map[string]string{
		"Accept":             "application/json, text/plain, */*",
		"Accept-Language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"Authorization":      basicAuth,
		"Connection":         "keep-alive",
		"Content-Length":     "0",
		"Origin":             "https://api.sitcline.com",
		"Referer":            "https://api.sitcline.com/sitcline/wel",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"Tenantid":           "2",
		"User-Agent":         h.userAgentGenerator.Generate(),
		"Istoken":            "false",
		"Sec-Ch-Ua":          "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "\"macOS\"",
	}
}

type urlGenerator struct {
}

func newUrlGenerator() *urlGenerator {
	return &urlGenerator{}
}

func (u *urlGenerator) generate(username, password, randomStr, solvedCaptcha string) string {
	return fmt.Sprintf(`https://api.sitcline.com/auth/oauth/token?username=%s&password=%s&randomStr=%s&code=%s&grant_type=password&scope=server`, username, password, randomStr, solvedCaptcha)
}

type Request struct {
	request          requests.IHttp
	headersGenerator *headersGenerator
	urlGenerator     *urlGenerator
}

func NewRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *Request {
	return &Request{request: request, headersGenerator: newHeadersGenerator(generator), urlGenerator: newUrlGenerator()}
}

func (r *Request) Login(ctx context.Context, basicAuth, username, password, randomStr, solvedCaptcha string) (*ApiResponse, error) {
	url := r.urlGenerator.generate(username, password, randomStr, solvedCaptcha)
	headers := r.headersGenerator.generate(basicAuth)
	response, err := r.request.Url(url).Method("POST").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 210 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s *ApiResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	return s, nil
}
