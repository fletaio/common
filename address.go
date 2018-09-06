package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/hash"
	"git.fleta.io/fleta/common/util"
)

// AddressSize TODO
const AddressSize = 20

// Address TODO
type Address [AddressSize]byte

// AddressFromPubkey TODO
func AddressFromPubkey(pubkey PublicKey) Address {
	var addr Address
	phash := hash.DoubleHash(pubkey[:])
	copy(addr[:], phash[len(phash)-len(addr):])
	return addr
}

// AddressFromString TODO
func AddressFromString(str string) (Address, error) {
	if len(str) != 2+AddressSize*2 {
		return Address{}, ErrInvalidAddressFormat
	}
	if str[:2] != "0x" {
		return Address{}, ErrInvalidAddressFormat
	}
	bs, err := hex.DecodeString(str[2:])
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
	return "0x" + hex.EncodeToString(addr[:])
}

// MarshalJSON TODO
func (addr Address) MarshalJSON() ([]byte, error) {
	return []byte(`"` + addr.String() + `"`), nil
}

// Debug TODO
func (addr Address) Debug() (string, error) {
	if bs, err := addr.MarshalJSON(); err != nil {
		return "", err
	} else {
		return string(bs), err
	}
}
