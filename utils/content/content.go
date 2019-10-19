package content

import (
	"fmt"
)

func GetVerficationText(company, module, code string) string {
	text := fmt.Sprintf("[%s] %s，%s verfication code, available in 10 minutes", company, code, module)
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

