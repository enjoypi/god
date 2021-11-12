package stdlib

import (
	"github.com/enjoypi/god/def"
	"github.com/spf13/viper"
)

var (
	applicationFactory map[string]def.NewApplication
	applications       []def.Application
)

func init() {
	applicationFactory = make(map[string]def.NewApplication)
	applications = make([]def.Application, 0)
}

func StartApplications(v *viper.Viper, apps []string) error {
	defer func() {
		// TODO: destroy on error
		for _, app := range applications {
			app.PrepareStop()
		}
	}()

	for _, name := range apps {
		if err := StartApplication(v, name); err != nil {
			return err
		}
	}
	return nil
}

func StartApplication(v *viper.Viper, name string) error {
	creator, ok := applicationFactory[name]
	if !ok {
		return ErrNoSuchApplication
	}

	app, err := creator(v)
	if err != nil {
		return err
	}

	if err := app.Start(v); err != nil {
		return err
	}

	applications = append(applications, app)
	return nil
}

func RegisterApplication(name string, creator def.NewApplication) {
	applicationFactory[name] = creator
}
