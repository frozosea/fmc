package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/frozosea/mailing"
	"time"
	"user-api/internal/domain"
	"user-api/pkg/logging"
)

type Service struct {
	repository   IRepository
	tokenManager ITokenManager
	hash         IHash
	*RecoveryUserTemplateGenerator
	emailSender mailing.IMailing
	logger      logging.ILogger
}

func NewService(repository IRepository, tokenManager ITokenManager, hash IHash, emailSender mailing.IMailing, logger logging.ILogger) *Service {
	return &Service{repository: repository, tokenManager: tokenManager, logger: logger, hash: hash, emailSender: emailSender, RecoveryUserTemplateGenerator: NewRecoveryUserTemplateGenerator()}
}

func (p *Service) RegisterUser(ctx context.Context, user *domain.RegisterUser) error {
	hashPassword, hashErr := p.hash.Hash(user.Password)
	if hashErr != nil {
		return hashErr
	}
	user.Password = hashPassword
	if regUserErr := p.repository.Register(ctx, user); regUserErr != nil {
		go p.logger.ExceptionLog(fmt.Sprintf(`user with username %s failed to register %s`, user.Email, regUserErr.Error()))
		return regUserErr
	}
	go p.logger.InfoLog(fmt.Sprintf(`user with username %s was registered`, user.Email))
	return nil
}
func (p *Service) Login(ctx context.Context, user *domain.User) (*Token, error) {
	userId, err := p.repository.Login(ctx, user)
	if err != nil {
		return nil, err
	}
	go p.logger.InfoLog(fmt.Sprintf(`user with id %d was login`, userId))
	token, genTokenErr := p.tokenManager.GenerateAccessRefreshTokens(userId)
	if genTokenErr != nil {
		go p.logger.ExceptionLog(fmt.Sprintf(`generate access refresh tokens for user-pb: %d error: %s`, userId, genTokenErr.Error()))
		return nil, genTokenErr
	}
	return token, genTokenErr
}
func (p *Service) RefreshToken(refreshToken string) (*Token, error) {
	var token *Token
	userId, decodeTokenErr := p.tokenManager.DecodeToken(refreshToken)
	if decodeTokenErr != nil {
		return token, decodeTokenErr
	}
	return p.tokenManager.GenerateAccessRefreshTokens(userId)
}
func (p *Service) CheckAccess(ctx context.Context, tokenString string) (bool, error) {
	userId, decodeTokenErr := p.tokenManager.DecodeToken(tokenString)
	if decodeTokenErr != nil || userId < 0 {
		return false, decodeTokenErr
	}
	return p.repository.CheckAccess(ctx, userId)
}
func (p *Service) GetUserIdByJwtToken(tokenString string) (int, error) {
	res, err := p.tokenManager.DecodeToken(tokenString)
	if err != nil {
		return -1, err
	}
	return res, nil
}
func (p *Service) SendRecoveryUserEmail(ctx context.Context, email string) error {
	if exist, err := p.repository.CheckEmailExist(ctx, email); !exist || err != nil {
		return &InvalidUserError{}
	}
	userId, err := p.repository.GetUserId(ctx, email)
	if err != nil {
		return err
	}
	token, err := p.tokenManager.GenerateResetPasswordToken(userId, time.Hour)
	if err != nil {
		return err
	}
	template, err := p.GetRecoveryUserTemplate(token)
	if err != nil {
		return err
	}
	return p.emailSender.SendSimple(ctx, []string{email}, "FindMyCargo recovery password", template, "text/html")
}

func (p *Service) RecoveryUser(ctx context.Context, token string, newPassword string) error {
	userId, operationType, err := p.tokenManager.DecodeResetPasswordToken(token)
	if err != nil {
		return err
	}
	if operationType != "reset_password" {
		return errors.New("invalid token")
	}
	hashPwd, err := p.hash.Hash(newPassword)
	if err != nil {
		return err
	}
	return p.repository.SetNewPassword(ctx, userId, hashPwd)
}
