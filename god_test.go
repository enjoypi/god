package god

import (
	"ext"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
)

func TestGod(t *testing.T) {
	ext.AssertNoError(t,
		Start("amqp://guest:guest@localhost:5672/", 0, ext.RandomUint64()),
		"start god")

	var producer, consumer *Session
	var err error

	consumer, err = NewSession()
	ext.AssertNoError(t, err, "new consumer")

	exchange := "god.test"
	routingKeyType := ext.RandomUint16()
	routingKey := ext.RandomUint64()

	q, err := consumer.Subscribe(exchange, routingKeyType, routingKey)
	ext.AssertNoError(t, err, "pull msgs")
	ci := int64(0)
	go consumer.Handle(q,
		func(method string, msg proto.Message) error {
			ext.CheckEqual(t, method, "Test")
			test := msg.(*Test)
			ext.CheckEqual(t, test.Count, ci)
			ci++
			return nil
		})

	producer, err = NewSession()
	ext.AssertNoError(t, err, "new producer")

	for i := int64(0); i < 1000; i++ {
		var test Test
		test.Count = i
		err = producer.Post(exchange,
			routingKeyType, routingKey,
			"Test", &test)
		ext.AssertNoError(t, err, "post")
	}

	time.Sleep(time.Millisecond * 100)
}

func BenchmarkStartNode(b *testing.B) {
}
