package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dsniels/email-service/internal/core"
	"github.com/dsniels/email-service/internal/service"
	rb "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn     *rb.Connection
	ch       *rb.Channel
	emailSvc *service.EmailSvc
}

func (r *Rabbit) Publish(ctx context.Context, body interface{}) {

}

func (r *Rabbit) StartConsuming(ctx context.Context, eventname string) {

	defer r.ch.Close()
	err := r.ch.ExchangeDeclare(eventname, rb.ExchangeFanout, false, true, false, false, nil)
	if err != nil {
		log.Fatalln("Exchange: ", err)
	}

	queue, err := r.ch.QueueDeclare("", true, false, true, false, nil)
	if err != nil {
		log.Fatalln("QueueDeclare: ", err)
	}

	err = r.ch.QueueBind(queue.Name, "", eventname, false, nil)
	if err != nil {
		log.Fatalln("Bindin: ", err)
	}

	msg, err := r.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("Consume: ", err)
	}

	var stop chan struct{}
	go func() {
		for d := range msg {
			message := new(core.Message)
			json.Unmarshal(d.Body, message)
			if err := r.emailSvc.SendEmail(message); err!=nil{
				log.Println(err)
			}
			
		}
	}()
	<-stop
}

func NewRabbit(svc *service.EmailSvc) (*Rabbit, error) {
	conn, err := rb.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Rabbit{
		emailSvc: svc,
		conn:     conn,
		ch:       ch,
	}, nil

}
