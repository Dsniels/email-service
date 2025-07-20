package service

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/dsniels/email-service/internal/core"
)

type EmailSvc struct {
	auth *smtp.Auth
	opts *SmtpOpts
}

func (s *EmailSvc) SendEmail(msg *core.Message) error {
	message := messageBuilder(msg)
	err := smtp.SendMail(
		s.opts.Addr,
		*s.auth,
		s.opts.Sender,
		msg.To,
		message,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Email Sent...")

	return nil
}

func messageBuilder(mail *core.Message) []byte {
	var msg string
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Content)
	return []byte(msg)
}

func NewEmailSvc() *EmailSvc {
	auth, opts := GetAuth()
	return &EmailSvc{auth: auth, opts: opts}

}
