package common

import (
	"bytes"
	"encoding/binary"
	"io"

	"git.fleta.io/fleta/common/util"
	"github.com/mr-tron/base58/base58"
)

// TagSize TODO
const TagSize = 4

// Tag TODO
type Tag [TagSize]byte

// NewTag TODO
func NewTag(v uint32) Tag {
	var tag Tag
	binary.LittleEndian.PutUint32(tag[:], v)
	return tag
}

// WriteTo TODO
func (addr *Tag) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(addr[:]); err != nil {
		return int64(n), err
	} else if n != TagSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (addr *Tag) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(addr[:]); err != nil {
		return int64(n), err
	} else if n != TagSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal TODO
func (addr Tag) Equal(b Tag) bool {
	return bytes.Equal(addr[:], b[:])
}

// String TODO
func (addr Tag) String() string {
	return base58.Encode(addr[:])
}

// TagFromString TODO
func TagFromString(str string) (Tag, error) {
	bs, err := base58.Decode(str)
	if err != nil {
		return Tag{}, err
	}
	if len(bs) != TagSize {
		return Tag{}, ErrInvalidTagFormat
	}
	var addr Tag
	copy(addr[:], bs)
	return addr, nil
}
