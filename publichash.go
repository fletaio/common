package common

import (
	"bytes"
	"io"

	"git.fleta.io/fleta/common/hash"
	"git.fleta.io/fleta/common/util"
	"github.com/mr-tron/base58/base58"
)

// PublicHashSize is 31 bytes
const PublicHashSize = 31

// PublicHash is the [PublicHashSize]byte with methods
type PublicHash [PublicHashSize]byte

// NewPublicHash returns the PublicHash of the pubkey
func NewPublicHash(pubkey PublicKey) PublicHash {
	h := hash.DoubleHash(pubkey[:])
	var ph PublicHash
	ph[0] = pubkey.Checksum()
	copy(ph[1:], h[:len(h)-2])
	return ph
}

// WriteTo is a serialization function
func (pubhash *PublicHash) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(pubhash[:]); err != nil {
		return int64(n), err
	} else if n != PublicHashSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// UnmarshalJSON is a unmarshaler function
func (pubhash *PublicHash) UnmarshalJSON(bs []byte) error {
	if len(bs) < 3 {
		return ErrInvalidPublicHashFormat
	}
	if bs[0] != '"' || bs[len(bs)-1] != '"' {
		return ErrInvalidPublicHashFormat
	}
	v, err := ParsePublicHash(string(bs[1 : len(bs)-1]))
	if err != nil {
		return err
	}
	copy(pubhash[:], v[:])
	return nil
}

// MarshalJSON is a marshaler function
func (pubhash *PublicHash) MarshalJSON() ([]byte, error) {
	return []byte(`"` + pubhash.String() + `"`), nil
}

// ReadFrom is a deserialization function
func (pubhash *PublicHash) ReadFrom(r io.Reader) (int64, error) {
	return util.FillBytes(r, pubhash[:])
}

// Equal checks that two values is same or not
func (pubhash PublicHash) Equal(b PublicHash) bool {
	return bytes.Equal(pubhash[:], b[:])
}

// String returns a base58 value of the public hash
func (pubhash PublicHash) String() string {
	return base58.Encode(bytes.TrimRight(pubhash[:], string([]byte{0})))
}

// Clone returns the clonend value of it
func (pubhash PublicHash) Clone() PublicHash {
	var cp PublicHash
	copy(cp[:], pubhash[:])
	return cp
}

// ParsePublicHash parse the public hash from the string
func ParsePublicHash(str string) (PublicHash, error) {
	bs, err := base58.Decode(str)
	if err != nil {
		return PublicHash{}, err
	}
	var pubhash PublicHash
	copy(pubhash[:], bs)
	return pubhash, nil
}

// MustParsePublicHash panic when error occurred
func MustParsePublicHash(str string) PublicHash {
	pubhash, err := ParsePublicHash(str)
	if err != nil {
		panic(err)
	}
	return pubhash
}
