package common

import (
	"bytes"
	"encoding/binary"
	"io"

	"git.fleta.io/fleta/common/util"
	"github.com/mr-tron/base58/base58"
)

// AddressSize TODO
const AddressSize = 7

// AddressType TODO
type AddressType uint8

// Address TODO
type Address [AddressSize]byte

// NewAddress TODO
func NewAddress(Type AddressType, height uint32, index uint16) Address {
	var addr Address
	addr[0] = byte(Type)
	binary.LittleEndian.PutUint32(addr[1:], height)
	binary.LittleEndian.PutUint16(addr[5:], index)
	return addr
}

// WriteTo TODO
func (addr *Address) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(addr[:]); err != nil {
		return int64(n), err
	} else if n != AddressSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (addr *Address) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(addr[:]); err != nil {
		return int64(n), err
	} else if n != AddressSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Type TODO
func (addr Address) Type() AddressType {
	return AddressType(addr[0])
}

// Equal TODO
func (addr Address) Equal(b Address) bool {
	return bytes.Equal(addr[:], b[:])
}

// String TODO
func (addr Address) String() string {
	return base58.Encode(addr[:])
}

// AddressFromString TODO
func AddressFromString(str string) (Address, error) {
	bs, err := base58.Decode(str)
	if err != nil {
		return Address{}, err
	}
	if len(bs) != AddressSize {
		return Address{}, ErrInvalidAddressFormat
	}
	var addr Address
	copy(addr[:], bs)
	return addr, nil
}
