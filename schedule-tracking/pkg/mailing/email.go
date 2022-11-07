package mailing

import (
	"encoding/json"
	"errors"
	"fmt"
	file_reader "github.com/frozosea/file-reader"
	"io"
	"net/http"
	"net/url"
	"schedule-tracking/pkg/logging"
)

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
	signature       string
}

func NewMailing(logger logging.ILogger, senderName string, senderEmail string, unisenderApiKey string, signature string) *Mailing {
	return &Mailing{reader: file_reader.New(), logger: logger, senderName: senderName, senderEmail: senderEmail, UnisenderApiKey: unisenderApiKey, signature: signature}
}

func (m *Mailing) getForm(toAddress, subject, fileName, body string, file string) url.Values {
	query := url.Values{}
	query.Set("format", "json")
	query.Set("api_key", m.UnisenderApiKey)
	query.Set("sender_name", m.senderName)
	query.Set("email", toAddress)
	query.Set("sender_email", m.senderEmail)
	query.Set("subject", subject)
	query.Set("body", body)
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
	form := m.getForm(toAddress, subject, fileName, m.signature, string(readFile))
	id, sendMailErr := m.sendEmail(form)
	if sendMailErr != nil {
		return sendMailErr
	}
	return m.checkStatusOfEmail(id)
}
