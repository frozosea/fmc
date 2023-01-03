package captcha_resolver

import (
	"context"
)

type ICaptcha interface {
	Resolve(ctx context.Context) (RandomString, SolvedCaptcha, error)
}

type Captcha struct {
	randomStringGenerator IRandomStringGenerator
	captchaGenerator      ICaptchaGenerator
	solver                ICaptchaSolver
}

func NewCaptcha(randomStringGenerator IRandomStringGenerator, captchaGenerator ICaptchaGenerator, solver ICaptchaSolver) *Captcha {
	return &Captcha{randomStringGenerator: randomStringGenerator, captchaGenerator: captchaGenerator, solver: solver}
}

func (c *Captcha) Resolve(ctx context.Context) (RandomString, SolvedCaptcha, error) {
	randomStr, err := c.randomStringGenerator.Generate()
	if err != nil {
		return "", "", err
	}
	captchaImage, err := c.captchaGenerator.GetCaptcha(ctx, randomStr)
	if err != nil {
		return "", "", err
	}
	solvedCaptcha, err := c.solver.Solve(ctx, captchaImage)
	if err != nil {
		return "", "", err
	}
	return RandomString(randomStr), SolvedCaptcha(solvedCaptcha), nil
}
