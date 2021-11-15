package events

import "net"

type EvNetConnected struct {
	net.Conn
}
