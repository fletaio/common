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
	n, err := w.Write(hash[:])
	if err != nil {
		return int64(n), err
	} else if n != hash256Size {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (hash *Hash256) ReadFrom(r io.Reader) (int64, error) {
	n, err := r.Read(hash[:])
	if err != nil {
		return int64(n), err
	} else if n != hash256Size {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// String TODO
func (hash Hash256) String() string {
	return hex.EncodeToString(hash[:])
}

// MarshalJSON TODO
func (hash Hash256) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hash.String() + `"`), nil
}

// Debug TODO
func (hash Hash256) Debug() (string, error) {
	if bs, err := hash.MarshalJSON(); err != nil {
		return "", err
	} else {
		return string(bs), err
	}
}

// Equal TODO
func (hash Hash256) Equal(h Hash256) bool {
	return bytes.Equal(hash[:], h[:])
}
