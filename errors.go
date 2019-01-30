package common

import (
	"errors"
)

// common errors
var (
	ErrInvalidAddressFormat   = errors.New("invalid address format")
	ErrInvalidAddressCheckSum = errors.New("invalid address checksum")
	ErrInvalidTagFormat       = errors.New("invalid tag format")
	ErrInvalidSignature       = errors.New("invalid signature")
	ErrInvalidPublicHash      = errors.New("invalid public hash")
)
