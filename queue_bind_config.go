package amqp

// QueueConfig represents the configuration for binding a queue to an exchange
type QueueBindConfig struct {
	// When a queue is declared as durable,
	// it means that RabbitMQ will make efforts to ensure that the component survives server restarts or failures.
	// Messages and configuration related to durable components are stored on disk.
	// This is useful when you want to ensure that important data is not lost in case of server failures.
	//
	// default: false
	Durable bool

	// An auto-delete queue is automatically deleted by RabbitMQ once there are no consumers or bindings left for it.
	// This is often used for temporary components that are only needed for a specific duration or purpose.
	// It's a way to clean up resources automatically when they are no longer needed.
	//
	// default: false
	AutoDelete bool

	// When you declare a queue as exclusive,
	// it means that the queue can only be accessed by the current connection.
	// The queue will be automatically deleted by RabbitMQ when the connection that declared it is closed.
	// Exclusive queues are often used in scenarios where you want to ensure that a queue is used only by a single consumer or where you want to create a temporary work queue for a particular client or session.
	//
	// When you declare a queue as not exclusive, multiple connections can access it concurrently.
	// The queue will not be automatically deleted when the connection that declared it is closed.
	// It will remain in RabbitMQ until it is explicitly deleted or until the server decides to remove it due to other factors, such as lack of use or expiration.
	// Non-exclusive queues are typically used for more persistent, shared message processing scenarios where multiple consumers may need to access the same queue.
	//
	// default: false
	Exclusive bool

	// When you declare a queue with the "NoWait" option,
	// it means that the method will not wait for a response from the server to confirm the declaration.
	// This can improve declaration speed but comes with the trade-off that you won't receive an immediate response indicating success or failure.
	// It's often used when you want to declare components quickly and are willing to skip the confirmation step.
	//
	// default: false
	NoWait bool

	// When declaring an queue in RabbitMQ, you can include a set of optional arguments to customize its behavior
	// These arguments are provided as a collection of key-value pairs, where the keys represent specific configuration options,
	// and the values determine the settings for those options.
	Args Table
}
