package common

import (
	"encoding/binary"
	"encoding/hex"
	"io"

	"github.com/fletaio/common/util"
)

// CoordinateSize is 6 bytes
const CoordinateSize = 6

// Coordinate is (BlockHeight, TransactionIndexOfTheBlock)
type Coordinate struct {
	Height uint32
	Index  uint16
}

// NewCoordinate returns a Coordinate
func NewCoordinate(Height uint32, Index uint16) *Coordinate {
	return &Coordinate{
		Height: Height,
		Index:  Index,
	}
}

// NewCoordinateByID returns a Coordinate using compacted id
func NewCoordinateByID(id uint64) *Coordinate {
	return &Coordinate{
		Height: uint32(id >> 32),
		Index:  uint16(id >> 16),
	}
}

// WriteTo is a serialization function
func (crd *Coordinate) WriteTo(w io.Writer) (int64, error) {
	var wrote int64
	if n, err := util.WriteUint32(w, crd.Height); err != nil {
		return wrote, err
	} else {
		wrote += n
	}
	if n, err := util.WriteUint16(w, crd.Index); err != nil {
		return wrote, err
	} else {
		wrote += n
	}
	return 0, nil
}

// ReadFrom is a deserialization function
func (crd *Coordinate) ReadFrom(r io.Reader) (int64, error) {
	var read int64
	if v, n, err := util.ReadUint32(r); err != nil {
		return read, err
	} else {
		read += n
		crd.Height = v
	}
	if v, n, err := util.ReadUint16(r); err != nil {
		return read, err
	} else {
		read += n
		crd.Index = v
	}
	return 0, nil
}

// MarshalJSON is a marshaler function
func (crd *Coordinate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + crd.String() + `"`), nil
}

// UnmarshalJSON is a unmarshaler function
func (crd *Coordinate) UnmarshalJSON(bs []byte) error {
	if len(bs) < 3 {
		return ErrInvalidCoordinateFormat
	}
	if bs[0] != '"' || bs[len(bs)-1] != '"' {
		return ErrInvalidCoordinateFormat
	}
	v, err := ParseCoordinate(string(bs[1 : len(bs)-1]))
	if err != nil {
		return err
	}
	crd.Height = v.Height
	crd.Index = v.Index
	return nil
}

// Equal checks that two values is same or not
func (crd *Coordinate) Equal(b *Coordinate) bool {
	return crd.Height == b.Height && crd.Index == b.Index
}

// Clone returns the clonend value of it
func (crd *Coordinate) Clone() *Coordinate {
	return &Coordinate{
		Height: crd.Height,
		Index:  crd.Index,
	}
}

// Bytes returns a byte array
func (crd *Coordinate) Bytes() []byte {
	bs := make([]byte, 6)
	binary.LittleEndian.PutUint32(bs, crd.Height)
	binary.LittleEndian.PutUint16(bs[4:], crd.Index)
	return bs
}

// ID returns a compacted id
func (crd *Coordinate) ID() uint64 {
	return uint64(crd.Height)<<32 | uint64(crd.Index)<<16
}

// String returns a hex value of the byte array
func (crd *Coordinate) String() string {
	return hex.EncodeToString(crd.Bytes())
}

// ParseCoordinate parse the public hash from the string
func ParseCoordinate(str string) (*Coordinate, error) {
	if len(str) != CoordinateSize*2 {
		return nil, ErrInvalidCoordinateFormat
	}
	bs, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	coord := &Coordinate{
		Height: binary.LittleEndian.Uint32(bs),
		Index:  binary.LittleEndian.Uint16(bs[4:]),
	}
	return coord, nil
}

// MustParseCoordinate panic when error occurred
func MustParseCoordinate(str string) *Coordinate {
	coord, err := ParseCoordinate(str)
	if err != nil {
		panic(err)
	}
	return coord
}
