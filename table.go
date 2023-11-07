package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Table represents the amqp extra arguments when declaring exchanges or binding queues
//
// When declaring an exchange or a queue in AMQP, you can include a set of optional arguments to customize its behavior
// These arguments are provided as a collection of key-value pairs, where the keys represent specific configuration options,
// and the values determine the settings for those options.
type Table amqp.Table

func (t Table) toAmqpTable() amqp.Table {
	if t == nil {
		return amqp.Table{}
	}

	return amqp.Table(t)
}
