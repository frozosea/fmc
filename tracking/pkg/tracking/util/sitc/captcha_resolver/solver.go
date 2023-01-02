package captcha_resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
	"time"
)

type ICaptchaSolver interface {
	Solve(ctx context.Context, image []byte) (string, error)
}

type captchaSolverFormGenerator struct {
}

func newCaptchaSolverFormGenerator() *captchaSolverFormGenerator {
	return &captchaSolverFormGenerator{}
}

func (c *captchaSolverFormGenerator) generate(authKey string) map[string]string {
	return map[string]string{
		"key":      authKey,
		"numeric":  "1",
		"json":     "1",
		"phrase":   "0",
		"regsense": "0",
		"calc":     "0",
		"min_len":  "4",
		"max_len":  "4",
		"language": "0",
	}
}

type CaptchaSolver struct {
	request       requests.IHttp
	authKey       string
	formGenerator *captchaSolverFormGenerator
}

func NewCaptchaSolver(request requests.IHttp, authKey string) *CaptchaSolver {
	return &CaptchaSolver{request: request, authKey: authKey, formGenerator: newCaptchaSolverFormGenerator()}
}

func (c *CaptchaSolver) sendRequestToSolve(ctx context.Context, image []byte) (string, error) {
	form := c.formGenerator.generate(c.authKey)
	const url = "http://2captcha.com/in.php"
	response, err := c.request.Url(url).Method("POST").MultipartForm(form, "file", "", image).Do(ctx)
	if err != nil {
		return "", err
	}
	if response.Status > 300 {
		return "", requests.NewStatusCodeError(response.Status)
	}
	var s *CaptchaSolverGetIdResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return "", err
	}
	return s.Request, nil
}

func (c *CaptchaSolver) getSolvedCaptcha(ctx context.Context, id string) (*CaptchaSolverResponse, error) {
	url := fmt.Sprintf("`http://2captcha.com/res.php?key=%s&action=get&id=%s&json=1", c.authKey, id)
	response, err := c.request.Url(url).Method("GET").Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 300 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	var s *CaptchaSolverResponse
	if err := json.Unmarshal(response.Body, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func (c *CaptchaSolver) Solve(ctx context.Context, image []byte) (string, error) {
	id, err := c.sendRequestToSolve(ctx, image)
	if err != nil {
		return "", err
	}
	for {
		time.Sleep(time.Millisecond * 1000)
		response, err := c.getSolvedCaptcha(ctx, id)
		if err != nil {
			return "", err
		}
		if response.Request != "CAPCHA_NOT_READY" {
			return response.Request, nil
		}
	}
}
