package transports

import (
	"context"
	"io"
)

// Encapsulates the I/O layer
type Transport interface {
	io.ReadWriteCloser
	Flush(ctx context.Context) (err error)
	RemainingBytes() (num_bytes uint64)

	// Opens the transport for communication
	Open() error

	// Returns true if the transport is open
	IsOpen() bool
}

type constructor func() Transport

var (
	constructors = make(map[string]constructor)
)

func NewTransport(transportType string) Transport {
	if c, ok := constructors[transportType]; ok {
		return c()
	}
	return nil
}
