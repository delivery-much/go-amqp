package amqp

import (
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type connectedStruct struct {
	ch *amqp091.Channel
}

func (cs *connectedStruct) OnClose(f func(err *amqp.Error)) {
	go func() {
		ch := make(chan *amqp.Error)
		cs.ch.NotifyClose(ch)

		for err := range ch {
			f(err)
		}
	}()
}
