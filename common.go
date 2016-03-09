package god

import (
	"encoding/binary"
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

func DefaultCompress(in []byte) []byte {
	return in
}

func DefaultDecompress(in []byte) []byte {
	return in
}
