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
	GrantTTL	 int64
	DefaultTimeout time.Duration
	Path	string
}

//const defaultTimeout = 5 * time.Second
