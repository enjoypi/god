package god

import (
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

const (
	adminExchange = "god.admin"
)

type node struct {
	*amqp.Connection
	*Session

	kind uint16
	ID   uint64
}

var self node

func Start(url string, nodeType uint16, nodeID uint64) error {
	c, err := amqp.Dial(url)
	if err == nil {
		self.Connection = c
		s, err := NewSession()
		if err != nil {
			s.Close()
			return err
		}

		q, err := s.Subscribe(adminExchange, nodeType, nodeID)
		if err != nil {
			s.Close()
			return err
		}

		self.Session = s
		go self.Handle(q, handleAdmin)
	}
	return err
}

func Close() {
	self.Close()
}

func postAdmin(msg proto.Message) error {
	return self.Post(adminExchange,
		self.kind, self.ID,
		msg)
}

func handleAdmin(*amqp.Delivery) error {
	return nil
}
