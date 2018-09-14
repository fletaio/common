package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/util"
)

// PublicKeySize TODO
const PublicKeySize = 33

// PublicKey TODO
type PublicKey [PublicKeySize]byte

// WriteTo TODO
func (pubkey *PublicKey) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(pubkey[:]); err != nil {
		return int64(n), err
	} else if n != PublicKeySize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (pubkey *PublicKey) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(pubkey[:]); err != nil {
		return int64(n), err
	} else if n != PublicKeySize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal TODO
func (pubkey PublicKey) Equal(b PublicKey) bool {
	return bytes.Equal(pubkey[:], b[:])
}

// String TODO
func (pubkey PublicKey) String() string {
	return hex.EncodeToString(pubkey[:])
}
