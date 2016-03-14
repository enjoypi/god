package god

import (
	"ext"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const (
	adminExchange = "god.admin"
)

// service consists of the information of the server serving this service and
// the methods in this service.
type service struct {
	server interface{} // the server for service methods
	md     map[string]*grpc.MethodDesc
}

type node struct {
	*amqp.Connection
	*Session

	kind uint16
	id   uint64
	m    map[string]*service // service name -> service info
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
	self.m = make(map[string]*service)

	var req AuthReq
	req.ID = nodeID
	postAdmin("Auth", &req)

	self.register(&_Node_serviceDesc, newServer())
	go self.Handle(q, self.dispatch)
	return nil
}

func Close() {
	self.Close()
}

func (s *node) register(sd *grpc.ServiceDesc, ss interface{}) {
	if _, ok := s.m[sd.ServiceName]; ok {
	}
	srv := &service{
		server: ss,
		md:     make(map[string]*grpc.MethodDesc),
	}
	for i := range sd.Methods {
		d := &sd.Methods[i]
		srv.md[d.MethodName] = d
	}
	s.m[sd.ServiceName] = srv
}

func postAdmin(method string, msg proto.Message) error {
	return self.Post(adminExchange,
		self.kind, self.id,
		method, msg)
}

func (n *node) dispatch(method string, msg proto.Message) error {
	srv := n.m[_Node_serviceDesc.ServiceName]
	if srv == nil {
		return nil
	}

	md := srv.md[method]
	if md == nil {
		return nil
	}

	out, err := md.Handler(srv.server, nil, func(interface{}) error { return nil })
	ext.LogDebug("%#v\t%#v", msg, out)
	return err
}
