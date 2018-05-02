// Package anybus provides ...
package anybus

import (
	"errors"

	"github.com/gilwil/anybus/driver"
)

// AnyBus is the primary handle to a bus
type AnyBus struct {
	d         driver.Driver
	namespace string
	name      string
}

// Config struct
type Config struct {
	// Driver must be fully initialised
	Driver driver.Driver

	// Namespace may be empty string
	Namespace string

	// Name must uniquely identify the application within the namespace
	Name string
}

// Create instantiates a new AnyBus connection based on the supplied driver interface
func Create(cfg Config) (*AnyBus, error) {

	switch {
	case cfg.Driver == nil:
		return nil, errors.New("nil driver in config")
	case cfg.Name == "":
		return nil, errors.New("name is required to be non-nil")
	}

	bus := &AnyBus{
		d:         cfg.Driver,
		namespace: cfg.Namespace,
		name:      cfg.Name,
	}

	return bus, nil
}
