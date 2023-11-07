package amqp

import (
	"context"
)

// HandlerFunc represents a funcion that handles amqp messages
type HandlerFunc func(context.Context, Delivery) HandleResponse

// PreHandleFunc represents a middleware function to be called before handling amqp messages.
//
// It can alter the messaging context and even the message itself before anything is done
type PreHandleFunc func(*context.Context, *Delivery)

// PostHandleFunc represents a middleware function to be called after handling amqp messages.
//
// It has access to a copy of the message handling context, the message, and the handle response.
type PostHandleFunc func(context.Context, Delivery, HandleResponse)
