package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/hash"
	"git.fleta.io/fleta/common/util"
)

// PublicHashSize is 31 bytes
const PublicHashSize = 31

// PublicHash is the [PublicHashSize]byte with methods
type PublicHash [PublicHashSize]byte

// NewPublicHash returns the PublicHash of the pubkey
func NewPublicHash(pubkey PublicKey) PublicHash {
	h := hash.DoubleHash(pubkey[:])
	var ph PublicHash
	ph[0] = pubkey.Checksum()
	copy(ph[1:], h[:len(h)-2])
	return ph
}

// WriteTo is a serialization function
func (pubhash *PublicHash) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(pubhash[:]); err != nil {
		return int64(n), err
	} else if n != PublicHashSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom is a deserialization function
func (pubhash *PublicHash) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(pubhash[:]); err != nil {
		return int64(n), err
	} else if n != PublicHashSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal checks that two values is same or not
func (pubhash PublicHash) Equal(b PublicHash) bool {
	return bytes.Equal(pubhash[:], b[:])
}

// String returns the hex string of the public hash
func (pubhash PublicHash) String() string {
	return hex.EncodeToString(pubhash[:])
}

// Clone returns the clonend value of it
func (pubhash PublicHash) Clone() PublicHash {
	var cp PublicHash
	copy(cp[:], pubhash[:])
	return cp
}
