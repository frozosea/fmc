package credit

import (
	"context"
	"user-api/internal/user_balance/balance"
)

type IService interface {
	CheckAccessToPaidTracking(ctx context.Context, userId int64) (bool, error)
}

type Service struct {
	repository             balance.IRepository
	minimalPossibleBalance float64
}

func NewService(repository balance.IRepository, minimalPossibleBalance float64) *Service {
	return &Service{repository: repository, minimalPossibleBalance: minimalPossibleBalance}
}

func (s *Service) CheckAccessToPaidTracking(ctx context.Context, userId int64) (bool, error) {
	userBalance, err := s.repository.Get(ctx, userId)
	if err != nil {
		return false, err
	}
	return userBalance > s.minimalPossibleBalance, nil
}
