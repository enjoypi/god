package t_nats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction()
	cfg       Config
	trans     *Transport
)

func init() {
	cfg.Nats.Url = "mypc:4222" //nats.DefaultURL
	trans = NewTransport(cfg, logger, 1)
}

func TestNewTransport(t *testing.T) {
	sub, err := trans.SubscribeSync("test")
	require.NoError(t, err)
	require.NotNil(t, sub)

	inbox := trans.NewRespInbox()
	require.NoError(t, trans.PublishRequest("test", inbox, []byte("data")))
	for msg, err := sub.NextMsg(time.Second); err == nil; msg, err = sub.NextMsg(time.Second) {
		require.Equal(t, "test", msg.Subject)
		require.Equal(t, "data", string(msg.Data))
		t.Log(msg.Reply)
		require.NoError(t, trans.Publish(msg.Reply, nil))
	}

	require.NoError(t, sub.Unsubscribe())
}

func BenchmarkNewTransport(b *testing.B) {
	//wg := sync.WaitGroup{}

	subject := "test"
	//recv := 0
	//sub, err := trans.Subscribe(subject, func(msg *nats.Msg) {
	//	recv++
	//	//_ = trans.Publish(msg.Reply, nil)
	//	//wg.Done()
	//})
	//require.NoError(b, err)
	//require.NotNil(b, sub)

	inbox := trans.NewRespInbox()
	data := []byte("data")
	sent := 0
	for i := 0; i < b.N; i++ {
		_ = trans.PublishRequest(subject, inbox, data)
		//wg.Add(1)
		sent++
	}
	//wg.Wait()
	//require.Equal(b, sent, recv)

	//require.NoError(b, sub.Unsubscribe())
}
