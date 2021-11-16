package mesh

import (
	"errors"
)

var (
	ErrEtcdDropped = errors.New("the connection of etcd is dropped")
)
