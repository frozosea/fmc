package dnyg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
)

type INumberInfoRequest interface {
	Send(ctx context.Context, number string, isContainer bool) (*NumberInfoResponse, error)
}

type numberInfoRequestBodyGenerator struct {
}

func newNumberInfoRequestBodyGenerator() *numberInfoRequestBodyGenerator {
	return &numberInfoRequestBodyGenerator{}
}

func (n *numberInfoRequestBodyGenerator) Generate(number string, isContainer bool) ([]byte, error) {
	var t struct {
		DmaSearchInfo struct {
			BUKRS      string `json:"BUKRS"`
			INPBNO     string `json:"INPBNO"`
			INPCNTRNO  string `json:"INPCNTRNO"`
			INPBKN     string `json:"INPBKN"`
			USRCCD     string `json:"USRCCD"`
			PROFILESEQ string `json:"PROFILESEQ"`
			LANGCD     string `json:"LANGCD"`
		} `json:"dma_searchInfo"`
	}

	t.DmaSearchInfo.BUKRS = "2000"

	if isContainer {
		t.DmaSearchInfo.INPCNTRNO = number
	} else {
		t.DmaSearchInfo.INPBNO = number
	}

	t.DmaSearchInfo.LANGCD = "en"

	return json.Marshal(t)
}

type headersGenerator struct {
	userAgentGenerator requests.IUserAgentGenerator
}

func newHeadersGenerator(userAgentGenerator requests.IUserAgentGenerator) *headersGenerator {
	return &headersGenerator{userAgentGenerator: userAgentGenerator}
}
func (h *headersGenerator) Generate() map[string]string {
	return map[string]string{
		"Accept":             "application/json",
		"Accept-Language":    "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"Connection":         "keep-alive",
		"Content-Type":       "application/json; charset=\"UTF-8\"",
		"Origin":             "https://ebiz.pcsline.co.kr",
		"Referer":            "https://ebiz.pcsline.co.kr/",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-origin",
		"User-Agent":         h.userAgentGenerator.Generate(),
		"sec-ch-ua":          "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"",
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": "\"macOS\"",
	}
}

type NumberInfoRequest struct {
	request                        requests.IHttp
	numberInfoRequestBodyGenerator *numberInfoRequestBodyGenerator
	headersGenerator               *headersGenerator
}

func NewNumberInfoRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *NumberInfoRequest {
	return &NumberInfoRequest{
		request:                        request,
		numberInfoRequestBodyGenerator: newNumberInfoRequestBodyGenerator(),
		headersGenerator:               newHeadersGenerator(generator),
	}
}

func (n *NumberInfoRequest) Send(ctx context.Context, number string, isContainer bool) (*NumberInfoResponse, error) {
	body, err := n.numberInfoRequestBodyGenerator.Generate(number, isContainer)
	if err != nil {
		return nil, err
	}

	headers := n.headersGenerator.Generate()

	resp, err := n.request.Url("https://ebiz.pcsline.co.kr/trk/trkE0710R01N").Method("POST").Headers(headers).Body(body).Do(ctx)
	if err != nil {
		return nil, err
	}

	if resp.Status > 200 {
		return nil, errors.New(fmt.Sprintf("status code is invalid, current status: %s", err.Error()))
	}

	var numberInfoResponse *NumberInfoResponse

	if err := json.Unmarshal(resp.Body, &numberInfoResponse); err != nil {
		return nil, err
	}

	return numberInfoResponse, nil

}

type InfoAboutMovingRequestBodyGenerator struct {
}

func NewInfoAboutMovingRequestBodyGenerator() *InfoAboutMovingRequestBodyGenerator {
	return &InfoAboutMovingRequestBodyGenerator{}
}

func (i *InfoAboutMovingRequestBodyGenerator) Generate(bkgNo, billno, cntrNo string) ([]byte, error) {
	var t struct {
		DmaParamInfo struct {
			BUKRS     string `json:"BUKRS"`
			INPBNO    string `json:"INPBNO"`
			INPBKN    string `json:"INPBKN"`
			INPCNTRNO string `json:"INPCNTRNO"`
			USRCCD    string `json:"USRCCD"`
			USRNAT    string `json:"USRNAT"`
			LANGCD    string `json:"LANGCD"`
		} `json:"dma_paramInfo"`
		DmaSearchInfo struct {
			BUKRS      string `json:"BUKRS"`
			INPBNO     string `json:"INPBNO"`
			INPCNTRNO  string `json:"INPCNTRNO"`
			INPBKN     string `json:"INPBKN"`
			USRCCD     string `json:"USRCCD"`
			PROFILESEQ string `json:"PROFILESEQ"`
			LANGCD     string `json:"LANGCD"`
		} `json:"dma_searchInfo"`
	}
	t.DmaParamInfo.BUKRS = "2000"
	t.DmaParamInfo.LANGCD = "en"

	t.DmaParamInfo.INPBKN = bkgNo
	t.DmaParamInfo.INPBNO = billno
	t.DmaParamInfo.INPCNTRNO = cntrNo
	t.DmaParamInfo.INPCNTRNO = cntrNo

	t.DmaSearchInfo.BUKRS = "2000"
	t.DmaSearchInfo.LANGCD = "en"
	t.DmaSearchInfo.INPCNTRNO = cntrNo

	return json.Marshal(t)

}

type IInfoAboutMovingRequest interface {
	Send(ctx context.Context, bkgNo, billno, cntrNo string) (*InfoAboutMovingResponse, error)
}

type InfoAboutMovingRequest struct {
	request          requests.IHttp
	bodyGenerator    *InfoAboutMovingRequestBodyGenerator
	headersGenerator *headersGenerator
}

func NewInfoAboutMovingRequest(request requests.IHttp, generator requests.IUserAgentGenerator) *InfoAboutMovingRequest {
	return &InfoAboutMovingRequest{
		request:          request,
		bodyGenerator:    NewInfoAboutMovingRequestBodyGenerator(),
		headersGenerator: newHeadersGenerator(generator),
	}
}

func (i *InfoAboutMovingRequest) Send(ctx context.Context, bkgNo, billno, cntrNo string) (*InfoAboutMovingResponse, error) {
	body, err := i.bodyGenerator.Generate(bkgNo, billno, cntrNo)
	if err != nil {
		return nil, err
	}

	headers := i.headersGenerator.Generate()

	resp, err := i.request.Url("https://ebiz.pcsline.co.kr/trk/trkE0710R03").Method("POST").Headers(headers).Body(body).Do(ctx)
	if err != nil {
		return nil, err
	}

	if resp.Status > 200 {
		return nil, errors.New(fmt.Sprintf("status code is invalid, current status: %s", err.Error()))
	}

	var infoAboutMovingResponse *InfoAboutMovingResponse

	if err := json.Unmarshal(resp.Body, &infoAboutMovingResponse); err != nil {
		return nil, err
	}

	return infoAboutMovingResponse, nil
}
