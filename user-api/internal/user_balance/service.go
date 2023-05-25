package user_balance

import (
	"context"
	"fmt"
	"user-api/internal/domain"
	"user-api/internal/user_balance/balance"
	"user-api/internal/user_balance/transactions"
	"user-api/pkg/logging"
)

type IService interface {
	GetCurrentTariff(ctx context.Context, userId int64) (*domain.Tariff, error)
	SubOneDayTrackingPriceFromBalance(ctx context.Context, userId int64, number string) error
	GetBalance(ctx context.Context, userId int64) (float64, error)
}

type Service struct {
	numbersRepository           IRepository
	oneDayPriceCalculator       IOneDayPriceCalculator
	balanceRepository           balance.IRepository
	numberTransactionRepository transactions.IRepository
	logger                      logging.ILogger
}

func NewService(numbersRepository IRepository, oneDayPriceCalculator IOneDayPriceCalculator, balanceRepository balance.IRepository, numberTransactionRepository transactions.IRepository, logger logging.ILogger) *Service {
	return &Service{numbersRepository: numbersRepository, oneDayPriceCalculator: oneDayPriceCalculator, balanceRepository: balanceRepository, numberTransactionRepository: numberTransactionRepository, logger: logger}
}

func (s *Service) GetCurrentTariff(ctx context.Context, userId int64) (*domain.Tariff, error) {
	numbersOnTrack, err := s.numbersRepository.GetAllNumbersOnTrack(ctx, userId)
	if err != nil {
		return nil, err
	}

	price, err := s.oneDayPriceCalculator.Calculate(ctx, numbersOnTrack)
	if err != nil {
		return nil, err
	}

	return &domain.Tariff{
		OneDayPrice:            price,
		NumbersOnTrackQuantity: numbersOnTrack,
	}, nil

}

func (s *Service) SubOneDayTrackingPriceFromBalance(ctx context.Context, userId int64, number string) error {
	tariff, err := s.GetCurrentTariff(ctx, userId)
	if err != nil {
		return err
	}

	subBalanceTransactionData, err := s.balanceRepository.Sub(ctx, userId, tariff.OneDayPrice)
	if err != nil {
		go s.logger.ExceptionLog(fmt.Sprint(`SubOneDayTrackingPriceFromBalance sub from balance repository error: %s`, err.Error()))
		return err
	}

	if err := s.numberTransactionRepository.Add(ctx, subBalanceTransactionData.ID, number); err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`SubOneDayTrackingPriceFromBalance sub from number transaction repository error: %s`, err.Error()))
		return err
	}

	return nil
}
func (s *Service) GetBalance(ctx context.Context, userId int64) (float64, error) {
	b, err := s.balanceRepository.Get(ctx, userId)
	if err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`get balance error: %s`, err.Error()))
		return 0, err
	}
	return b, nil
}
