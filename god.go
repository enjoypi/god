package god

import (
	_ "github.com/enjoypi/god/actors"
	_ "github.com/enjoypi/god/kernel"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

func Initialize(v *viper.Viper) error {

	if err := stdlib.Initialize(v); err != nil {
		return err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := stdlib.StartApplication(v, "kernel"); err != nil {
		return err
	}

	return stdlib.StartApplications(v, cfg.Apps)
}

func Wait() {
	stdlib.Wait()
}
