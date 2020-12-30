module github.com/enjoypi/god

go 1.15

replace github.com/enjoypi/gostatechart => ../gostatechart

require (
	github.com/enjoypi/gostatechart v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3
	github.com/nats-io/nats.go v1.10.0
	github.com/stretchr/testify v1.6.0
	go.etcd.io/etcd v3.4.14+incompatible
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
	google.golang.org/grpc v1.29.1
)
