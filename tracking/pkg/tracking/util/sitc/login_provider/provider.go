package login_provider

import (
	"context"
	"golang_tracking/pkg/tracking/util/sitc/captcha_resolver"
)

type Provider struct {
	username        string
	password        string
	basicAuth       string
	request         *Request
	captchaResolver captcha_resolver.ICaptcha
}

func NewProvider(username string, password string, basicAuth string, request *Request, captchaResolver captcha_resolver.ICaptcha) *Provider {
	return &Provider{username: username, password: password, basicAuth: basicAuth, request: request, captchaResolver: captchaResolver}
}

func (p *Provider) Login(ctx context.Context) (*ApiResponse, error) {
	randomString, solvedCaptcha, err := p.captchaResolver.Resolve(ctx)
	if err != nil {
		return nil, err
	}
	return p.request.Login(ctx, p.basicAuth, p.username, p.password, string(randomString), string(solvedCaptcha))
}
