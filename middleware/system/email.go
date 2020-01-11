package system

import (
	"gitlab.com/pangold/auth/config"
	"gopkg.in/gomail.v2"
)

type DefaultEmail struct {
	message *gomail.Message
	config   config.Email
}

func NewDefaultEmail(c config.Email) *DefaultEmail {
	return &DefaultEmail{
		message: gomail.NewMessage(),
		config: c,
	}
}

func (this *DefaultEmail) SendVerificationCode(to, vcode string) error {
	return this.SendEmail("Verification Code", vcode, to)
}

func (this *DefaultEmail) SendActivationEmail(to, url string) error {
	return this.SendEmail("Account Activation", url, to)
}

func (this *DefaultEmail) SendResetPasswordEmail(to, url string) error {
	return this.SendEmail("Reset Password", url, to)
}

func (this *DefaultEmail) SendEmail(subject, content, to string) error {
	this.message.SetHeader("From", this.config.Address)
	this.message.SetHeader("To", to)
	this.message.SetHeader("Subject", subject)
	// e.m.SetHeader("Cc", cc)
	// e.m.Attach(attachPath)
	this.message.SetBody("text/html", content)
	dialer := gomail.NewDialer(this.config.Host, this.config.Port, this.config.Address, this.config.Password)
	if err := dialer.DialAndSend(this.message); err != nil {
		return err
	}
	return nil
}