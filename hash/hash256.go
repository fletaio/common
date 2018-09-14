package hash

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/util"
)

const hash256Size = 32

// Hash256 TODO
type Hash256 [hash256Size]byte

// WriteTo TODO
func (hash *Hash256) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(hash[:]); err != nil {
		return int64(n), err
	} else if n != hash256Size {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (hash *Hash256) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(hash[:]); err != nil {
		return int64(n), err
	} else if n != hash256Size {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal TODO
func (hash Hash256) Equal(h Hash256) bool {
	return bytes.Equal(hash[:], h[:])
}

// String TODO
func (hash Hash256) String() string {
	return hex.EncodeToString(hash[:])
}
