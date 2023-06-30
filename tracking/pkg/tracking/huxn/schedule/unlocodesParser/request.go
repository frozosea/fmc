package unlocodesParser

import (
	"bytes"
	"context"
	"github.com/PuerkitoBio/goquery"
	"golang_tracking/pkg/tracking/util/requests"
)

type Request struct {
	request requests.IHttp
}

func NewRequest(request requests.IHttp) *Request {
	return &Request{request: request}
}

func (r *Request) Send(ctx context.Context) (*goquery.Document, error) {
	const url = "http://dc.hxlines.com:8099/HX_WeChat/HX_Schedules"

	response, err := r.request.Url(url).Method("GET").Do(ctx)
	if err != nil {
		return nil, err
	}

	if response.Status > 200 {
		return nil, requests.NewStatusCodeError(response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
	if err != nil {
		return nil, err
	}

	return doc, nil
}
