package common

import (
	"encoding/binary"
	"io"
	"strconv"
)

const coordinateSize = 6

// Coordinate TODO
type Coordinate struct {
	Height uint32
	Index  uint16
}

// WriteTo TODO
func (crd *Coordinate) WriteTo(w io.Writer) (int64, error) {
	/*
		if n, err := w.Write(crd[:]); err != nil {
			return int64(n), err
		} else if n != coordinateSize {
			return int64(n), util.ErrInvalidLength
		} else {
			return int64(n), nil
		}
	*/
	return 0, nil
}

// ReadFrom TODO
func (crd *Coordinate) ReadFrom(r io.Reader) (int64, error) {
	/*
		if n, err := r.Read(crd[:]); err != nil {
			return int64(n), err
		} else if n != coordinateSize {
			return int64(n), util.ErrInvalidLength
		} else {
			return int64(n), nil
		}
	*/
	return 0, nil
}

// Equal TODO
func (crd *Coordinate) Equal(b *Coordinate) bool {
	//return bytes.Equal(crd[:], b[:])
	return crd.Height == b.Height && crd.Index == b.Index
}

// Clone TODO
func (crd *Coordinate) Clone() *Coordinate {
	return &Coordinate{
		Height: crd.Height,
		Index:  crd.Index,
	}
}

// Bytes TODO
func (crd *Coordinate) Bytes() []byte {
	bs := make([]byte, 6)
	binary.LittleEndian.PutUint32(bs, crd.Height)
	binary.LittleEndian.PutUint16(bs[4:], crd.Index)
	//return hex.EncodeToString(crd[:])
	return bs
}

// String TODO
func (crd *Coordinate) String() string {
	return strconv.FormatUint(uint64(crd.Height)<<16+uint64(crd.Index), 10)
	//return hex.EncodeToString(crd[:])
}
