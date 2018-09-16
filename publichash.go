package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/hash"
	"git.fleta.io/fleta/common/util"
)

// PublicHashSize TODO
const PublicHashSize = 33

// PublicHash TODO
type PublicHash [PublicHashSize]byte

// NewPublicHash TODO
func NewPublicHash(pubkey PublicKey) PublicHash {
	h := hash.DoubleHash(pubkey[:])
	var ph PublicHash
	ph[0] = ChecksumFromPublicKey(pubkey)
	copy(ph[1:], h[:])
	return ph
}

// WriteTo TODO
func (pubkey *PublicHash) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(pubkey[:]); err != nil {
		return int64(n), err
	} else if n != PublicHashSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (pubkey *PublicHash) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(pubkey[:]); err != nil {
		return int64(n), err
	} else if n != PublicHashSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal TODO
func (pubkey PublicHash) Equal(b PublicHash) bool {
	return bytes.Equal(pubkey[:], b[:])
}

// String TODO
func (pubkey PublicHash) String() string {
	return hex.EncodeToString(pubkey[:])
}

// ChecksumFromPublicKey TODO
func ChecksumFromPublicKey(pubkey PublicKey) byte {
	var cs byte
	for i := 0; i < len(pubkey); i++ {
		cs = cs ^ pubkey[i]
	}
	return cs
}
