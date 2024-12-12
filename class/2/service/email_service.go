package service

import (
	"bytes"
	"context"
	"github.com/mailersend/mailersend-go"
	"html/template"
	"time"
)

type EmailService interface {
	Send(to mailersend.Recipient, subject, template string, data interface{}) (string, error)
}

type emailService struct {
	Mailer *mailersend.Mailersend
}

func NewEmailService() EmailService {
	ms := mailersend.NewMailersend("mlsn.34598713dfd7d71e7dfc181a26b904c029cbdae88595a27e42eee94f1b426811")

	return &emailService{Mailer: ms}
}

func (s *emailService) Send(to mailersend.Recipient, subject, htmlTemplate string, data interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := s.Mailer.Email.NewMessage()

	from := mailersend.From{
		Name:  "Toko Mantap",
		Email: "MS_q1mJQr@trial-351ndgw0nrd4zqx8.mlsender.net",
	}
	message.SetFrom(from)
	recipients := []mailersend.Recipient{to}
	message.SetRecipients(recipients)
	message.SetSubject(subject)

	tmpl, err := template.ParseFiles("../email/" + htmlTemplate + ".html")
	if err != nil {
		return "", err
	}

	// Apply template dengan data
	var body bytes.Buffer
	if err = tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	message.SetHTML(body.String())

	res, err := s.Mailer.Email.Send(ctx, message)
	if err != nil {
		return "", err
	}

	return res.Header.Get("X-Message-Id"), nil
}
