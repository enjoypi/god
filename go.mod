module github.com/enjoypi/god

go 1.14

replace go.etcd.io/etcd => github.com/coreos/etcd v3.3.22+incompatible

replace github.com/coreos/etcd => go.etcd.io/etcd v3.3.22+incompatible

require (
	github.com/coreos/etcd v0.0.0-00010101000000-000000000000 // indirect
	go.etcd.io/etcd v3.3.22+incompatible
	go.uber.org/zap v1.15.0
)
