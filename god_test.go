package god

import (
	"ext"
	"fmt"
	"testing"
	"time"
)

func TestStartNode(t *testing.T) {
	ext.AssertNoError(t,
		Start("amqp://guest:guest@localhost:5672/", 0, ext.RandomUint64()),
		"start god")

	var producer, consumer *Session
	var err error

	consumer, err = NewSession()
	ext.AssertNoError(t, err, "new consumer")

	exchange := "logs"
	routingKeyType := ext.RandomUint16()
	routingKey := ext.RandomUint64()
	count := 10

	q, err := consumer.Subscribe(exchange, routingKeyType, routingKey)
	ext.AssertNoError(t, err, "pull msgs")
	go func() {
		msgs, err := consumer.Pull(q)
		ext.AssertNoError(t, err, "pull msgs")
		i := 0
		for d := range msgs {
			ext.CheckEqual(t, string(d.Body), fmt.Sprintf("hello world %d", i))
			d.Ack(false)
			i++
		}
		ext.CheckEqual(t, i, count)
	}()

	producer, err = NewSession()
	ext.AssertNoError(t, err, "new producer")
	err = producer.Declare(exchange)
	ext.AssertNoError(t, err, "declare exchange")
	for i := 0; i < count; i++ {
		err = producer.Post(exchange,
			routingKeyType, routingKey,
			[]byte(fmt.Sprintf("hello world %d", i)))
		ext.AssertNoError(t, err, "post")
	}

	time.Sleep(time.Millisecond * 100)
}

func BenchmarkStartNode(b *testing.B) {
}
