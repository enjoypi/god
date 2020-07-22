package net

type Config struct {
	Net NetConfig
}

type NetConfig struct {
	ListenAddress string
}