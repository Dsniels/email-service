package service

import (
	"net/smtp"
	"os"
)

type SmtpOpts struct {
	Host     string
	Port     string
	Email    string
	Addr     string
	Password string
	Sender   string
}

func GetAuth() (*smtp.Auth, *SmtpOpts) {
	s := &SmtpOpts{
		Host:     os.Getenv("HOST"),
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Addr:     os.Getenv("HOST") + os.Getenv("PORT"),
		Sender:   os.Getenv("SENDER"),
	}

	auth := smtp.PlainAuth("", s.Email, s.Password, s.Host)

	return &auth, s

}
