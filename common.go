package god

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"ext"
	"io"
)

func WriteBytes(w io.Writer, data []byte) {
	ext.ANoError(binary.Write(w, DEFAULT_BYTE_ORDER, uint16(len(data))))
	ext.ANoError(binary.Write(w, DEFAULT_BYTE_ORDER, data))
	ext.LogDebug("WRITTEN\t%d", len(data))
}

func ReadBytes(r io.Reader) []byte {
	var size uint16
	ext.ANoError(binary.Read(r, DEFAULT_BYTE_ORDER, &size))
	data := make([]byte, size)
	ext.ANoError(binary.Read(r, DEFAULT_BYTE_ORDER, data))
	ext.LogDebug("READ\t%d", size)
	return data
}

func GobDecode(bin []byte, data interface{}) {
	buf := bytes.NewBuffer(bin)
	decoder := gob.NewDecoder(buf)
	ext.ANoError(decoder.Decode(data))
}

func GobEncode(data interface{}) []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	ext.ANoError(encoder.Encode(data))
	return buf.Bytes()
}

func DefaultDecode(b []byte) Message {
	var m Message
	GobDecode(b, &m)
	return m
}

func DefaultEncode(m Message) []byte {
	return GobEncode(m)
}

func DefaultCompress(in []byte) []byte {
	return in
}

func DefaultDecompress(in []byte) []byte {
	return in
}
