package mesh

import (
	"time"

	"github.com/enjoypi/god/services/net"
	etcdclient "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type Config struct {
	Mesh MeshConfig
	Etcd etcdclient.Config
}

type MeshConfig struct {
	Net net.Config

	AdvertiseAddress string
	DefaultTimeout   time.Duration // 未设置Timeout时默认超时
	GrantTTL         int64
	Path             string
	RetryTimes       int // 尝试连接次数
}

func normalizeConfig(config Config) Config {
	if config.Etcd.DialTimeout == 0 {
		config.Etcd.DialTimeout = config.Mesh.DefaultTimeout
	}

	config.Etcd.DialOptions = append(config.Etcd.DialOptions, grpc.WithBlock())
	return config
}
