package god

import (
	"github.com/enjoypi/god/kernel"
	_ "github.com/enjoypi/god/kernel"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

func Initialize(v *viper.Viper) error {

	if err := stdlib.Initialize(v); err != nil {
		return err
	}

	if err := kernel.Start(v); err != nil {
		return err
	}
	return nil
}

func Wait() {
	stdlib.Wait()
}
