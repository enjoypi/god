package god

import (
	"time"
)

const defaultTimeout = 60 * time.Second

type Config struct {
	ListenAddress string
	NodePath      string
	EtcdTTL       int64

	//Etcd     etcdclient.Config
	Node NodeConfig
}

type NodeConfig struct {
	Type string
	ID   uint16
}
