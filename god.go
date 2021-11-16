package god

import (
	_ "github.com/enjoypi/god/actor"
	_ "github.com/enjoypi/god/application"
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/option"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

func Initialize(v *viper.Viper) error {
	if err := logger.Initialize(v); err != nil {
		return err
	}

	var opt struct {
		option.Node
	}
	if err := v.Unmarshal(&opt); err != nil {
		return err
	}

	return stdlib.StartApplications(v, opt.Apps)
}

func Wait() {
	stdlib.Wait()
}
