package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dsniels/email-service/internal/core"
	"github.com/dsniels/email-service/pkg"
)

func (e *EmailHandler) HandleSendEmail(w http.ResponseWriter, r *http.Request) {

	msg := new(core.Message)
	err := json.NewDecoder(r.Body).Decode(msg)
	if err != nil {
		pkg.BadRequestError(err.Error())
	}
	err = e.smtpSvc.SendEmail(msg)
	if err != nil {
		pkg.BadRequestError(err.Error())
	}

	pkg.WriteReponse(w, http.StatusOK, struct{}{})

}
