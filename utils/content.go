package utils

import (
	"./content"
)

func GetVerificationText(company, code string) string {
	return content.GetVerificationText(company, code)
}

func GetActivationText(company, url string) string {
	return content.GetActivationText(company, url)
}

func GetForgotPasswordText(company, url string) string {
	return content.GetForgotPasswordText(company, url)
}

