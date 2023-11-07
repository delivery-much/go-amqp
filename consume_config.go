package amqp

// ConsumeConfig represents the configuration that can be provided when consuming a queue
type ConsumeConfig struct {
	// When AutoAck is set to true, it means that as soon as a message is delivered to the consumer,
	// RabbitMQ automatically considers the message as acknowledged (ack) without the consumer having to send an explicit acknowledgment.
	//
	// default: false
	AutoAck bool

	// When a consumer is declared as Exclusive,
	// it means that no other consumer can access the same queue on the same channel/connection.
	// Exclusive consumers are often used for scenarios where you want to ensure that
	// only one consumer processes messages from a specific queue.
	//
	// default: false
	Exclusive bool

	// The NoLocal parameter, when set to true,
	// prevents a consumer from receiving messages that it publishes to the same connection.
	// In other words, it prevents consumers from consuming their own messages.
	//
	// default: false
	NoLocal bool

	// When NoWait is set to True,
	// it means that the method will not wait for a response from the server to confirm the consumption.
	// In other words, it makes the declaration non-blocking.
	// If any error occurs during the declaration, it won't be reported immediately.
	//
	// default: false
	NoWait bool

	// The name for the queue consumer.
	// When a consumer name is not provided, the library will generate one based on the queue information.
	ConsumerName string

	// When declaring an consumer in AMQP, you can include a set of optional arguments to customize its behavior
	// These arguments are provided as a collection of key-value pairs, where the keys represent specific configuration options,
	// and the values determine the settings for those options.
	Args Table
}
