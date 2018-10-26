package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/util"
)

// PublicKeySize is 33 bytes
const PublicKeySize = 33

// PublicKey is the [PublicKeySize]byte with methods
type PublicKey [PublicKeySize]byte

// WriteTo is a serialization function
func (pubkey *PublicKey) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(pubkey[:]); err != nil {
		return int64(n), err
	} else if n != PublicKeySize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom is a deserialization function
func (pubkey *PublicKey) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(pubkey[:]); err != nil {
		return int64(n), err
	} else if n != PublicKeySize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal checks compare two values and returns true or false
func (pubkey PublicKey) Equal(b PublicKey) bool {
	return bytes.Equal(pubkey[:], b[:])
}

// String returns the hex string of the public key
func (pubkey PublicKey) String() string {
	return hex.EncodeToString(pubkey[:])
}

// Clone returns the clonend value of it
func (pubkey PublicKey) Clone() PublicKey {
	var cp PublicKey
	copy(cp[:], pubkey[:])
	return cp
}

// Checksum returns the checksum byte
func (pubkey PublicKey) Checksum() byte {
	var cs byte
	for i := 0; i < len(pubkey); i++ {
		cs = cs ^ pubkey[i]
	}
	return cs
}
