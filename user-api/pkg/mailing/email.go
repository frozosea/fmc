package mailing

import (
	"context"
	"gopkg.in/gomail.v2"
	"net/smtp"
	"sync"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	// return "LOGIN", []byte{}, nil
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

type IMailing interface {
	SendSimple(ctx context.Context, toAddresses []string, subject, body, textType string) error
}

type Mailing struct {
	smtpHost  string
	smtpPort  int
	fromEmail string
	password  string
	authKey   string
}

func NewMailing(smtpHost string, smtpPort int, fromEmail string, password string, authKey string) *Mailing {
	return &Mailing{smtpHost: smtpHost, smtpPort: smtpPort, fromEmail: fromEmail, password: password, authKey: authKey}
}
func (w *Mailing) SendSimple(ctx context.Context, toAddress []string, subject, message, textType string) error {
	errCh := make(chan error, 1)
	var wg sync.WaitGroup
	for _, toEmail := range toAddress {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m := gomail.NewMessage()
			m.SetHeader("From", w.fromEmail)
			m.SetHeader("To", toEmail)
			m.SetHeader("Subject", subject)
			m.SetBody(textType, message)
			d := gomail.NewDialer(w.smtpHost, w.smtpPort, w.fromEmail, w.password)
			if err := d.DialAndSend(m); err != nil {
				errCh <- err
			}
		}()
		wg.Wait()
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errCh:
			return err
		default:
			return nil
		}
	}
}
