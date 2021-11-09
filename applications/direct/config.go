package direct

type Config struct {
	Net NetConfig
}

type NetConfig struct {
	ListenAddress string
}
