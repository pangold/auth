package verifier

import (
	"fmt"
	"testing"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func TestSendEmail(t *testing.T) {
	m := gomail.NewMessage()
	// 
	m.SetHeader("From", "panking@126.com")
	//
	m.SetHeader("To", "panking@126.com")
	//
	m.SetAddressHeader("Cc", "pangold@163.com", "pangold")
	//
	m.SetHeader("Subject", "Hello")
	//
	m.SetBody("text/html", "Golang Email Test...")
	//
	// m.Attach("/home/dora/Desktop/888.jpeg")
	//
	d := gomail.NewDialer("pop.126.com", 110, "panking@126.com", "******")
	// d := gomail.NewPlainDialer("smtp.163.com", 465, "pangold@163.com", "******")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("Send email successfully")
}
