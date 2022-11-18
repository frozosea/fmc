package auth

import (
	"fmt"
	"os"
)

func generateRecoveryPasswordUrl(token string) string {
	baseUrl := os.Getenv("BASE_DOMAIN_URL")
	return fmt.Sprintf("%s/recovery?token=%s", baseUrl, token)
}
