package common

import (
	"errors"
)

// common errors
var (
	ErrInvalidAddressFormat = errors.New("invalid address format")
	ErrInvalidTagFormat     = errors.New("invalid tag format")
	ErrInvalidSignature     = errors.New("invalid signature")
)
