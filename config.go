package god

type Config struct {
	Node NodeConfig
	Log  Logging
	Apps []string
}

type NodeConfig struct {
	Type string
	ID   uint32
}

type Logging struct {
	Level string
}
