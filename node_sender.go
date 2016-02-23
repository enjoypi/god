package god

import (
	"bytes"
	"encoding/gob"
	"ext"
	"net"
	"time"
)

type nodeSender struct {
	net.Conn
	Encode
	Compress

	*stopper
	MessageQueue
}

func NewNodeSender(conn net.Conn, encode Encode, compress Compress) MessageQueue {
	s := &nodeSender{Conn: conn, Encode: encode, Compress: compress, stopper: NewStopper(), MessageQueue: NewMessageQueue(32)}
	go ext.PCall(s.Run)
	return s
}

func (s *nodeSender) Run() {
	defer s.Stopped()

	if RELEASE {
		ext.ANoError(s.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second)))
	}
	WriteBytes(s.Conn, GobEncode(MyInfo()))

	for !s.StopRequested() {
		source, target, msg := s.Pop()
		data := s.Encode(msg)
		ext.ATrue(data != nil)

		var buf bytes.Buffer

		enc := gob.NewEncoder(&buf)
		ext.ANoError(enc.Encode(source))
		ext.ANoError(enc.Encode(target))
		ext.ANoError(enc.Encode(len(data)))
		ext.ANoError(enc.Encode(data))

		data = buf.Bytes()
		if s.Compress != nil {
			data = s.Compress(buf.Bytes())
		}

		if RELEASE {
			ext.ANoError(s.Conn.SetWriteDeadline(time.Now().Add(time.Minute)))
		}
		WriteBytes(s.Conn, data)
	}
}
