package utils

import (
	"./verifier"
)

// TODO: 
// I think content generator be invoked here is better.
// ActivatedContentGenerator
// ForgotContentGenerator
// VerificationCodeGenerator

func SendActivationEmail(email, content string) error {
	return verifier.SendEmail("Account Activation", content, email)
}

func SendForgotEmail(email, content string) error {
	return verifier.SendEmail("Forgot Password", content, email)
}

func SendVerificationEmail(email, content string) error {
	return verifier.SendEmail("Verification Code", content, email)
}

func SendSMS(phone, content string) error {
	return verifier.SendSMS(phone, content)
}
