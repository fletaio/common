package common

import (
	"bytes"
	"io"

	"github.com/mr-tron/base58/base58"

	"git.fleta.io/fleta/common/hash"
	"git.fleta.io/fleta/common/util"
)

// AddressSize TODO
const AddressSize = 40

// Address TODO
type Address [AddressSize]byte

// AddressType TODO
type AddressType uint8

// Type TODO
func (addr Address) Type() AddressType {
	return AddressType(addr[6])
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
	return base58.Encode(addr[:])
}

// AddressFromPubkey TODO
func AddressFromPubkey(crd Coordinate, t AddressType, pubkey PublicKey) Address {
	phash := hash.DoubleHash(pubkey[:])
	addr := AddressFromHash(crd, t, phash, ChecksumFromPublicKey(pubkey))
	return addr
}

// ConvertAddressType TODO
func ConvertAddressType(addr Address, t AddressType) Address {
	var a Address
	copy(a[:], addr[:])
	a[coordinateSize] = byte(t)
	return a
}

// AddressFromHash TODO
func AddressFromHash(crd Coordinate, t AddressType, h hash.Hash256, checksum byte) Address {
	var addr Address
	idx := 0
	copy(addr[idx:], crd[:]) // 6 bytes
	idx += len(crd)
	copy(addr[idx:], []byte{byte(t)}) // 1 bytes
	idx++
	copy(addr[idx:], h[:]) // 32 bytes
	idx += len(h)
	addr[idx] = checksum // 1 bytes
	return addr
}

// ChecksumFromPublicKey TODO
func ChecksumFromPublicKey(pubkey PublicKey) byte {
	var cs byte
	for i := 0; i < len(pubkey); i++ {
		cs = cs ^ pubkey[i]
	}
	return cs
}

// ChecksumFromAddresses TODO
func ChecksumFromAddresses(addrs []Address) byte {
	var cs byte
	for _, addr := range addrs {
		for i := 0; i < len(addr); i++ {
			cs = cs ^ addr[i]
		}
	}
	return cs
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
