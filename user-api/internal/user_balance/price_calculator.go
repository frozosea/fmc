package user_balance

import "context"

// IOneDayPriceCalculator TODO realize interface by formula
type IOneDayPriceCalculator interface {
	Calculate(ctx context.Context, numbersOnTrackQuantity int64) (float64, error)
}

type PriceCalculator struct{}

func NewPriceCalculator() *PriceCalculator {
	return &PriceCalculator{}
}

func (PriceCalculator) Calculate(_ context.Context, _ int64) (float64, error) {
	return 0, nil
}
