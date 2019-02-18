package hash

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/util"
)

// Hash256Size is 32 bytes
const Hash256Size = 32

// Hash256 is the [Hash256Size]byte with methods
type Hash256 [Hash256Size]byte

// WriteTo is a sereialization function
func (hash *Hash256) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(hash[:]); err != nil {
		return int64(n), err
	} else if n != Hash256Size {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom is a desereialization function
func (hash *Hash256) ReadFrom(r io.Reader) (int64, error) {
	return util.FillBytes(r, hash[:])
}

// Equal checks that two values is same or not
func (hash Hash256) Equal(h Hash256) bool {
	return bytes.Equal(hash[:], h[:])
}

// String returns the hex string of the hash
func (hash Hash256) String() string {
	return hex.EncodeToString(hash[:])
}
