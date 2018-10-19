package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/util"
)

const signatureSize = 65

// Signature TODO
type Signature [signatureSize]byte

// WriteTo TODO
func (sig *Signature) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(sig[:]); err != nil {
		return int64(n), err
	} else if n != signatureSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom TODO
func (sig *Signature) ReadFrom(r io.Reader) (int64, error) {
	if n, err := r.Read(sig[:]); err != nil {
		return int64(n), err
	} else if n != signatureSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// Equal TODO
func (sig Signature) Equal(b Signature) bool {
	return bytes.Equal(sig[:], b[:])
}

// String TODO
func (sig Signature) String() string {
	return hex.EncodeToString(sig[:])
}

// Clone TODO
func (sig Signature) Clone() Signature {
	var cp Signature
	copy(cp[:], sig[:])
	return cp
}
