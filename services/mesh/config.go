package mesh

import (
	"time"

	"google.golang.org/grpc"
	etcdclient "go.etcd.io/etcd/clientv3"
)

type Config struct {
	Etcd etcdclient.Config
	Mesh MeshConfig
}

type MeshConfig struct {
	DefaultTimeout time.Duration // 未设置Timeout时默认超时
	RetryTimes     int           // 未设置Timeout时默认超时
	GrantTTL       int64
	Path           string
}

func normalizeConfig(config Config) Config {
	if config.Etcd.DialTimeout == 0 {
		config.Etcd.DialTimeout = config.Mesh.DefaultTimeout
	}

	config.Etcd.DialOptions = append(config.Etcd.DialOptions, grpc.WithBlock())
	return config
}
