package amqp

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// PublishConfig represents the configuration to publish a message
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

	// Message specific fields

	// Application or exchange specific fields,
	// the headers exchange will inspect this field.
	Headers Table

	ContentType     string    // MIME content type
	ContentEncoding string    // MIME content encoding
	DeliveryMode    uint8     // Transient (0 or 1) or Persistent (2)
	Priority        uint8     // 0 to 9
	CorrelationId   string    // correlation identifier
	ReplyTo         string    // address to to reply to (ex: RPC)
	Expiration      string    // message expiration spec
	MessageId       string    // message identifier
	Timestamp       time.Time // message timestamp
	Type            string    // message type name
	UserId          string    // creating user id - ex: "guest"
	AppId           string    // creating application id
}

// getPublishingFromConfig returns a new amqp.Publishing struct with no body, using the PublishConfig values
func (c PublishConfig) getPublishingFromConfig() amqp.Publishing {
	return amqp.Publishing{
		ContentType:     c.ContentType,
		ContentEncoding: c.ContentEncoding,
		DeliveryMode:    c.DeliveryMode,
		Priority:        c.Priority,
		CorrelationId:   c.CorrelationId,
		ReplyTo:         c.ReplyTo,
		Expiration:      c.Expiration,
		MessageId:       c.MessageId,
		Timestamp:       c.Timestamp,
		Type:            c.Type,
		UserId:          c.UserId,
		AppId:           c.AppId,
	}
}
