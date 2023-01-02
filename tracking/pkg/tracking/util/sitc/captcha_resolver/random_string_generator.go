package captcha_resolver

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

type IRandomStringGenerator interface {
	Generate() (string, error)
}

type RandomStringGenerator struct {
}

func NewRandomStringGenerator() *RandomStringGenerator {
	return &RandomStringGenerator{}
}

func (r *RandomStringGenerator) Generate() (string, error) {
	const desiredMaxLength = 17
	maxLimit := int64(int(math.Pow10(desiredMaxLength)) - 1)
	lowLimit := int(math.Pow10(desiredMaxLength - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return "", err
	}
	randomNumberInt := int(randomNumber.Int64())

	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return fmt.Sprintf(`%d`, randomNumberInt), nil
}
