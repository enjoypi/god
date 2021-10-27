package mesh

import (
	"context"

	etcdclient "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

func dialEtcd(cfg Config, logger *zap.Logger) (*etcdclient.Client, etcdclient.LeaseID, error) {

	meshCfg := cfg.Mesh
	etcdCfg := cfg.Etcd

	cli, err := etcdclient.New(etcdCfg)
	if err != nil {
		return nil, 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), meshCfg.DefaultTimeout)
	resp, err := cli.Grant(ctx, meshCfg.GrantTTL)
	if err != nil {
		_ = cli.Close()
		return nil, 0, err
	} else {
	}
	cancel()

	logger.Info("lease id: ", zap.Any("lease id", int64(resp.ID)), zap.Any("endpoints", cli.Endpoints()))
	return cli, resp.ID, err
}
