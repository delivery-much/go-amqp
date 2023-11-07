package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Config represents the optional configuration that the user
// can provide when connecting to the AMQP client
type Config amqp.Config
