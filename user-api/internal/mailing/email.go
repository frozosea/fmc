package mailing

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
type Response struct {
	Result struct {
		Statuses []struct {
			Id     int64  `json:"id"`
			Status string `json:"status"`
		} `json:"statuses"`
	} `json:"result"`
}

//Mailing can send email
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

func (m *Mailing) getForm(toAddress, subject, fileName, htmlBody string, file string) url.Values {
	query := url.Values{}
	query.Set("format", "json")
	query.Set("api_key", m.UnisenderApiKey)
	query.Set("sender_name", m.senderName)
	query.Set("email", toAddress)
	query.Set("sender_email", m.senderEmail)
	query.Set("subject", subject)
	query.Set("body", htmlBody)
	query.Set("wrap_type", "STRING")
	query.Set("list_id", "1")
	query.Set(fmt.Sprintf(`attachments[%s]`, fileName), file)
	return query
}
func (m *Mailing) checkStatusOfEmail(id string) error {
	client := http.Client{}
	checkStatusUrl := fmt.Sprintf(`https://api.unisender.com/ru/api/checkEmail?format=json&api_key=%s&email_id=%s`, m.UnisenderApiKey, id)
	r, err := client.Get(checkStatusUrl)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		return readErr
	}
	var s Response
	if unmarshalErr := json.Unmarshal(body, &s); unmarshalErr != nil {
		return unmarshalErr
	}
	for _, v := range s.Result.Statuses {
		if v.Status != "ok_sent" {
			return errors.New("email was not sent successfully")
		}
	}
	return nil
}
func (m *Mailing) sendEmail(form url.Values) (string, error) {
	client := http.Client{}
	r, err := client.PostForm("https://api.unisender.com/ru/api/sendEmail", form)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	if r.StatusCode > 250 {
		return "", errors.New("bad status code")
	}
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		return "", readErr
	}
	go m.logger.InfoLog(fmt.Sprintf(`send email result: %s`, string(body)))

	return "", nil
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
	form := m.getForm(toAddress, subject, fileName, m.htmlTemplate.GetTrackingTemplate(fileName), string(readFile))
	id, sendMailErr := m.sendEmail(form)
	if sendMailErr != nil {
		return sendMailErr
	}
	return m.checkStatusOfEmail(id)
}
