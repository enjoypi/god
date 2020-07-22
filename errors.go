package god

import "errors"

var (
	ErrInvalidNodeID        = errors.New("invalid node ID")
	ErrDuplicateService     = errors.New("the service is already exists")
	ErrDuplicateActor       = errors.New("the actor is already exists")
	ErrFailedInitialization = errors.New("failed initialization")
)
