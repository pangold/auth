package utils

import (
	"./verifier"
)

func SendEmail(email, content string) error {
	return verifier.SendEmail(email, content)
}

func SendSMS(phone, content string) error {
	return verifier.SendSMS(phone, content)
}
