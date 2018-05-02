// Package driver defines the interface for bus drivers
package driver

// HandleFunc is a receive message handler
type HandleFunc func(msg []byte)

// Subscription is managed by the driver and is a handle to a subscription
type Subscription interface {
	ImASubscription()
}

// Driver interfaces provides interaction with the actual bus
type Driver interface {

	// Subcribe activates a subscription
	Subscribe(topic string, group string, f HandleFunc) (Subscription, error)

	// Unsubscribe deactivates a subscription
	Unsubscripte(Subscription)

	// Post sends a message to a topic
	Post(topic string, buff []byte) error
}
