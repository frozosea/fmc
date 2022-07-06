package mailing

import (
	"fmt"
	"net/http"
	file_reader "user-api/internal/file-reader"
	"user-api/internal/logging"
)

//var (
//	SENDER_NAME  = os.Getenv("SENDER_NAME")
//	SENDER_EMAIL = os.Getenv("SENDER_EMAIL")
//)

type IMailing interface {
	SendWithFile(toAddress, subject, filePath string) error
}

//Mailing TODO test
type Mailing struct {
	reader          *file_reader.FileReader
	logger          logging.ILogger
	senderName      string
	senderEmail     string
	UnisenderApiKey string
	htmlTemplate    IHtmlTemplate
}

func NewMailing(logger logging.ILogger, senderName string, senderEmail string, unisenderApiKey string, htmlTemplate IHtmlTemplate) *Mailing {
	return &Mailing{reader: file_reader.NewFileReader(), logger: logger, senderName: senderName, senderEmail: senderEmail, UnisenderApiKey: unisenderApiKey, htmlTemplate: htmlTemplate}
}

func (m *Mailing) getUrl(toAddress, subject, fileName, textBody, htmlBody string, file []byte) string {
	return fmt.Sprintf("https://api.unisender.com/ru/api/createEmailMessage?format=json&api_key=%s&sender_name=%s&sender_email=%ssubject=%s&body=%s&attachements[%s]=%b&lang=RU&wrap_type=STRING&text_body=%s", m.UnisenderApiKey, m.senderName, m.senderEmail, subject, htmlBody, fileName, file, textBody)
}
func (m *Mailing) sendEmail(url string) error {
	client := http.Client{}
	r, err := client.Post(url, "application/json", nil)
	defer r.Body.Close()
	return err
}
func (m *Mailing) SendWithFile(toAddress, subject, filePath string) error {
	fileName, err := m.reader.GetFileName(filePath)
	if err != nil {
		return err
	}
	readFile, err := m.reader.ReadFile(filePath)
	if err != nil {
		return err
	}
	url := m.getUrl(toAddress, subject, fileName, "", m.htmlTemplate.GetTrackingTemplate(fileName), readFile)
	if err := m.sendEmail(url); err != nil {
		go m.logger.ExceptionLog(fmt.Sprintf(`send email: %s failed: %s`, toAddress, err.Error()))
		return err
	}
	return nil
}
