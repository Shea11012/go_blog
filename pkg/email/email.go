package email

import (
	"gopkg.in/gomail.v2"
)

type SMTPInfo struct {
	Host string
	Port int
	IsSSL bool
	UserName string
	Password string
	From string
}

type Email struct {
	*SMTPInfo
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{info}
}

func (e *Email) SendEmail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From",e.From)
	m.SetHeader("To",to...)
	m.SetHeader("Subject",subject)
	m.SetHeader("text/html",body)
	return nil
}
