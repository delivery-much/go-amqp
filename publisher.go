package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// amqpPublisher represents a amqp publisher that can publish messages on a exchange
type amqpPublisher struct {
	connectedStruct
	// exchangeName represents the name of the exchange that the publisher publishes messages to
	exchangeName string

	// channel its the channel that the publisher is on
	channel *amqp.Channel

	// waitConfirmation defines if the publisher is configured to wait for the server confirmation when publishing messages
	waitConfirmation bool
}

func newPublisher(exchangeName string, ch *amqp.Channel) *amqpPublisher {
	return &amqpPublisher{
		exchangeName: exchangeName,
		channel:      ch,
		connectedStruct: connectedStruct{
			ch: ch,
		},
	}
}

// Publish publishes a message on a exchange
func (p *amqpPublisher) Publish(body []byte, key string, conf ...PublishConfig) (err error) {
	c := PublishConfig{}
	if len(conf) > 0 {
		c = conf[0]
	}

	publishing := c.getPublishingFromConfig()
	publishing.Body = body

	err = p.channel.PublishWithContext(
		context.TODO(),
		p.exchangeName,
		key,
		c.Mandatory,
		c.Imediate,
		publishing,
	)
	if err != nil {
		err = fmt.Errorf("Failed to publish message, %v", err)
		return
	}

	if p.waitConfirmation && c.WaitConfirmation {
		confirmation := <-p.channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		if !confirmation.Ack {
			err = errors.New("The server did not acknowledge the message publishing")
		}
	}

	return
}

// PublishJSON publishes a json encoded struct on a exchange
func (p *amqpPublisher) PublishJSON(v any, key string, conf ...PublishConfig) (err error) {
	c := PublishConfig{}
	if len(conf) > 0 {
		c = conf[0]
	}

	body, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("Failed to encode payload to a JSON, %v", err)
	}

	publishing := c.getPublishingFromConfig()
	publishing.Body = body

	err = p.channel.PublishWithContext(
		context.TODO(),
		p.exchangeName,
		key,
		c.Mandatory,
		c.Imediate,
		publishing,
	)
	if err != nil {
		err = fmt.Errorf("Failed to publish message, %v", err)
		return
	}

	if p.waitConfirmation && c.WaitConfirmation {
		confirmation := <-p.channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		if !confirmation.Ack {
			err = errors.New("The server did not acknowledge the message publishing")
		}
	}

	return
}
