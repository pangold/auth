package content

import (
	"fmt"
)

func GetVerificationText(company, code string) string {
	text := fmt.Sprintf("[%s]，verfication code %s, available in 10 minutes", company, code)
	return text
}

func GetActivationText(company, url string) string {
	text := fmt.Sprintf("[%s]<br><br>Click the link below to activate your account. <br><br>%s", company, url)
	return text
}

func GetForgotPasswordText(company, url string) string {
	text := fmt.Sprintf("[%s]\n\nClick the link below to reset your password. This link will be available in 10 minutes.\n %s", company, url)
	return text
}

