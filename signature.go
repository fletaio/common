package common

import (
	"bytes"
	"encoding/hex"
	"io"

	"git.fleta.io/fleta/common/util"
)

// SignatureSize is 65 bytes
const SignatureSize = 65

// Signature is the [SignatureSize]byte with methods
type Signature [SignatureSize]byte

// WriteTo is a serialization function
func (sig *Signature) WriteTo(w io.Writer) (int64, error) {
	if n, err := w.Write(sig[:]); err != nil {
		return int64(n), err
	} else if n != SignatureSize {
		return int64(n), util.ErrInvalidLength
	} else {
		return int64(n), nil
	}
}

// ReadFrom is a deserialization function
func (sig *Signature) ReadFrom(r io.Reader) (int64, error) {
	return util.FillBytes(r, sig[:])
}

// MarshalJSON is a marshaler function
func (sig *Signature) MarshalJSON() ([]byte, error) {
	return []byte(`"` + sig.String() + `"`), nil
}

// UnmarshalJSON is a unmarshaler function
func (sig *Signature) UnmarshalJSON(bs []byte) error {
	if len(bs) < 3 {
		return ErrInvalidSignatureFormat
	}
	if bs[0] != '"' || bs[len(bs)-1] != '"' {
		return ErrInvalidSignatureFormat
	}
	v, err := ParseSignature(string(bs[1 : len(bs)-1]))
	if err != nil {
		return err
	}
	copy(sig[:], v[:])
	return nil
}

// Equal checks that two values is same or not
func (sig Signature) Equal(b Signature) bool {
	return bytes.Equal(sig[:], b[:])
}

// String returns the hex string of the signature
func (sig Signature) String() string {
	return hex.EncodeToString(sig[:])
}

// Clone returns the clonend value of it
func (sig Signature) Clone() Signature {
	var cp Signature
	copy(cp[:], sig[:])
	return cp
}

// ParseSignature parse the public hash from the string
func ParseSignature(str string) (Signature, error) {
	if len(str) != SignatureSize*2 {
		return Signature{}, ErrInvalidSignatureFormat
	}
	bs, err := hex.DecodeString(str)
	if err != nil {
		return Signature{}, err
	}
	var sig Signature
	copy(sig[:], bs)
	return sig, nil
}

// MustParseSignature panic when error occurred
func MustParseSignature(str string) Signature {
	sig, err := ParseSignature(str)
	if err != nil {
		panic(err)
	}
	return sig
}
