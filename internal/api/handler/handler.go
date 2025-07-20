package handler

import "github.com/dsniels/email-service/internal/service"

type EmailHandler struct {
	smtpSvc *service.EmailSvc
}

func NewEmailHandler(mailSvc *service.EmailSvc) *EmailHandler {
	return &EmailHandler{
		smtpSvc: mailSvc,
	}
}
