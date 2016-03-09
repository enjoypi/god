package god

import (
	"encoding/binary"
)

var (
	DEFAULT_BYTE_ORDER = binary.LittleEndian
)

type Compress func([]byte) []byte
type Decompress func([]byte) []byte

type Encrypt func([]byte) []byte
type Decrypt func([]byte) []byte
