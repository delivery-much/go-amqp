package amqp

// ExchangeConfig represents the configuration for a rabbitmq exchange
type ExchangeConfig struct {
	// When a exchange is declared as durable,
	// it means that RabbitMQ will make efforts to ensure that the component survives server restarts or failures.
	// Messages and configuration related to durable components are stored on disk.
	// This is useful when you want to ensure that important data is not lost in case of server failures.
	//
	// default: false
	Durable bool

	// An auto-delete exchange is automatically deleted by RabbitMQ once there are no consumers or bindings left for it.
	// This is often used for temporary components that are only needed for a specific duration or purpose.
	// It's a way to clean up resources automatically when they are no longer needed.
	//
	// default: false
	AutoDelete bool

	// When an exchange is declared as internal,
	// it means that it can't be directly published to by clients.
	// Internal exchanges are used for internal RabbitMQ mechanisms and cannot be used for normal publishing of messages.
	// They are useful for building advanced routing topologies within RabbitMQ.
	//
	// default: false
	Internal bool

	// When you declare a exchange with the "NoWait" option,
	// it means that the method will not wait for a response from the server to confirm the declaration.
	// This can improve declaration speed but comes with the trade-off that you won't receive an immediate response indicating success or failure.
	// It's often used when you want to declare components quickly and are willing to skip the confirmation step.
	//
	// default: false
	NoWait bool

	// When declaring an exchange in AMQP, you can include a set of optional arguments to customize the behavior of the exchange
	// These arguments are provided as a collection of key-value pairs, where the keys represent specific configuration options,
	// and the values determine the settings for those options.
	Args Table
}
