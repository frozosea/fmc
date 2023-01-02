package captcha_resolver

import (
	"context"
	"fmt"
	"golang_tracking/pkg/tracking/util/requests"
)

type ICaptchaGenerator interface {
	GetCaptcha(ctx context.Context, randomString string) ([]byte, error)
}
type headersGeneratorForCaptchaGetter struct {
	userAgentGenerator requests.IUserAgentGenerator
}

func newHeadersGeneratorForCaptchaGetter(userAgentGenerator requests.IUserAgentGenerator) *headersGeneratorForCaptchaGetter {
	return &headersGeneratorForCaptchaGetter{userAgentGenerator: userAgentGenerator}
}

func (h *headersGeneratorForCaptchaGetter) generate() map[string]string {
	return map[string]string{
		"accept":          "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8",
		"accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
		"Referer":         "http://api.sitcline.com/app/cargoTrackSearch",
		"Referrer-Policy": "strict-origin-when-cross-origin",
		"User-Agent":      h.userAgentGenerator.Generate(),
	}
}

type urlGeneratorForCaptchaGetter struct {
}

func newUrlGeneratorForCaptchaGetter() *urlGeneratorForCaptchaGetter {
	return &urlGeneratorForCaptchaGetter{}
}

func (u *urlGeneratorForCaptchaGetter) generate(randomStr string) string {
	return fmt.Sprintf("http://api.sitcline.com/code?randomStr=%s", randomStr)
}

type CaptchaGetter struct {
	request          requests.IHttp
	headersGenerator *headersGeneratorForCaptchaGetter
	urlGenerator     *urlGeneratorForCaptchaGetter
}

func NewCaptchaGetter(request requests.IHttp, generator requests.IUserAgentGenerator) *CaptchaGetter {
	return &CaptchaGetter{
		request:          request,
		headersGenerator: newHeadersGeneratorForCaptchaGetter(generator),
		urlGenerator:     newUrlGeneratorForCaptchaGetter(),
	}
}

func (c *CaptchaGetter) GetCaptcha(ctx context.Context, randomString string) ([]byte, error) {
	url := c.urlGenerator.generate(randomString)
	headers := c.headersGenerator.generate()
	response, err := c.request.Url(url).Method("GET").Headers(headers).Do(ctx)
	if err != nil {
		return nil, err
	}
	if response.Status > 300 {
		return nil, requests.NewStatusCodeError(response.Status)
	}
	return response.Body, nil
}
