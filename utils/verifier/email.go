package verifier

import (
	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	Host     string
	Port     int
	Email    string
	Pwd      string
}

var (
	mail *gomail.Message
	conf EmailConfig
)

func init() {
	mail = gomail.NewMessage()
	conf = EmailConfig{
		Host: "smtp.163.com",
		Port: 465,
		Email: "pangold@163.com",
		Pwd: "OMEap67zeA",
	}
}

func SendEmail(subject, content, to string) error {
	mail.SetHeader("From", conf.Email)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	// mail.SetHeader("Cc", cc)
	// mail.Attach(attachPath)
	mail.SetBody("text/html", content)
	dialer := gomail.NewDialer(conf.Host, conf.Port, conf.Email, conf.Pwd)
	if err := dialer.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
