// This file contains all event based comms code. This is pub-sub type behaviour.

package anybus

import (
	"errors"

	"github.com/gilwil/anybus/driver"
)

// Publisher is a handle for publishing events
type Publisher struct {
	bus   *AnyBus
	topic string
}

// CreatePublisher is factory method for a publisher
func (b *AnyBus) CreatePublisher(topic string) (*Publisher, error) {
	switch {
	case b == nil:
		return nil, errors.New("nil bus")
	case topic == "":
		return nil, errors.New("nil topic")
	}

	return &Publisher{
		bus:   b,
		topic: topic,
	}, nil
}

// Post sends event to topic
func (p *Publisher) Post( /*TODO*/ ) error {
	/* TODO */
	return errors.New("not yet implemented")
}

// Subscription handle
type Subscription struct {
	bus *AnyBus
	sub driver.Subscription
}

// Subscribe starts a subscription on a topic
func (b *AnyBus) Subscribe(topic string, group string) (*Subscription, error) {
	switch {
	case b == nil:
		return nil, errors.New("nil bus")
	case topic == "":
		return nil, errors.New("nil topic")
	}

	sub, err := b.d.Subscribe(topic, group, nil /* TODO */)
	if err != nil {
		/* TODO error handling */
	}

	return &Subscription{
		bus: b,
		sub: sub,
	}, nil
}
