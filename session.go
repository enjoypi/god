package god

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Session struct {
	*amqp.Channel
	exchanges map[string]bool
}

func NewSession() (*Session, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	var s Session
	s.Channel = ch
	s.exchanges = make(map[string]bool)
	return &s, nil
}

func combine(routingKeyType uint16, routingKey uint64) string {
	return fmt.Sprintf("%d-%d", routingKeyType, routingKey)
}

func (s *Session) Post(exchange string,
	routingKeyType uint16, routingKey uint64,
	msgID uint64, msg []byte) error {

	if !s.exchanges[exchange] {
		err := s.ExchangeDeclare(
			exchange, // name
			"direct", // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		if err != nil {
			return err
		}
		s.exchanges[exchange] = true
	}

	err := s.Publish(
		exchange, // exchange
		combine(routingKeyType, routingKey), // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	return err
}

func (s *Session) Subscribe(exchange string,
	routingKeyType uint16, routingKey uint64) (string, error) {
	err := s.ExchangeDeclare(
		exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return "", err
	}

	q, err := s.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return "", err
	}

	err = s.QueueBind(
		q.Name, // queue name
		combine(routingKeyType, routingKey), // routing key
		exchange, // exchange
		false,
		nil)
	if err != nil {
		return "", err
	}

	return q.Name, nil

}

func (s *Session) Pull(queue string) (<-chan amqp.Delivery, error) {
	return s.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}
