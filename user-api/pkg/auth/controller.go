package auth

import (
	"context"
	"fmt"
	"user-api/internal/logging"
	"user-api/pkg/domain"
)

type Controller struct {
	repository   IRepository
	tokenManager ITokenManager
	logger       logging.ILogger
}

func NewController(repository IRepository, tokenManager ITokenManager, logger logging.ILogger) *Controller {
	return &Controller{repository: repository, tokenManager: tokenManager, logger: logger}
}

func (c *Controller) RegisterUser(ctx context.Context, user domain.User) error {
	if regUserErr := c.repository.Register(ctx, user); regUserErr != nil {
		//go c.logger.ExceptionLog(fmt.Sprintf(`user-pb with username %s`, user.Username))
		return regUserErr
	}
	//go c.logger.InfoLog(fmt.Sprintf(`user-pb with username %s was registered`, user.Username))
	return nil
}
func (c *Controller) Login(ctx context.Context, user domain.User) (*Token, error) {
	var token *Token
	userId, err := c.repository.Login(ctx, user)
	if err != nil {
		return token, err
	}
	token, genTokenErr := c.tokenManager.GenerateAccessRefreshTokens(userId)
	if genTokenErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`generate access refresh tokens for user-pb: %d error: %s`, userId, genTokenErr.Error()))
	}
	return token, genTokenErr
}
func (c *Controller) RefreshToken(refreshToken string) (*Token, error) {
	var token *Token
	userId, decodeTokenErr := c.tokenManager.DecodeToken(refreshToken)
	if decodeTokenErr != nil {
		return token, decodeTokenErr
	}
	return c.tokenManager.GenerateAccessRefreshTokens(userId)
}
func (c *Controller) CheckAccess(ctx context.Context, tokenString string) (bool, error) {
	userId, decodeTokenErr := c.tokenManager.DecodeToken(tokenString)
	if decodeTokenErr != nil {
		return false, decodeTokenErr
	}
	return c.repository.CheckAccess(ctx, userId)
}
