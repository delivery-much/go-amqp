package amqp

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type amqpExchange struct {
	name string

	// exchangeChannel its the amqp channel that the Exchange is on
	channel *amqp.Channel

	// preHandleFuncs are the functions that will be called before the message handling
	preHandleFuncs []PreHandleFunc

	// postHandleFuncs are the functions that will be called after the message handling
	postHandleFuncs []PostHandleFunc
}

// BindQueue declares a new queue on the exchange given a queue config and binds it to the exchange
func (e *amqpExchange) BindQueue(queueName, routingKey string, conf ...QueueBindConfig) (q Queue, err error) {
	config := QueueBindConfig{}
	if len(conf) > 0 {
		config = conf[0]
	}

	_, err = e.channel.QueueDeclare(
		queueName,
		config.Durable,
		config.AutoDelete,
		config.Exclusive,
		config.NoWait,
		config.Args.toAmqpTable(),
	)
	if err != nil {
		err = fmt.Errorf("Failed to declare queue, %v", err)
		return
	}

	err = e.channel.QueueBind(
		queueName,
		routingKey,
		e.name,
		config.NoWait,
		config.Args.toAmqpTable(),
	)
	if err != nil {
		err = fmt.Errorf("Failed to bind queue, %v", err)
		return
	}

	q = &amqpQueueBind{
		name:     queueName,
		exchange: e,
		channel:  e.channel,
	}
	return
}

// Name returns the exchange name
func (e *amqpExchange) Name() string {
	return e.name
}

// PreHandleFuncs returns the pre handle funcs for the exchange
func (e *amqpExchange) PreHandleFuncs() []PreHandleFunc {
	return e.preHandleFuncs
}

// PostHandleFuncs returns the post handle funcs for the exchange
func (e *amqpExchange) PostHandleFuncs() []PostHandleFunc {
	return e.postHandleFuncs
}

// Before adds functions that will be called in the exchange before the message handling
func (e *amqpExchange) Before(funcs ...PreHandleFunc) {
	e.preHandleFuncs = append(e.preHandleFuncs, funcs...)
}

// After adds functions that will be called in the queue before the message handling
func (q *amqpExchange) After(funcs ...PostHandleFunc) {
	q.postHandleFuncs = append(q.postHandleFuncs, funcs...)
}
