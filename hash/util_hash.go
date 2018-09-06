package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash TODO
func Hash(data []byte) Hash256 {
	h := sha256.New()
	h.Write(data)
	bs := h.Sum(nil)
	var hash Hash256
	copy(hash[:], bs)
	return hash
}

// DoubleHash TODO
func DoubleHash(data []byte) Hash256 {
	h1 := Hash(data)
	return Hash(h1[:])
}

// TwoHash TODO
func TwoHash(h1 Hash256, h2 Hash256) Hash256 {
	data := make([]byte, hash256Size*2+1)
	copy(data, h1[:])
	data[hash256Size] = 'h'
	copy(data[hash256Size+1:], h2[:])
	return Hash(data)
}

// ParseHex TODO
func ParseHex(str string) (Hash256, error) {
	bs, err := hex.DecodeString(str)
	if err != nil {
		return Hash256{}, err
	}
	if len(bs) != hash256Size {
		return Hash256{}, ErrInvalidHashSize
	}
	var h Hash256
	copy(h[:], bs)
	return h, nil
}

// MustParseHex TODO
func MustParseHex(str string) Hash256 {
	h, err := ParseHex(str)
	if err != nil {
		panic(err)
	}
	return h
}
