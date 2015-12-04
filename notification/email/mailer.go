package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	dialer *gomail.Dialer
}

func NewMailer() *Mailer {
	m := &Mailer{
		gomail.NewPlainDialer("smtp.easervice.com.br", 587, "erichnascimento@easervice.com.br", "pid96sqdi"),
	}
	m.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return m
}

func (m *Mailer) SendMail(content string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "erichnascimento@easervice.com.br")
	msg.SetHeader("To", "erichnascimento@gmail.com")
	msg.SetHeader("Subject", "Cloud-Watch Notification")
	msg.SetBody("text/html", content)

	return m.dialer.DialAndSend(msg)
}
