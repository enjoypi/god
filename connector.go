package god

import (
	"ext"
	"net"
)

type connector struct {
	NewAgent
}

func NewConnector(newAgent NewAgent) Connector {
	return &connector{NewAgent: newAgent}
}

func (c *connector) Dial(address string) {
	conn, err := net.Dial("tcp", address)
	ext.ANoError(err)
	c.NewAgent(conn)
}
