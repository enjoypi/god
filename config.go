package god

type Config struct {
	Node NodeConfig
	Apps []string
}

type NodeConfig struct {
	Type string
	ID   uint32
}
