package common

import (
	"bytes"
	"encoding/binary"
	"io"

	"git.fleta.io/fleta/common/util"
	"github.com/mr-tron/base58/base58"
)

// AddressSize is 20 bytes
const AddressSize = 20

// Address is the [AddressSize]byte with methods
type Address [AddressSize]byte

// NewAddress returns a Address by the AccountCoordinate, by the ChainCoordinate and by the nonce
func NewAddress(accCoord *Coordinate, chainCoord *Coordinate, nonce uint64) Address {
	var addr Address
	copy(addr[:], accCoord.Bytes())
	if chainCoord != nil && (chainCoord.Height > 0 || chainCoord.Index > 0) || nonce > 0 {
		copy(addr[6:], chainCoord.Bytes())
	}
	if nonce > 0 {
		binary.LittleEndian.PutUint64(addr[12:], nonce)
	}
	return addr
}

// WriteTo is a serialization function
func (addr *Address) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(addr[:]); err != nil {
		return int64(n), err
	} else if n != AddressSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// MarshalJSON is a marshaler function
func (addr *Address) MarshalJSON() ([]byte, error) {
	return []byte(`"` + addr.String() + `"`), nil
}

// UnmarshalJSON is a unmarshaler function
func (addr *Address) UnmarshalJSON(bs []byte) error {
	if len(bs) < 3 {
		return ErrInvalidAddressFormat
	}
	if bs[0] != '"' || bs[len(bs)-1] != '"' {
		return ErrInvalidAddressFormat
	}
	v, err := ParseAddress(string(bs[1 : len(bs)-1]))
	if err != nil {
		return err
	}
	copy(addr[:], v[:])
	return nil
}

// ReadFrom is a deserialization function
func (addr *Address) ReadFrom(r io.Reader) (int64, error) {
	return util.FillBytes(r, addr[:])
}

// Equal checks that two values is same or not
func (addr Address) Equal(b Address) bool {
	return bytes.Equal(addr[:], b[:])
}

// String returns a base58 value of the address
func (addr Address) String() string {
	var bs []byte
	checksum := addr.Checksum()
	result := bytes.TrimRight(addr[:], string([]byte{0}))
	if len(result) < 7 {
		bs = make([]byte, 7)
		copy(bs[1:], result[:])
	} else if len(result) < 13 {
		bs = make([]byte, 13)
		copy(bs[1:], result[:])
	} else if len(result) < 21 {
		bs := make([]byte, 21)
		copy(bs[1:], result[:])
	}
	bs[0] = checksum

	return base58.Encode(bs)
}

// Clone returns the clonend value of it
func (addr Address) Clone() Address {
	var cp Address
	copy(cp[:], addr[:])
	return cp
}

// WithNonce returns derive the address using the nonce
func (addr Address) WithNonce(nonce uint64) Address {
	var cp Address
	copy(cp[:], addr[:])
	binary.LittleEndian.PutUint64(cp[12:], nonce)
	return cp
}

// Checksum returns the checksum byte
func (addr Address) Checksum() byte {
	var cs byte
	for _, c := range addr {
		cs = cs ^ c
	}
	return cs
}

// ParseAddress parse the address from the string
func ParseAddress(str string) (Address, error) {
	bs, err := base58.Decode(str)
	if err != nil {
		return Address{}, err
	}
	if len(bs) != 7 && len(bs) != 13 && len(bs) != 21 {
		return Address{}, ErrInvalidAddressFormat
	}
	cs := bs[0]
	var addr Address
	copy(addr[:], bs[1:])
	if cs != addr.Checksum() {
		return Address{}, ErrInvalidAddressCheckSum
	}
	return addr, nil
}

// MustParseAddress panic when error occurred
func MustParseAddress(str string) Address {
	addr, err := ParseAddress(str)
	if err != nil {
		//panic(err)
	}
	return addr
}
