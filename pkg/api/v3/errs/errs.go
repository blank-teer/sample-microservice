package errs

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrNotValid = errors.New("invalid")
)
