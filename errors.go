package god

import "errors"

var (
	ErrInvalidNodeID    = errors.New("invalid node ID")
	ErrDuplicateService = errors.New("the service is already exists")
)
