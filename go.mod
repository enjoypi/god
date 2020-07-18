module github.com/enjoypi/god

go 1.14

//replace github.com/coreos/etcd => go.etcd.io/etcd v3.3.22+incompatible
//replace go.etcd.io/etcd => github.com/coreos/etcd v3.3.22+incompatible

require (
	github.com/golang/protobuf v1.4.2
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.30.0
)
