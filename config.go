package god

import "go.etcd.io/etcd/clientv3"

type Config struct {
	NodePath string
	Etcd     clientv3.Config
}
