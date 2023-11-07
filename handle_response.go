package amqp

// HandleResponse represents the response when handling a message
type HandleResponse struct {
	// Nack defines if the message should NOT be acknowledged, and should be requeued (default: false)
	Nack bool
	// Err is the error that could have occurred during the message handling (default: nil)
	Err error
}
