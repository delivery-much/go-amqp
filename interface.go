package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// ConnectedStruct represents a struct that's connected to an AMQP channel
type ConnectedStruct interface {
	// OnClose defines the function to execute when the AMQP channel for this struct is closed in any way.
	//
	// OnClose will start a new goroutine that listens to this connection 'closed' events.
	OnClose(f func(err *amqp.Error))
}

// Client represents a Client connection, and contains functions to manage the AMQP client
type Client interface {
	// Close closes the AMQP connection
	Close() error
	// Ping checks if the AMQP connection is active
	Ping() error

	// StartExchange starts a AMQP exchange with its own channel and returns the exchange as an entity
	StartExchange(exchangeName string, exchangeType ExchangeType, conf ...ExchangeConfig) (Exchange, error)

	// CreatePublisher creates a new publisher with its own channel to publish messages on an exchange, given the exchange name.
	//
	// When the optional NoWait flag is set to true, the publisher will not be created in confirmation mode.
	// This means that when a message is published using this publisher, the library will not wait for confirmation a from the server.
	//
	// The NoWait flag should be used when your server does not support publishers in confirmation mode, or when you specifically want the publisher to be asynchronous.
	CreatePublisher(exchangeName string, NoWait ...bool) (Publisher, error)
}

// Exchange represents a AMQP message exchange
type Exchange interface {
	ConnectedStruct
	// BindQueue declares a new queue on the exchange given a queue config and binds it to the exchange
	BindQueue(queueName, routingKey string, conf ...QueueBindConfig) (Queue, error)

	// Before adds functions that will be called in the exchange before the message handling
	Before(funcs ...PreHandleFunc)

	// After adds functions that will be called in the exchange after the message handling
	After(funcs ...PostHandleFunc)

	// Name returns the exchange name
	Name() string
	// PreHandleFuncs returns the pre handle funcs for the exchange
	PreHandleFuncs() []PreHandleFunc
	// PostHandleFuncs returns the post handle funcs for the exchange
	PostHandleFuncs() []PostHandleFunc
}

// Queue represents a AMQP queue
type Queue interface {
	// Consume subscribes a consumer in the routing key to handle the messages.
	//
	// Consume will start a new goroutine that listens to message publishings
	// and handles them with the provided handler function
	Consume(handlerFn HandlerFunc, conf ...ConsumeConfig) error

	// Before adds functions that will be called in the queue before the message handling
	Before(funcs ...PreHandleFunc)

	// After adds functions that will be called in the queue after the message handling
	After(funcs ...PostHandleFunc)

	// Name returns the queue name
	Name() string
	// Exchange returns the Exchange that the queue is on
	Exchange() Exchange
	// RoutingKey returns the queue routing-key that was used to bind to the exchange
	RoutingKey() string
	// PreHandleFuncs returns the pre handle funcs for the queue
	PreHandleFuncs() []PreHandleFunc
	// PostHandleFuncs returns the post handle funcs for the queue
	PostHandleFuncs() []PostHandleFunc
}

// Publisher represents a AMQP message publisher
type Publisher interface {
	ConnectedStruct
	// Publish publishes a message payload, in bytes format, on the publisher exchange.
	// The user can also provide a routing-key to publish the message and some extra configuration for that message, if needed.
	//
	// It is important to note that the message publishing, by default, is asynchronous.
	// However, you can make it synchronous by setting the WaitConfirmation flag from the PublishConfig as true.
	Publish(payload []byte, key string, conf ...PublishConfig) error

	// PublishJSON encodes the 'payload' param into a json string, and publishes it on the publisher exchange.
	// The user can also provide a routing-key to publish the message and some extra configuration for that message, if needed.
	//
	// It is important to note that the message publishing, by default, is asynchronous.
	// However, you can make it synchronous by setting the WaitConfirmation flag from the PublishConfig as true.
	PublishJSON(payload any, key string, conf ...PublishConfig) error
}
