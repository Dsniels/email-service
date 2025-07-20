package queue

import (
	"context"
	"log"

	rb "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn *rb.Connection
	ch   *rb.Channel
}

func (r *Rabbit) Publish(ctx context.Context, body interface{}) {

}

func (r *Rabbit) StartConsuming(ctx context.Context, eventname string) {

	defer r.ch.Close()
	err := r.ch.ExchangeDeclare(eventname, rb.ExchangeFanout, true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	queue, err := r.ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = r.ch.QueueBind(queue.Name, "", eventname, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	msg, err := r.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalln(err)
	}

	var stop chan struct{}
	go func() {
		for d := range msg {
			log.Println(d)
		}
	}()
	<-stop
}

func NewRabbit() (*Rabbit, error) {
	conn, err := rb.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Rabbit{
		conn: conn,
		ch:   ch,
	}, nil

}
