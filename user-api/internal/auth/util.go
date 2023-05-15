package auth

import (
	"errors"
	"fmt"
	"os"
	"user-api/pkg/htmlTemplateReader"
)

func generateRecoveryPasswordUrl(token string) string {
	baseUrl := os.Getenv("BASE_BACKEND_URL")
	return fmt.Sprintf("https://%s/recovery?token=%s", baseUrl, token)
}
func getBaseBackendUrl() (string, error) {
	baseUrl := os.Getenv("BASE_BACKEND_URL")
	if baseUrl == "" {
		return "", errors.New("no env variable")
	}
	return baseUrl, nil
}

func getBaseFrontendUrl() (string, error) {
	baseUrl := os.Getenv("BASE_FRONTEND_URL")
	if baseUrl == "" {
		return "", errors.New("no env variable")
	}
	return baseUrl, nil
}
func getTemplatesFilePath() (string, error) {
	return "./templates", nil
	//wd, err := os.Getwd()
	//if err != nil {
	//	return "", err
	//}
	//sep := reader.New().GetSeparator()
	//splitCwd := strings.Split(wd, sep)
	//templateFolderFilePath := strings.Join(splitCwd[:len(splitCwd)-1], sep) + sep + "templates"
	//return templateFolderFilePath, nil
}

type RecoveryUserTemplateGenerator struct {
	reader *htmlTemplateReader.HTMLReader
}

func NewRecoveryUserTemplateGenerator() *RecoveryUserTemplateGenerator {
	return &RecoveryUserTemplateGenerator{reader: htmlTemplateReader.New()}
}

func (r *RecoveryUserTemplateGenerator) GetRecoveryUserTemplate(token string) (string, error) {
	templateFolderFilePath, err := getTemplatesFilePath()
	if err != nil {
		return "", err
	}
	e, err := NewEmailTemplate(token)
	if err != nil {
		return "", err
	}
	template, err := r.reader.GetStringHtml(templateFolderFilePath, "resetPasswordEmailTemplate.html", e)
	if err != nil {
		return "", err
	}
	return template, nil
}
