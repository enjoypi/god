package god

type Config struct {
	Node NodeConfig
}

type NodeConfig struct {
	Type string
	ID   uint16
}
