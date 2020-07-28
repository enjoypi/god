module github.com/enjoypi/god

go 1.14

//replace github.com/coreos/etcd => go.etcd.io/etcd v3.4.9+incompatible

//replace go.etcd.io/etcd => github.com/coreos/etcd v3.4.9+incompatible

//replace github.com/coreos/bbolt => go.etcd.io/etcd/bolt v1.3.5

replace github.com/enjoypi/gostatechart => ../gostatechart

require (
	github.com/enjoypi/gostatechart v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.2
	github.com/nats-io/nats.go v1.10.0
	go.etcd.io/etcd v3.4.9+incompatible
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.29.0
)
