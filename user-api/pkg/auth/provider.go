package auth

import (
	"context"
	"fmt"
	"user-api/internal/logging"
	"user-api/pkg/domain"
)

type Provider struct {
	repository   IRepository
	tokenManager ITokenManager
	logger       logging.ILogger
}

func NewProvider(repository IRepository, tokenManager ITokenManager, logger logging.ILogger) *Provider {
	return &Provider{repository: repository, tokenManager: tokenManager, logger: logger}
}

func (p *Provider) RegisterUser(ctx context.Context, user domain.User) error {
	if regUserErr := p.repository.Register(ctx, user); regUserErr != nil {
		go p.logger.ExceptionLog(fmt.Sprintf(`user with username %s failed to register %s`, user.Username, regUserErr.Error()))
		return regUserErr
	}
	go p.logger.InfoLog(fmt.Sprintf(`user with username %s was registered`, user.Username))
	return nil
}
func (p *Provider) Login(ctx context.Context, user domain.User) (*Token, error) {
	userId, err := p.repository.Login(ctx, user)
	p.logger.InfoLog(fmt.Sprintf(`user with id %d was login`, userId))
	switch err.(type) {
	case *InvalidUserError:
		return nil, err
	}
	token, genTokenErr := p.tokenManager.GenerateAccessRefreshTokens(userId)
	if genTokenErr != nil {
		p.logger.ExceptionLog(fmt.Sprintf(`generate access refresh tokens for user-pb: %d error: %s`, userId, genTokenErr.Error()))
		return nil, genTokenErr
	}
	return token, genTokenErr
}
func (p *Provider) RefreshToken(refreshToken string) (*Token, error) {
	var token *Token
	userId, decodeTokenErr := p.tokenManager.DecodeToken(refreshToken)
	if decodeTokenErr != nil {
		return token, decodeTokenErr
	}
	return p.tokenManager.GenerateAccessRefreshTokens(userId)
}
func (p *Provider) CheckAccess(ctx context.Context, tokenString string) (bool, error) {
	userId, decodeTokenErr := p.tokenManager.DecodeToken(tokenString)
	if decodeTokenErr != nil || userId < 0 {
		return false, decodeTokenErr
	}
	return p.repository.CheckAccess(ctx, userId)
}
func (p *Provider) GetUserIdByJwtToken(ctx context.Context, tokenString string) (int, error) {
	res, err := p.tokenManager.DecodeToken(tokenString)
	if err != nil {
		return -1, err
	}
	select {
	case <-ctx.Done():
		return -1, ctx.Err()
	default:
	}
	return res, nil
}
