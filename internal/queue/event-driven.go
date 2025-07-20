package queue

import "context"

type EventDriven interface {
	Publish(ctx context.Context, body interface{})
	StartConsuming(ctx context.Context, queueName string)
}
