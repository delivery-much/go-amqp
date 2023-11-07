package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Delivery represents a AMQP message that was received
type Delivery amqp.Delivery
