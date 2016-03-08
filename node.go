package god

import "github.com/streadway/amqp"

var (
	conn     *amqp.Connection = nil
	nodeType uint16           = 0
	nodeID   uint64           = 0
)

func Start(url string, nodeType uint16, nodeID uint64) error {
	c, err := amqp.Dial(url)
	if err == nil {
		conn = c
		s, err := NewSession()
		if err != nil {
			s.Close()
			return err
		}

		q, err := s.Subscribe("nodes", nodeType, nodeID)
		if err != nil {
			s.Close()
			return err
		}
		go handleNodeMsg(s, q)
	}
	return err
}

func Close() {
	conn.Close()
}

func handleNodeMsg(session *Session, queue string) {
	msgs, err := session.Pull(queue)
	if err != nil {
		return
	}
	for d := range msgs {
		d.Ack(false)
	}
}
