package feedback

import (
	"context"
	"fmt"
	"github.com/frozosea/mailing"
	"golang.org/x/sync/errgroup"
	"user-api/pkg/logging"
)

type messageGenerator struct {
}

func newMessageGenerator() *messageGenerator {
	return &messageGenerator{}
}

func (m *messageGenerator) Gen(fb *Feedback) string {
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
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := s.repository.Save(ctx, fb); err != nil {
			s.logger.ExceptionLog(fmt.Sprintf(`save feedback to repository error: %s`, err.Error()))
			return err
		}
		return nil
	})
	g.Go(func() error {
		subject := fmt.Sprintf("container tracking feedback by %s", fb.Email)
		if err := s.mailing.SendSimple(ctx, s.sendToEmails, subject, s.msgGen.Gen(fb), s.msgGen.getTextType()); err != nil {
			s.logger.ExceptionLog(fmt.Sprintf(`send email with feedback error: %s`, err.Error()))
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
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
