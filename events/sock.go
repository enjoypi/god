package events

import "net"

type EvSocketConnected struct {
	net.Conn
}
