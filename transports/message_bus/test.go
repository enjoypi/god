package message_bus

import (
	"strconv"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type test struct {
	*zap.Logger
	producer *nats.Conn
	replyTo  string
	sub      *nats.Subscription
	serial   int
}

func newTest(logger *zap.Logger) *test {
	t := &test{Logger: logger}
	opts := nats.GetDefaultOptions()
	opts.Url = "nats://pchost:4222"
	opts.NoEcho = true
	t.producer, _ = opts.Connect()
	t.sub, _ = t.producer.Subscribe(">", t.onMsg)
	t.serial++
	t.producer.Publish("hello", []byte(strconv.Itoa(t.serial)))
	return t
}

func (t *test) onMsg(msg *nats.Msg) {
	t.Logger.Debug("onMsg", zap.String("subject", msg.Subject), zap.Binary("data", msg.Data))
	t.serial++
	t.producer.PublishMsg(msg)
}
