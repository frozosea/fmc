package auth

type EmailTemplate struct {
	ResetUrl    string
	FrontendUrl string
	BackendUrl  string
}

func NewEmailTemplate(token string) (*EmailTemplate, error) {
	resetUrl := generateRecoveryPasswordUrl(token)
	backendUrl, err := getBaseBackendUrl()
	if err != nil {
		return nil, err
	}
	frontendUrl, err := getBaseFrontendUrl()
	if err != nil {
		return nil, err
	}
	return &EmailTemplate{ResetUrl: resetUrl, FrontendUrl: frontendUrl, BackendUrl: backendUrl}, nil
}
