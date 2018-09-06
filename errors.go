package common

import (
	"errors"
)

var (
	// ErrInvalidAddressFormat TODO
	ErrInvalidAddressFormat = errors.New("invalid address format")
	// ErrInvalidSignature TODO
	ErrInvalidSignature = errors.New("invalid signature")
)
