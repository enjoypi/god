package god

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Session struct {
	*amqp.Channel
}

func NewSession() (*Session, error) {
	ch, err := self.Connection.Channel()
	if err != nil {
		return nil, err
	}

	var s Session
	s.Channel = ch
	return &s, nil
}

func combine(routingKeyType uint16, routingKey uint64) string {
	return fmt.Sprintf("%d.%d", routingKeyType, routingKey)
}

func (s *Session) Declare(exchange string) error {
	return s.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
}

func (s *Session) Post(exchange string,
	routingKeyType uint16, routingKey uint64,
	msg []byte) error {
	return s.post(exchange, combine(routingKeyType, routingKey), msg)
}

func (s *Session) post(exchange string, routingKey string, msg []byte) error {
	return s.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         msg,
		})
}

func (s *Session) Subscribe(exchange string,
	routingKeyType uint16, routingKey uint64) (string, error) {
	err := s.Declare(exchange)
	if err != nil {
		return "", err
	}

	q, err := s.declareQueue()
	if err != nil {
		return "", err
	}

	return q.Name,
		s.bind(exchange,
			q.Name,
			combine(routingKeyType, routingKey))
}

func (s *Session) declareQueue() (amqp.Queue, error) {
	return s.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func (s *Session) bind(exchange string, queue string, routingKey string) error {
	return s.QueueBind(
		queue,      // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil)
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

type Handler func(delivery *amqp.Delivery) error

func (s *Session) Handle(queue string, handler Handler) error {
	msgs, err := s.Pull(queue)
	if err != nil {
		return err
	}

	for d := range msgs {
		err := handler(&d)
		if err != nil {
			return err
		}
		d.Ack(false)
	}
	return nil
}
