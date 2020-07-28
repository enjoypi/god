package message_bus

import (
	"github.com/nats-io/nats.go"
)

type Config struct {
	Nats nats.Options
}
