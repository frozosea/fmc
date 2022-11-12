package mailing

import (
	"gopkg.in/gomail.v2"
)

type IMailing interface {
	//SendWithFile(toAddress, subject, filePath string) error
	SendSimple(toAddresses []string, subject, body, textType string) error
}

type Mailing struct {
	smtpHost  string
	smtpPort  int
	fromEmail string
	password  string
}

func NewMailing(smtpHost string, smtpPort int, fromEmail string, password string) *Mailing {
	return &Mailing{smtpHost: smtpHost, smtpPort: smtpPort, fromEmail: fromEmail, password: password}
}

func (w *Mailing) SendSimple(toAddress []string, subject, message, textType string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", w.fromEmail)
	m.SetHeader("Cc", toAddress...)
	m.SetHeader("Subject", subject)
	m.SetBody(textType, message)
	d := gomail.NewDialer(w.smtpHost, w.smtpPort, w.fromEmail, w.password)
	return d.DialAndSend(m)
}
