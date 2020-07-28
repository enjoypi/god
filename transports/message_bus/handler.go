package message_bus

//import (
//	sc "github.com/enjoypi/gostatechart"
//	"github.com/nats-io/nats.go"
//	"go.uber.org/zap"
//)
//
//type Handler struct {
//	sc.SimpleState
//	*Service
//
//	producer *nats.Conn
//	replyTo  string
//	sub      *nats.Subscription
//}
//
//func (h *Handler) Begin(context interface{}, event sc.Event) sc.Event {
//	h.Service = context.(*Service)
//
//	h.RegisterReaction((*nats.Msg)(nil), h.onNatsMsg)
//	conn, _ := nats.Connect("nats://pchost:4222")
//	h.producer = conn
//	h.replyTo = conn.NewRespInbox()
//	h.sub, _ = conn.Subscribe(">", h.onReply)
//	return h.producer.Publish("hello", []byte("world"))
//}
//
//func (h *Handler) onNatsMsg(event sc.Event) sc.Event {
//	msg := event.(*nats.Msg)
//	h.Logger.Debug("onNatsMsg", zap.ByteString("msg", msg.Data))
//	return msg.Respond(msg.Data)
//	//return nil
//}
//
//func (h *Handler) onReply(msg *nats.Msg) {
//	h.Logger.Debug("onNatsMsg", zap.ByteString("msg", msg.Data))
//	h.producer.Publish("hello", []byte("world"))
//}
