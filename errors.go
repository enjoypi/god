package god

import "errors"

var (
	ErrInvalidNodeID        = errors.New("invalid node ID")
	ErrNoService            = errors.New("the service is not exists")
	ErrDuplicateApplication     = errors.New("the service is already exists")
	ErrDuplicateActor       = errors.New("the actor is already exists")
	ErrFailedInitialization = errors.New("failed initialization")
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
