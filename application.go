package god

import "github.com/spf13/viper"

type Application interface {
	//ChangeConfig(Changed, New, Removed) error
	PrepareStop()
	Name() string
	Start(v *viper.Viper) error
	Stop()
}

type NewApplication func(v *viper.Viper) (Application, error)

var (
	applicationFactory map[string]NewApplication
	applications       []Application
)

func initializeApplications(v *viper.Viper, apps []string) error {
	defer func() {
		// TODO: destroy on error
		for _, app := range applications {
			app.PrepareStop()
		}
	}()

	for _, name := range apps {
		if err := startApplication(v, name); err != nil {
			return err
		}
	}
	return nil
}

func startApplication(v *viper.Viper, name string) error {
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

func RegisterApplication(name string, creator NewApplication) {

}
