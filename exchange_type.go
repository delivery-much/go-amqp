package amqp

// ExchangeType represents a exchange type that defines its behaviour
type ExchangeType string

const (
	// A direct exchange routes messages to queues based on a specified routing key.
	// When a message is published to a direct exchange with a particular routing key,
	// it will be delivered to the queue(s) that are bound to the same exchange with a matching routing key.
	ExchangeTypeDirect = ExchangeType("direct")

	// A fanout exchange broadcasts messages to all queues that are bound to it,
	// regardless of routing keys. It is a simple publish/subscribe mechanism where
	// all queues receive a copy of each message sent to the exchange.
	ExchangeTypeFanout = ExchangeType("fanout")

	// A topic exchange is more flexible than direct exchanges.
	// It routes messages to queues based on wildcard patterns in the routing keys.
	// Routing keys can include wildcards like '*' (matches one word) and '#' (matches zero or more words),
	// allowing for complex message routing based on patterns.
	ExchangeTypeTopic = ExchangeType("topic")

	// Headers exchanges use message header attributes to determine message routing, rather than routing keys.
	// The exchange will match headers against predefined criteria to determine which queues should receive the message.
	ExchangeTypeHeaders = ExchangeType("headers")
)

// ToString returns the string notation of the exchange type
func (t ExchangeType) ToString() string {
	return string(t)
}
