package feedback

import (
	"context"
	"fmt"
	"user-api/pkg/logging"
	"user-api/pkg/mailing"
)

type messageGenerator struct {
}

func newMessageGenerator() *messageGenerator {
	return &messageGenerator{}
}

func (m *messageGenerator) Gen(fb *Feedback) string {
	//TODO message generator
	return fmt.Sprintf("email: %s\n %s", fb.Email, fb.Message)
}
func (m *messageGenerator) getTextType() string {
	return "text/plain"
}

type Service struct {
	mailing      mailing.IMailing
	repository   IRepository
	logger       logging.ILogger
	msgGen       *messageGenerator
	sendToEmails []string
}

func NewService(mailing mailing.IMailing, repository IRepository, logger logging.ILogger, sendToEmails []string) *Service {
	return &Service{mailing: mailing, repository: repository, logger: logger, msgGen: newMessageGenerator(), sendToEmails: sendToEmails}
}

func (s *Service) Add(ctx context.Context, fb *Feedback) error {
	errCh := make(chan error, 1)
	go func() {
		if err := s.repository.Save(ctx, fb); err != nil {
			errCh <- err
		}
	}()
	go func() {
		subject := fmt.Sprintf("container tracking feedback by %s", fb.Email)
		if err := s.mailing.SendSimple(s.sendToEmails, subject, s.msgGen.Gen(fb), s.msgGen.getTextType()); err != nil {
			errCh <- err
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return err
		default:
			return nil
		}
	}
}
func (s *Service) GetByEmail(ctx context.Context, email string) ([]*Feedback, error) {
	return s.repository.GetByEmail(ctx, email)
}

func (s *Service) GetAll(ctx context.Context) ([]*Feedback, error) {
	return s.repository.GetAll(ctx)
}
