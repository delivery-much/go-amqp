package amqp

type PublishConfig struct {
	// When you set the mandatory flag to true while publishing a message,
	// it indicates that the message must be routed to at least one queue.
	// If the message cannot be routed to any queue, RabbitMQ will return the message to the publisher.
	// This flag is typically used when you want to ensure that your message is not lost and must be delivered to at least one queue.
	//
	// default: false
	Mandatory bool

	// When you set the immediate flag to true while publishing a message,
	// it indicates that the message should be delivered to a consumer as soon as possible.
	// If there are no available consumers to immediately accept the message, RabbitMQ will return the message to the publisher.
	//
	// default: false
	Imediate bool

	// Message publishing is, by default, an asynchronous event.
	// However, when the WaitConfirmation flag is set to true, and the publisher was created in confirmation mode,
	// the Publish method will wait for the server to return a response confirming that the message was indeed published.
	// It is important to note that this could potentially increase the message publishing time, depending on your server's latency.
	//
	// default: false
	WaitConfirmation bool
}
