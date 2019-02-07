package hash

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// Hash returns the Hash256 value of the data
func Hash(data []byte) Hash256 {
	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		panic(err)
	}
	bs := h.Sum(nil)
	var hash Hash256
	copy(hash[:], bs)
	return hash
}

// DoubleHash returns the result of Hash(Hash(data))
func DoubleHash(data []byte) Hash256 {
	h1 := Hash(data)
	return Hash(h1[:])
}

// TwoHash returns the result of Hash(h1,'h',h2)
func TwoHash(h1 Hash256, h2 Hash256) Hash256 {
	data := make([]byte, Hash256Size*2+1)
	copy(data, h1[:])
	data[Hash256Size] = 'h'
	copy(data[Hash256Size+1:], h2[:])
	return Hash(data)
}

// ParseHex parses the Hash256 from the hex string
func ParseHex(str string) (Hash256, error) {
	bs, err := hex.DecodeString(str)
	if err != nil {
		return Hash256{}, err
	}
	if len(bs) != Hash256Size {
		return Hash256{}, ErrInvalidHashSize
	}
	var h Hash256
	copy(h[:], bs)
	return h, nil
}

// MustParseHex panics when parsing fails.
func MustParseHex(str string) Hash256 {
	h, err := ParseHex(str)
	if err != nil {
		panic(err)
	}
	return h
}

// DoubleHashByWriterTo returns the result of Hash(io.Writer)
func DoubleHashByWriterTo(w io.WriterTo) Hash256 {
	var buffer bytes.Buffer
	if _, err := w.WriteTo(&buffer); err != nil {
		panic(err)
	}
	return DoubleHash(buffer.Bytes())
}
