package sitc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
	"golang_tracking/pkg/tracking/util/sitc/login_provider"
)

type containerRequestUrlGenerator struct {
}

func newContainerRequestUrlGenerator() *containerRequestUrlGenerator {
	return &containerRequestUrlGenerator{}
}

func (u *containerRequestUrlGenerator) generate(number string) string {
	return fmt.Sprintf("https://api.sitcline.com/ecm/cmcontainerhistory/movementSearch?blNo=&containerNo=%s&randomStr=", number)
}

type headersGenerator struct {
	userAgentGenerator requests.IUserAgentGenerator
	store              *login_provider.Store
}

func newHeadersGenerator(userAgentGenerator requests.IUserAgentGenerator, store *login_provider.Store) *headersGenerator {
	return &headersGenerator{userAgentGenerator: userAgentGenerator, store: store}
}

func (h *headersGenerator) generate() map[string]string {
	return map[string]string{
		"Accept":             "application/json, text/plain, */*",
		"Accept-Language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"Authorization":      fmt.Sprintf("Bearer %s", h.store.AuthToken()),
		"Connection":         "keep-alive",
		"Content-Length":     "0",
		"Origin":             "https://api.sitcline.com",
		"Referer":            "https://api.sitcline.com/sitcline/query/cargoTrack",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"Tenantid":           "2",
		"User-Agent":         h.userAgentGenerator.Generate(),
		"Sec-Ch-Ua":          "\"Google Chrome\";v=\"107\", \"Chromium\";v=\"107\", \"Not=A?Brand\";v=\"24\"",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "\"macOS\"",
	}
}

type ContainerTrackingRequest struct {
	request          requests.IHttp
	urlGenerator     *containerRequestUrlGenerator
	headersGenerator *headersGenerator
}

func NewContainerTrackingRequest(request requests.IHttp, generator requests.IUserAgentGenerator, store *login_provider.Store) *ContainerTrackingRequest {
	return &ContainerTrackingRequest{request: request, urlGenerator: newContainerRequestUrlGenerator(), headersGenerator: newHeadersGenerator(generator, store)}
}
func (c *ContainerTrackingRequest) Send(ctx context.Context, number string) (*ContainerApiResponse, error) {
	url := c.urlGenerator.generate(number)
	headers := c.headersGenerator.generate()
	response, err := c.request.Url(url).Method("POST").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 210 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s *ContainerApiResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	if len(s.Data.List) == 0 {
		return nil, errors.New("no len")
	}
	return s, nil
}

type billNumberUrlGenerator struct {
}

func newBillNumberUrlGenerator() *billNumberUrlGenerator {
	return &billNumberUrlGenerator{}
}

func (b *billNumberUrlGenerator) generateForGetBillInfo(number, randomString, solvedCaptcha string) string {
	return fmt.Sprintf("http://api.sitcline.com/doc/cargoTrack/searchApp?blNo=%s&code=%s&randomStr=%s", number, solvedCaptcha, randomString)
}

func (b *billNumberUrlGenerator) generateForGetContainerInfo(containerNo, billNo string) string {
	return fmt.Sprintf("http://api.sitcline.com/doc/cargoTrack/movementDetailApp?blNo=%s&containerNo=%s", billNo, containerNo)
}

type IBillRequest interface {
	GetBillNumberInfo(ctx context.Context, number, randomString, solvedCaptcha string) (*BillNumberApiResponse, error)
	GetContainerInfo(ctx context.Context, billNo, containerNo string) (*BillNumberInfoAboutContainerApiResponse, error)
}

type BillTrackingRequest struct {
	request          requests.IHttp
	urlGenerator     *billNumberUrlGenerator
	headersGenerator *headersGenerator
}

func NewBillTrackingRequest(request requests.IHttp, generator requests.IUserAgentGenerator, store *login_provider.Store) *BillTrackingRequest {
	return &BillTrackingRequest{
		request: request, urlGenerator: newBillNumberUrlGenerator(),
		headersGenerator: newHeadersGenerator(generator, store),
	}
}

func (b *BillTrackingRequest) GetBillNumberInfo(ctx context.Context, number, randomString, solvedCaptcha string) (*BillNumberApiResponse, error) {
	url := b.urlGenerator.generateForGetBillInfo(number, randomString, solvedCaptcha)
	headers := b.headersGenerator.generate()
	response, err := b.request.Url(url).Method("POST").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 250 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s *BillNumberApiResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func (b *BillTrackingRequest) GetContainerInfo(ctx context.Context, billNo, containerNo string) (*BillNumberInfoAboutContainerApiResponse, error) {
	url := b.urlGenerator.generateForGetContainerInfo(containerNo, billNo)
	headers := b.headersGenerator.generate()
	response, err := b.request.Url(url).Method("POST").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 250 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s *BillNumberInfoAboutContainerApiResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	return s, nil
}
