package event

import "net"

type EvSocketConnected struct {
	net.Conn
}

type EvTCPConnected struct {
	*net.TCPConn
}
