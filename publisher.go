package amqp

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// amqpPublisher represents a amqp publisher that can publish messages on a exchange
type amqpPublisher struct {
	// exchangeName represents the name of the exchange that the publisher publishes messages to
	exchangeName string

	// channel its the channel that the publisher is on
	channel *amqp.Channel

	// waitConfirmation defines if the publisher is configured to wait for the server confirmation when publishing messages
	waitConfirmation bool
}

// Publish publishes a message on a exchange
func (p *amqpPublisher) Publish(m Publishing, key string, conf ...PublishConfig) (err error) {
	c := PublishConfig{}
	if len(conf) > 0 {
		c = conf[0]
	}

	err = p.channel.PublishWithContext(
		context.TODO(),
		p.exchangeName,
		key,
		c.Mandatory,
		c.Imediate,
		amqp.Publishing(m),
	)
	if err != nil {
		err = fmt.Errorf("Failed to publish message, %v", err)
		return
	}

	if p.waitConfirmation && c.WaitConfirmation {
		confirmation := <-p.channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		if !confirmation.Ack {
			err = fmt.Errorf("The server did not acknowledge the message publishing, %v", err)
		}
	}

	return
}
