package mesh

import (
	"context"
	etcdclient "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)


func initEtcd(cfg Config, logger *zap.Logger) error {

	mcfg := cfg.Mesh
	ecfg := cfg.Etcd
	if ecfg.DialTimeout == 0 {
		ecfg.DialTimeout = mcfg.DefaultTimeout
	}

	cli, err := etcdclient.New(ecfg)
	if err != nil {
		return err
	}
	defer cli.Close()

	//ctx, cancel := context.WithTimeout(context.Background(), mcfg.DefaultTimeout)
	_, err = cli.Grant(context.TODO(), mcfg.GrantTTL)
	if err != nil {
		return err
	}

	//ctx, cancel := context.WithTimeout(context.Background(), mcfg.DefaultTimeout)
	resp, err := cli.Get(context.TODO(), mcfg.Path)
	//cancel()
	if err != nil {
		return err
	}
	resp.OpResponse()
	return nil
}
