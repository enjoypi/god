package mesh

import (
	"context"

	etcdclient "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

func dialEtcd(cfg Config, logger *zap.Logger) (*etcdclient.Client, error) {

	mcfg := cfg.Mesh
	ecfg := cfg.Etcd
	if ecfg.DialTimeout == 0 {
		ecfg.DialTimeout = mcfg.DefaultTimeout
	}

	cli, err := etcdclient.New(ecfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), mcfg.DefaultTimeout)
	if resp, err := cli.Grant(ctx, mcfg.GrantTTL); err != nil {
		err = cli.Close()
		return nil, err
	} else {
		logger.Info("lease id: ", zap.Any("lease id", int64(resp.ID)), zap.Any("endpoints", cli.Endpoints()))
	}
	cancel()

	return cli, err
}
