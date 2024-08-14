package amqp

import (
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// client represents the client with connection to AMQP.
type client struct {
	conn *amqp.Connection
}

// NewClient connects to the AMQP server using the provided configuration, and returns the AMQP Client.
func NewClient(URL string, conf ...Config) (c Client, err error) {
	amqpConfig := amqp.Config{}
	if len(conf) > 0 {
		amqpConfig = conf[0].toAMQPConfig()
	}

	conn, err := amqp.DialConfig(URL, amqpConfig)
	if err != nil {
		err = fmt.Errorf("Error when connecting to AMQP, %v", err)
		return
	}

	c = &client{conn}
	return
}

// Close will close the rabbitmq connection.
func (c *client) Close() (err error) {
	if c.conn != nil {
		err = c.conn.Close()
	}

	return
}

// Ping checks the rabbitmq connection health
func (c *client) Ping() (err error) {
	if c.conn == nil {
		err = errors.New("The AMQP connection is not open")
		return
	}

	if c.conn.IsClosed() {
		err = errors.New("AMQP disconnected")
	}

	return
}

// StartExchange starts a amqp exchange and returns a channel with the exchange declared
func (c *client) StartExchange(exchangeName string, exchangeType ExchangeType, conf ...ExchangeConfig) (e Exchange, err error) {
	if c.conn == nil {
		err = errors.New("The AMQP connection is not open")
		return
	}

	ch, err := c.conn.Channel()
	if err != nil {
		err = fmt.Errorf("Failed to create a new channel for the %s exchange, %v", exchangeName, err)
		return
	}

	config := ExchangeConfig{}
	if len(conf) > 0 {
		config = conf[0]
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeType.ToString(),
		config.Durable,
		config.AutoDelete,
		config.Internal,
		config.NoWait,
		config.Args.toAmqpTable(),
	)

	e = newExchange(exchangeName, ch)
	return
}

// CreatePublisher creates a new publisher to publish messages on an exchange
func (c *client) CreatePublisher(exchangeName string, NoWait ...bool) (p Publisher, err error) {
	waitConfirmation := true
	if len(NoWait) > 0 {
		waitConfirmation = !NoWait[0]
	}

	channel, err := c.conn.Channel()
	if err != nil {
		err = fmt.Errorf("Failed to create a new channel for the %s exchange publisher, %v", exchangeName, err)
		return
	}

	if waitConfirmation {
		err = channel.Confirm(false)
		if err != nil {
			err = fmt.Errorf("Failed to set the %s exchange publisher into confirmation mode. Try setting the NoWait flag as true. %v", exchangeName, err)
			return
		}
	}

	p = newPublisher(exchangeName, channel)
	return
}
