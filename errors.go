package common

import (
	"errors"
)

var (
	// ErrInvalidAddressFormat TODO
	ErrInvalidAddressFormat = errors.New("invalid address format")
	// ErrInvalidTagFormat TODO
	ErrInvalidTagFormat = errors.New("invalid tag format")
	// ErrInvalidSignature TODO
	ErrInvalidSignature = errors.New("invalid signature")
)
