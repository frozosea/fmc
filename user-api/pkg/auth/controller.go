package auth

import (
	"context"
	"fmt"
	"user-api/internal/logging"
	"user-api/pkg/user"
)

type Controller struct {
	repository      IRepository
	hash            IHash
	tokenManager    ITokenManager
	tokenExpiration int
	logger          logging.ILogger
}

func (c *Controller) RegisterUser(ctx context.Context, user user.User) error {
	if regUserErr := c.repository.Register(ctx, user); regUserErr != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`user with username %s`, user.Username))
		return regUserErr
	}
	go c.logger.InfoLog(fmt.Sprintf(`user with username %s was registered`, user.Username))
	return nil
}
func (c *Controller) Login(ctx context.Context, user user.User) (*Token, error) {
	var token *Token
	userId, err := c.repository.Login(ctx, user)
	if err != nil {
		return token, err
	}
	token, genTokenErr := c.tokenManager.GenerateAccessRefreshTokens(userId)
	if genTokenErr != nil {
		c.logger.ExceptionLog(fmt.Sprintf(`generate access refresh tokens for user: %d error: %s`, userId, genTokenErr.Error()))
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
