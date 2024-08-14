package amqp

import (
	"crypto/tls"
	"net"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Authentication interface provides a means for different SASL authentication
// mechanisms to be used during connection tuning.
type Authentication interface {
	Mechanism() string
	Response() string
}

// Config represents the optional configuration that the user
// can provide when connecting to the AMQP client
type Config struct {
	// The SASL mechanisms to try in the client request, and the successful
	// mechanism used on the Connection object.
	// If SASL is nil, PlainAuth from the URL is used.
	SASL []Authentication

	// Vhost specifies the namespace of permissions, exchanges, queues and
	// bindings on the server.  Dial sets this to the path parsed from the URL.
	Vhost string

	ChannelMax int           // 0 max channels means 2^16 - 1
	FrameSize  int           // 0 max bytes means unlimited
	Heartbeat  time.Duration // less than 1s uses the server's interval

	// TLSClientConfig specifies the client configuration of the TLS connection
	// when establishing a tls transport.
	// If the URL uses an amqps scheme, then an empty tls.Config with the
	// ServerName from the URL is used.
	TLSClientConfig *tls.Config

	// Properties is table of properties that the client advertises to the server.
	// This is an optional setting - if the application does not set this,
	// the underlying library will use a generic set of client properties.
	Properties Table

	// Connection locale that we expect to always be en_US
	// Even though servers must return it as per the AMQP 0-9-1 spec,
	// we are not aware of it being used other than to satisfy the spec requirements
	Locale string

	// Dial returns a net.Conn prepared for a TLS handshake with TSLClientConfig,
	// then an AMQP connection handshake.
	// If Dial is nil, net.DialTimeout with a 30s connection and 30s deadline is
	// used during TLS and AMQP handshaking.
	Dial func(network, addr string) (net.Conn, error)
}

func (c Config) toAMQPConfig() amqp.Config {
	sasl := []amqp.Authentication{}
	for _, auth := range c.SASL {
		converted, ok := auth.(amqp.Authentication)
		if ok {
			sasl = append(sasl, converted)
		}
	}

	return amqp.Config{
		SASL:            sasl,
		Vhost:           c.Vhost,
		ChannelMax:      c.ChannelMax,
		FrameSize:       c.FrameSize,
		Heartbeat:       c.Heartbeat,
		TLSClientConfig: c.TLSClientConfig,
		Properties:      c.Properties.toAmqpTable(),
		Locale:          c.Locale,
		Dial:            c.Dial,
	}
}
