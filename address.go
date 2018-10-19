package common

import (
	"bytes"
	"encoding/binary"
	"io"

	"git.fleta.io/fleta/common/util"
	"github.com/mr-tron/base58/base58"
)

// AddressSize TODO
const AddressSize = 20

// Address TODO
type Address [AddressSize]byte

// NewAddress TODO
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

// Equal TODO
func (addr Address) Equal(b Address) bool {
	return bytes.Equal(addr[:], b[:])
}

// String TODO
func (addr Address) String() string {
	return base58.Encode(bytes.TrimRight(addr[:], string([]byte{0})))
}

// Clone TODO
func (addr Address) Clone() Address {
	var cp Address
	copy(cp[:], addr[:])
	return cp
}

// AddressFromString TODO
func AddressFromString(str string) (Address, error) {
	bs, err := base58.Decode(str)
	if err != nil {
		return Address{}, err
	}
	var addr Address
	copy(addr[:], bs)
	return addr, nil
}
