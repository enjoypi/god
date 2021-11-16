package event

import "net"

type EvSocketConnected struct {
	net.Conn
}
