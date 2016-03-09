package god

import (
	"ext"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
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
	count := int64(10)

	q, err := consumer.Subscribe(exchange, routingKeyType, routingKey)
	ext.AssertNoError(t, err, "pull msgs")
	i := int64(0)
	go consumer.Handle(q,
		func(d *amqp.Delivery) error {
			var test Test
			err := proto.Unmarshal(d.Body, &test)
			ext.AssertNoError(t, err, "post")
			ext.CheckEqual(t, test.Count, i)
			i++
			return nil
		})

	producer, err = NewSession()
	ext.AssertNoError(t, err, "new producer")
	err = producer.Declare(exchange)
	ext.AssertNoError(t, err, "declare exchange")
	for i := int64(0); i < count; i++ {
		var test Test
		test.Count = i
		err = producer.Post(exchange,
			routingKeyType, routingKey,
			&test)
		ext.AssertNoError(t, err, "post")
	}

	time.Sleep(time.Millisecond * 100)
}

func BenchmarkStartNode(b *testing.B) {
}
