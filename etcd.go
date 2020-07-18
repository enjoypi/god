package god

import (
	"go.uber.org/zap"
)

func initEtcd(cfg *Config, logger *zap.Logger) error {

	//ecfg := cfg.Etcd
	//if ecfg.DialTimeout == 0 {
	//	ecfg.DialTimeout = defaultTimeout
	//}
	//
	//cli, err := etcdclient.New(ecfg)
	//if err != nil {
	//	return err
	//}
	//defer cli.Close()
	//
	////ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	////_, err = cli.Grant(ctx, cfg.EtcdTTL)
	////if err != nil {
	////	return err
	////}
	//
	////ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	//resp, err := cli.Get(context.TODO(), cfg.NodePath)
	////cancel()
	//if err != nil {
	//	return err
	//}
	//// use the response
	//resp.OpResponse()
	return nil
}
