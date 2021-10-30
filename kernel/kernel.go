package kernel

import (
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"
)

const name = "kernel"

func init() {
	RegisterApplication(name, newKernel)
}

type kernel struct {
}

func newKernel(v *viper.Viper) (types.Application, error) {
	return &kernel{}, nil
}

func (k *kernel) PrepareStop() {

}

func (k *kernel) Name() string {
	return name
}

func (k *kernel) Start(v *viper.Viper) error {

	return nil
}

func (k *kernel) Stop() {

}

func Start(v *viper.Viper) error {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := startApplication(v, "kernel"); err != nil {
		return err
	}

	if err := initializeApplications(v, cfg.Apps); err != nil {

	}
	return nil
}
