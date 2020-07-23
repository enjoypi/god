package mesh

import (
	"time"

	etcdclient "go.etcd.io/etcd/clientv3"
)

type Config struct {
	Etcd etcdclient.Config
	Mesh MeshConfig
}

type MeshConfig struct {
	DefaultTimeout time.Duration
	GrantTTL       int64
	Path           string
}

//const defaultTimeout = 5 * time.Second
