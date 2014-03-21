package domain

import (
	"bytes"
	"fmt"
	"github.com/robfig/revel"
	"net/smtp"
	"os"
	"text/template"
)

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}

type SmtpTemplateData struct {
	From  string
	To    string
	Name  string
	Email string
	Body  string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}}
Subject: You got an email from your website

Somebody named {{.Name}} at '{{.Email}}' sent you a
message through the contact form at your website

{{.Body}}
`

func SendMail(name, email, message string) error {
	emailUsername := os.Getenv("EMAIL_USERNAME")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailTo := os.Getenv("EMAIL_DESTINATION")
	emailServer, _ := revel.Config.String("email.server")
	emailPort, _ := revel.Config.Int("email.port")

	emailUser := &EmailUser{emailUsername, emailPassword, emailServer, emailPort}
	auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.EmailServer)

	context := &SmtpTemplateData{emailUser.Username, emailTo, name, email, message}
	t := template.New("emailTemplate")
	t, err := t.Parse(emailTemplate)
	if err != nil {
		return err
	}

	var doc bytes.Buffer
	err = t.Execute(&doc, context)
	if err != nil {
		return err
	}

	err = smtp.SendMail(fmt.Sprintf("%v:%v", emailUser.EmailServer, emailUser.Port),
		auth,
		emailUser.Username,
		[]string{emailTo},
		doc.Bytes())

	if err != nil {
		return err
	}
	return nil
}
