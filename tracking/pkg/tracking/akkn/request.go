package akkn

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang_tracking/pkg/tracking/util/requests"
)

type formGenerator struct {
}

func newFormGenerator() *formGenerator {
	return &formGenerator{}
}

func (f *formGenerator) Generate(number string) map[string]string {
	return map[string]string{
		"__EVENTTARGET":        "",
		"__EVENTARGUMENT":      "",
		"__VIEWSTATE":          "E1gQ1eFziTB7ZKSZ56VMxIrMhCREZZdQI6T3i0RUHv4+SJHFd0/BvZEZGfeJ7qc2Pd2Pu+5RqyRl2j17Z5HHiyOh2k6ke9ZUQVsdn6DOWLZqtqFtt0FjRBmPmxnaCvDlCBcfUZkiUH7CoBd3DSou6knZjsVCHBHFOPUud9GKSmpigtSniK3lbXznGA9HBgUy",
		"__VIEWSTATEGENERATOR": "90059987",
		"__PREVIOUSPAGE":       "Q6XXvMPmGNYA2w5c7Nq4T3u3LOxXQkp3lutE2VdecsGNEjGnVBSd6o4qFNseNufEQszZBM9KSdiGm9e3IMIZLw2",
		"__EVENTVALIDATION":    "qqOpIrpG+CiPE0zaMZ6hAC5MM9MpHGicz4+Gs1ZmwHEi5BFIOEvwrmDCjM5w9txD6aoNBJBR4HKMv13Yb5WmRpJ3lgk4exAY6AfLwFGry0ZcmQ4RNn8K+MJ3iwEHpHD97fmFlPQh86AHGvse9tbouXBleZt9WdW+W4TnznligrM=",
		"TextBox1":             number,
		"TextBox2":             "",
		"Button1":              "",
	}
}

type headersGenerator struct {
	uaGenerator requests.IUserAgentGenerator
}

func newHeadersGenerator(uaGenerator requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{uaGenerator: uaGenerator}
}
func (h *headersGenerator) Generate() map[string]string {
	return map[string]string{
		"authority":                 "sap.akkonlines.com",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-language":           "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"cache-control":             "max-age=0",
		"content-type":              "application/x-www-form-urlencoded",
		"origin":                    "https://sap.akkonlines.com",
		"referer":                   "https://sap.akkonlines.com/index.aspx",
		"sec-ch-ua":                 "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "\"macOS\"",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-user":            "?1",
		"upgrade-insecure-requests": "1",
		"user-agent":                h.uaGenerator.Generate(),
	}
}

type Request struct {
	request          requests.IHttp
	formGenerator    *formGenerator
	headersGenerator *headersGenerator
}

func NewRequest(request requests.IHttp, uaGenerator requests.IUserAgentGenerator) *Request {
	return &Request{request: request, formGenerator: newFormGenerator(), headersGenerator: newHeadersGenerator(uaGenerator)}
}

func (r *Request) Send(ctx context.Context, number string) (*goquery.Document, error) {
	form := r.formGenerator.Generate(number)
	headers := r.headersGenerator.Generate()

	response, err := r.request.Url("https://sap.akkonlines.com/index.aspx").Method("POST").Headers(headers).Form(form).Do(ctx)
	if err != nil {
		return nil, err
	}

	if response.Status > 200 {
		return nil, errors.New(fmt.Sprintf("status code invalid: %d", response.Status))
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
	if err != nil {
		return nil, err
	}
	return doc, nil
}
