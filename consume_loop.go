package amqp

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// consumeLoop its the function that will be called whenever a message is consumed.
// - calls the exchange and queue middlewares
// - calls the handlerFunc to consume the message
// - treats the messaging response
func consumeLoop(deliveries <-chan amqp.Delivery, q Queue, handlerFn HandlerFunc) {
	for d := range deliveries {
		msg := Delivery(d)
		ctx := context.TODO()

		for _, preFunc := range q.Exchange().PreHandleFuncs() {
			preFunc(&ctx, &msg)
		}
		for _, preFunc := range q.PreHandleFuncs() {
			preFunc(&ctx, &msg)
		}

		res := handlerFn(ctx, msg)

		for _, postFunc := range q.Exchange().PostHandleFuncs() {
			postFunc(ctx, msg, res)
		}
		for _, postFunc := range q.PostHandleFuncs() {
			postFunc(ctx, msg, res)
		}

		if res.Nack {
			_ = d.Nack(false, true)
			continue
		}

		_ = d.Ack(false)
	}
}
