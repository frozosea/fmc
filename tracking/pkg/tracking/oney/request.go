package oney

import (
	"context"
	"encoding/json"
	"errors"
	"golang_tracking/pkg/tracking/util/requests"
	"strconv"
)

type headersGenerator struct {
	userAgentGenerator requests.IUserAgentGenerator
}

func newHeadersGenerator(userAgentGenerator requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{userAgentGenerator: userAgentGenerator}
}

func (h *headersGenerator) generate() map[string]string {
	return map[string]string{
		`Host`:               `ecomm.one-line.com`,
		`Accept`:             `application/json, text/javascript, */*; q=0.01`,
		`Accept-Language`:    `en-US,en;q=0.5`,
		`Accept-Encoding`:    `gzip, deflate, br`,
		`Connection`:         `keep-alive`,
		`content-type`:       `application/x-www-form-urlencoded`,
		`origin`:             `https://ecomm.one-line.com`,
		`referer`:            `https://ecomm.one-line.com/ecom/CUP_HOM_3301.do`,
		`sec-ch-ua`:          `" Not;A Brand";v = "99", "Google Chrome";v = "97", "Chromium";v = "97"`,
		`sec-ch-ua-mobile`:   `?0`,
		`sec-ch-ua-platform`: `macOS`,
		`sec-fetch-dest`:     `empty`,
		`sec-fetch-mode`:     `cors`,
		`sec-fetch-site`:     `same - origin`,
		`user-agent`:         h.userAgentGenerator.Generate(),
		`x-requested-with`:   `XMLHttpRequest`,
	}
}

type bodyGenerator struct {
}

func newBodyGenerator() *bodyGenerator {
	return &bodyGenerator{}
}
func (b *bodyGenerator) generateToGetBkgAndCopNumber(number string) map[string]string {
	return map[string]string{
		"f_cmd":       "122",
		"cust_id":     "",
		"cntr_no":     number,
		"search_type": "C",
	}
}
func (b *bodyGenerator) generateForInfoAboutMovingRequest(containerNumber, bookingNumber, copNo string) map[string]string {
	return map[string]string{
		"f_cmd":   "125",
		"cntr_no": containerNumber,
		"bkg_no":  bookingNumber,
		"cop_no":  copNo,
	}
}

func (b *bodyGenerator) generateForContainerSizeRequest(containerNumber, bookingNumber, copNo string) map[string]string {
	return map[string]string{
		"f_cmd":   "123",
		"cntr_no": containerNumber,
		"bkg_no":  bookingNumber,
		"cop_no":  copNo,
	}

}
func (b *bodyGenerator) generateForCheckContainerAccessoryToThisLine(number string) map[string]string {
	return map[string]string{
		"_search":     "false",
		"nd":          "1672021473638",
		"rows":        "10000",
		"page":        "1",
		"sidx":        "",
		"sord":        "acs",
		"f_cmd":       "121",
		"search_type": "A",
		"search_name": number,
		"cust_cd":     "",
	}
}

type Request struct {
	request          requests.IHttp
	headersGenerator *headersGenerator
	bodyGenerator    *bodyGenerator
}

func NewRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *Request {
	return &Request{request: request, headersGenerator: newHeadersGenerator(generator), bodyGenerator: newBodyGenerator()}
}

func (r *Request) sendRequestWithQuery(ctx context.Context, query map[string]string, method string) ([]byte, error) {
	const url = "https://ecomm.one-line.com/ecom/CUP_HOM_3301GS.do"
	response, err := r.request.Url(url).Method(method).Headers(r.headersGenerator.generate()).Query(query).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 300 {
		return response.Body, requests.NewStatusCodeError(response.Status)
	}
	return response.Body, nil
}
func (r *Request) SendRequestForGetBkgNoAndCopNo(ctx context.Context, number string) (*BkgAndCopNosApiResponseSchema, error) {
	query := r.bodyGenerator.generateToGetBkgAndCopNumber(number)
	var s *BkgAndCopNosApiResponseSchema
	b, err := r.sendRequestWithQuery(ctx, query, "POST")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}
func (r *Request) SendForInfoAboutMoving(ctx context.Context, number, bookingNumber, copNo string) (*InfoAboutMovingApiResponseSchema, error) {
	query := r.bodyGenerator.generateForInfoAboutMovingRequest(number, bookingNumber, copNo)
	var s *InfoAboutMovingApiResponseSchema
	b, err := r.sendRequestWithQuery(ctx, query, "POST")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}
func (r *Request) SendForContainerSize(ctx context.Context, number, bookingNumber, copNo string) (*ContainerSizeApiResponseSchema, error) {
	query := r.bodyGenerator.generateForContainerSizeRequest(number, bookingNumber, copNo)
	var s *ContainerSizeApiResponseSchema
	b, err := r.sendRequestWithQuery(ctx, query, "POST")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s, nil
}
func (r *Request) CheckContainerAccessoryToThisLine(ctx context.Context, number string) (bool, error) {
	query := r.bodyGenerator.generateForCheckContainerAccessoryToThisLine(number)
	var s *BaseApiResponseEntity
	b, err := r.sendRequestWithQuery(ctx, query, "POST")
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return false, err
	}
	countValue, err := strconv.Atoi(s.Count)
	if err != nil {
		return false, err
	} else if countValue == 0 {
		return false, nil
	} else if s.Exception != "" {
		return false, errors.New(s.Exception)
	}
	return true, nil
}
