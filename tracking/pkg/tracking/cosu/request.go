package cosu

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
	"time"
)

type IRequest interface {
	GetInfoAboutMovingRawResponse(ctx context.Context, number string) (*ApiResponseSchema, error)
	GetEtaRawResponse(ctx context.Context, number string) (*EtaApiResponseSchema, error)
}

type headersGenerator struct {
	userAgentGenerator requests.IUserAgentGenerator
}

func newHeadersGenerator(userAgentGenerator requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{userAgentGenerator: userAgentGenerator}
}

func (h *headersGenerator) generate() map[string]string {
	return map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"Connection":      "keep-alive",
		//"Cookie":             "HMF_CI=b31b2b5c10296bc7404f41ed06e6a85ad6a03767fd9219ce634ebda8b93eb030044b25e56fc282d7bd330a91e92e07f15c597c43e4b4a04c94e9201d8abc6b9e80; Hm_lvt_3dc23c86163f0d1cd70ef10ea94e0263=1672568806; COSCON_ACCESS_I18N=en-US; HBB_HC=252ecf4e91e8434adf73bc354872ecd63d5a8095e2de78335ab3dd6b17b10e91dff93e47fd5b733ac2794bb07b8b488d37; HMY_JC=4a571730370db1ecc4bf648f9248f1613c2503e3f05210e0c66414ee686bef2ce5,; Hm_lpvt_3dc23c86163f0d1cd70ef10ea94e0263=1672568819",
		"Referer":            "https://elines.coscoshipping.com/ebusiness/cargotracking",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"User-Agent":         h.userAgentGenerator.Generate(),
		"X-Client-Timestamp": fmt.Sprintf(`%d`, time.Now().UTC().Unix()),
		"language":           "en_US",
		"sec-ch-ua":          "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"macOS\"",
		"sys":                "eb",
	}
}

type urlGenerator struct {
}

func newUrlGenerator() *urlGenerator {
	return &urlGenerator{}
}

func (u *urlGenerator) generateGetInfoAboutMovingUrl(number string) string {
	return fmt.Sprintf(`https://elines.coscoshipping.com/ebtracking/public/containers/%s?timestamp=%d`, number, time.Now().UTC().Unix())
}
func (u *urlGenerator) generateGetEtaUrl(number string) string {
	return fmt.Sprintf(`https://elines.coscoshipping.com/ebtracking/public/container/eta/%s?timestamp=%d`, number, time.Now().UTC().Unix())
}

type Request struct {
	requests   requests.IHttp
	urlGen     *urlGenerator
	headersGen *headersGenerator
}

func NewRequest(requests requests.IHttp, generator requests.IUserAgentGenerator) *Request {
	return &Request{requests: requests, urlGen: newUrlGenerator(), headersGen: newHeadersGenerator(generator)}
}

func (r *Request) GetInfoAboutMovingRawResponse(ctx context.Context, number string) (*ApiResponseSchema, error) {
	url := r.urlGen.generateGetInfoAboutMovingUrl(number)
	headers := r.headersGen.generate()
	response, err := r.requests.Url(url).Method("GET").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 300 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s *ApiResponseSchema
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	if len(s.Data.Content.Containers) == 0 {
		return nil, errors.New("no len")
	}
	return s, nil
}

func (r *Request) GetEtaRawResponse(ctx context.Context, number string) (*EtaApiResponseSchema, error) {
	url := r.urlGen.generateGetEtaUrl(number)
	headers := r.headersGen.generate()
	httpResponse, err := r.requests.Url(url).Method("GET").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if httpResponse.Status > 300 {
		return nil, requests.NewStatusCodeError(httpResponse.Status)
	}
	var eta *EtaApiResponseSchema
	if err := json.Unmarshal(httpResponse.Body, &eta); err != nil {
		return nil, err
	}
	return eta, nil
}
