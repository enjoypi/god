package stdlib

import "github.com/spf13/viper"

func Initialize(v *viper.Viper) error {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := initializeLogger(cfg); err != nil {
		return err
	}

	return nil
}
