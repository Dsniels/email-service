package api

import (
	"github.com/dsniels/email-service/internal/api/handler"
	"github.com/dsniels/email-service/internal/queue"
	"github.com/dsniels/email-service/internal/service"
)

type App struct {
	H   *handler.EmailHandler
	Rab *queue.Rabbit
}

func NewApp() (*App, error) {

	svc := service.NewEmailSvc()
	hdl := handler.NewEmailHandler(svc)
	rabbit, err := queue.NewRabbit(svc)
	if err != nil {
		return nil, err
	}

	return &App{
		Rab: rabbit,
		H:   hdl,
	}, nil
}
