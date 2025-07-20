package api

import (
	"github.com/dsniels/email-service/internal/api/handler"
	"github.com/dsniels/email-service/internal/events"
	"github.com/dsniels/email-service/internal/service"
)

type App struct {
	H   *handler.EmailHandler
	Rab *events.Rabbit
}

func NewApp() (*App, error) {

	svc := service.NewEmailSvc()
	hdl := handler.NewEmailHandler(svc)
	rabbit, err := events.NewRabbit()
	if err != nil {
		return nil, err
	}

	return &App{
		Rab: rabbit,
		H:   hdl,
	}, nil
}
