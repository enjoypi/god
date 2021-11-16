package god

import (
	_ "github.com/enjoypi/god/actors/implement"
	_ "github.com/enjoypi/god/applications"
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/options"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

func Initialize(v *viper.Viper) error {
	if err := logger.Initialize(v); err != nil {
		return err
	}

	var opt struct {
		options.Node
	}
	if err := v.Unmarshal(&opt); err != nil {
		return err
	}

	return stdlib.StartApplications(v, opt.Apps)
}

func Wait() {
	stdlib.Wait()
}
