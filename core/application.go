package core

type Application interface {
	//ChangeConfig(Changed, New, Removed)
	Name() string
	Start() error
	Stop()
}

var (
	applications = NewManager()
)

func StartApplication(app Application) error {
	//_, ok := services.Load(app.Name())
	//if ok {
	//	return god.ErrDuplicateApplication
	//}
	//
	//if err := app.Start(); err != nil {
	//	return err
	//}
	//services.Store(service.Name(), service)
	return nil
}

func Serve() error {
	Wait()
	return nil
}
