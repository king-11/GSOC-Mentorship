package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"strings"

	"github.com/king-11/mentorship/extractor"
)

const (
	// smtp server configuration
	SMTPHOST = "smtp.gmail.com"
	SMTPPORT = "587"
)

type Sender struct {
	Email string
	// app password
	Password string
}

func NewSender(Email, Password string) *Sender {
	return &Sender{Email, Password}
}

func (sender *Sender) getAuth() smtp.Auth {
	return smtp.PlainAuth("", sender.Email, sender.Password, SMTPHOST)
}

func (sender *Sender) SendMail(Dest []string, bodyMessage []byte) error {
	auth := sender.getAuth()
	return smtp.SendMail(
		SMTPHOST+":"+SMTPPORT,
		auth,
		sender.Email,
		Dest,
		bodyMessage,
	)
}

func getTemplate(name string) (*template.Template, error) {
	return template.ParseFiles(name)
}

func (sender *Sender) writeEmail(dest []string, contentType, subject string, bodyMessage extractor.BasicMentorship, template string) []byte {

	receipient := strings.Join(dest, ",")
	header := make(map[string]string)
	header["From"] = sender.Email
	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Disposition"] = "inline"

	var message bytes.Buffer
	for key, value := range header {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}

	t, err := getTemplate(template)
	if err != nil {
		log.Fatal(err)
	}

	message.WriteString("\r\n")
	err = t.Execute(&message, bodyMessage)
	if err != nil {
		log.Fatal(err)
	}
	return message.Bytes()
}

func (sender *Sender) WriteHTMLEmail(dest []string, subject string, bodyMessage extractor.BasicMentorship, template string) []byte {
	return sender.writeEmail(dest, "text/html", subject, bodyMessage, template)
}
