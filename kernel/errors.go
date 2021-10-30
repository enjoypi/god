package kernel

import "errors"

var (
	ErrNoSuchApplication = errors.New("no such application")
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
