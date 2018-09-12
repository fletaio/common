package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/util"
)

const coordinateSize = 6

// Coordinate TODO
type Coordinate [coordinateSize]byte

// WriteTo TODO
func (crd *Coordinate) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(crd[:]); err != nil {
		return int64(n), err
	} else if n != coordinateSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (crd *Coordinate) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(crd[:]); err != nil {
		return int64(n), err
	} else if n != coordinateSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal TODO
func (crd Coordinate) Equal(b Coordinate) bool {
	return bytes.Equal(crd[:], b[:])
}

// String TODO
func (crd Coordinate) String() string {
	return hex.EncodeToString(crd[:])
}
