package god

import (
	"ext"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
)

const (
	adminExchange = "god.admin"
)

type node struct {
	*amqp.Connection
	*Session

	kind uint16
	id   uint64
}

var self node

func Start(url string, nodeType uint16, nodeID uint64) error {
	c, err := amqp.Dial(url)
	if err != nil {
		return err
	}

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
	self.kind = nodeType
	self.id = nodeID

	var req AuthReq
	req.ID = nodeID
	postAdmin("Auth", &req)

	self.register(&_Node_serviceDesc, &self)
	go self.Handle(q, nil)
	return nil
}

func Close() {
	self.Close()
}

func postAdmin(method string, msg proto.Message) error {
	return self.Post(adminExchange,
		self.kind, self.id,
		"god.Node", method, msg)
}

func (n *node) Auth(c context.Context, req *AuthReq) (*AuthAck, error) {
	ext.LogDebug("%#v", req)
	return &AuthAck{Code: ErrorCode_OK}, nil
}
