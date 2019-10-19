package utils

import (
	"./content"
)

func GetVerificationText(company, module, code string) string {
	return content.GetVerificationText(company, module, code)
}

func GetActivationText(company, url string) string {
	return content.GetActivationText(company, url)
}

func GetForgotPasswordText(company, url string) string {
	return content.GetForgotPasswordText(company, url)
}

