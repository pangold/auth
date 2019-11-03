package verifier

import (
	"gopkg.in/gomail.v2"
)

type Email interface {
	SendEmail(subject, content, to string) error
}

type Gomail struct {
	m       *gomail.Message
	Host     string
	Port     int
	Address  string
	Password string
}

func UseGomail(host string, port int, addr, pwd string) Email {
	return Gomail{m: gomail.NewMessage(), Host: host, Port: port, Address: addr, Password: pwd}
}

func (e Gomail) SendEmail(subject, content, to string) error {
	e.m.SetHeader("From", e.Address)
	e.m.SetHeader("To", to)
	e.m.SetHeader("Subject", subject)
	// e.m.SetHeader("Cc", cc)
	// e.m.Attach(attachPath)
	e.m.SetBody("text/html", content)
	dialer := gomail.NewDialer(e.Host, e.Port, e.Address, e.Password)
	if err := dialer.DialAndSend(e.m); err != nil {
		return err
	}
	return nil
}
