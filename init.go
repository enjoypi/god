package god

import "github.com/spf13/viper"

func Initialize(v *viper.Viper) error {

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	if err := initializeLogger(cfg); err != nil {
		return err
	}

	if err := startApplication(v, "core"); err != nil {
		return err
	}

	if err := initializeApplications(v, cfg.Apps); err != nil {

	}
	return nil
}
