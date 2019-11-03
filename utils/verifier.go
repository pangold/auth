package utils

import (
	"./verifier"
)

// TODO: 
// I think content generator be invoked here is better.
// ActivatedContentGenerator
// ForgotContentGenerator
// VerificationCodeGenerator

var (
	e verifier.Email
)

func getEmail() verifier.Email {
	if e == nil {
		e = verifier.UseGomail(GetConfig().Email.Host, GetConfig().Email.Port, GetConfig().Email.Address, GetConfig().Email.Password)
	}
	return e
}

func SendActivationEmail(email, content string) error {
	return getEmail().SendEmail("Account Activation", content, email)
}

func SendForgotEmail(email, content string) error {
	return getEmail().SendEmail("Forgot Password", content, email)
}

func SendVerificationEmail(email, content string) error {
	return getEmail().SendEmail("Verification Code", content, email)
}

func SendSMS(phone, content string) error {
	return verifier.SendSMS(phone, content)
}
