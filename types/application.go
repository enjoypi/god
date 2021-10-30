package types

import "github.com/spf13/viper"

type ID int64

type Application interface {
	//ChangeConfig(Changed, New, Removed) error
	PrepareStop()
	Name() string
	Start(v *viper.Viper) error
	Stop()
}
type NewApplication func(v *viper.Viper) (Application, error)
