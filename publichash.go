package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/hash"
	"git.fleta.io/fleta/common/util"
)

// PublicHashSize TODO
const PublicHashSize = 31

// PublicHash TODO
type PublicHash [PublicHashSize]byte

// NewPublicHash TODO
func NewPublicHash(pubkey PublicKey) PublicHash {
	h := hash.DoubleHash(pubkey[:])
	var ph PublicHash
	ph[0] = ChecksumFromPublicKey(pubkey)
	copy(ph[1:], h[:len(h)-2])
	return ph
}

// WriteTo TODO
func (pubhash *PublicHash) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(pubhash[:]); err != nil {
		return int64(n), err
	} else if n != PublicHashSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (pubhash *PublicHash) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(pubhash[:]); err != nil {
		return int64(n), err
	} else if n != PublicHashSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal TODO
func (pubhash PublicHash) Equal(b PublicHash) bool {
	return bytes.Equal(pubhash[:], b[:])
}

// String TODO
func (pubhash PublicHash) String() string {
	return hex.EncodeToString(pubhash[:])
}

// Clone TODO
func (pubhash PublicHash) Clone() PublicHash {
	var cp PublicHash
	copy(cp[:], pubhash[:])
	return cp
}

// ChecksumFromPublicKey TODO
func ChecksumFromPublicKey(pubhash PublicKey) byte {
	var cs byte
	for i := 0; i < len(pubhash); i++ {
		cs = cs ^ pubhash[i]
	}
	return cs
}
