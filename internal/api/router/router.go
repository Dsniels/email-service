package router

import (
	"net/http"

	"github.com/dsniels/email-service/internal/api"
)

func InitRoutes(a *api.App) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /api/email/SendEmail", a.H.HandleSendEmail)

	return router
}
