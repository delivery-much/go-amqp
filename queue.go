package amqp

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// amqpQueueBind represents an amqp queue that is bound to an exchange
type amqpQueueBind struct {
	name,
	routingKey string

	// exchange its the exchange that the queue is on
	exchange Exchange

	// exchangeChannel its the amqp channel that the Queue is on
	channel *amqp.Channel

	// preHandleFuncs are the functions that will be called before the message handling
	preHandleFuncs []PreHandleFunc

	// postHandleFuncs are the functions that will be called after the message handling
	postHandleFuncs []PostHandleFunc
}

// Consume subscribes a consumer in the routing key to handle the messages with the handler function
func (q *amqpQueueBind) Consume(handlerFn HandlerFunc, conf ...ConsumeConfig) (err error) {
	queueName := q.name
	exchangeName := q.exchange.Name()
	config := ConsumeConfig{}
	if len(conf) > 0 {
		config = conf[0]
	}

	consumerName := config.ConsumerName
	if consumerName == "" {
		consumerName = fmt.Sprintf("%s-%s-%s-consumer", exchangeName, queueName, q.routingKey)
	}

	msgs, err := q.channel.Consume(
		queueName,
		consumerName,
		config.AutoAck,
		config.Exclusive,
		config.NoLocal,
		config.NoWait,
		config.Args.toAmqpTable(),
	)
	if err != nil {
		err = fmt.Errorf("Failed to consume queue, %v", err)
		return
	}

	go consumeLoop(msgs, q, handlerFn)
	return
}

// Before adds functions that will be called in the queue before the message handling
func (q *amqpQueueBind) Before(funcs ...PreHandleFunc) {
	q.preHandleFuncs = append(q.preHandleFuncs, funcs...)
}

// After adds functions that will be called in the queue before the message handling
func (q *amqpQueueBind) After(funcs ...PostHandleFunc) {
	q.postHandleFuncs = append(q.postHandleFuncs, funcs...)
}

// Name returns the queue name
func (q *amqpQueueBind) Name() string {
	return q.name
}

// Exchange returns the queue exchange
func (q *amqpQueueBind) Exchange() Exchange {
	return q.exchange
}

// RoutingKey returns the queue routing-key that was used to bind to the exchange
func (q *amqpQueueBind) RoutingKey() string {
	return q.routingKey
}

// PreHandleFuncs returns the pre handle funcs for the queue
func (q *amqpQueueBind) PreHandleFuncs() []PreHandleFunc {
	return q.preHandleFuncs
}

// PostHandleFuncs returns the post handle funcs for the queue
func (e *amqpQueueBind) PostHandleFuncs() []PostHandleFunc {
	return e.postHandleFuncs
}
