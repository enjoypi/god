package god

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

const defaultTimeout = 5 * time.Second

func StartNode(cfg *Config, logger *zap.Logger) error {
	ecfg := cfg.Etcd
	if ecfg.DialTimeout == 0 {
		ecfg.DialTimeout = defaultTimeout
	}

	cli, err := clientv3.New(ecfg)
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	resp, err := cli.Get(ctx, cfg.NodePath)
	cancel()
	if err != nil {
		return err
	}
	// use the response
	resp.OpResponse()
	return nil
}
