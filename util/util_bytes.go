package util

import (
	"encoding/binary"
	"io"
)

// WriteUint64 writes the uint64 number to the writer
func WriteUint64(w io.Writer, num uint64) (int64, error) {
	BNum := make([]byte, 8)
	binary.LittleEndian.PutUint64(BNum, num)
	if n, err := w.Write(BNum); err != nil {
		return int64(n), err
	} else if n != 8 {
		return int64(n), ErrInvalidLength
	} else {
		return 8, nil
	}
}

// WriteUint32 writes the uint32 number to the writer
func WriteUint32(w io.Writer, num uint32) (int64, error) {
	BNum := make([]byte, 4)
	binary.LittleEndian.PutUint32(BNum, num)
	if n, err := w.Write(BNum); err != nil {
		return int64(n), err
	} else if n != 4 {
		return int64(n), ErrInvalidLength
	} else {
		return 4, nil
	}
}

// WriteUint16 writes the uint16 number to the writer
func WriteUint16(w io.Writer, num uint16) (int64, error) {
	BNum := make([]byte, 2)
	binary.LittleEndian.PutUint16(BNum, num)
	if n, err := w.Write(BNum); err != nil {
		return int64(n), err
	} else if n != 2 {
		return int64(n), ErrInvalidLength
	} else {
		return 2, nil
	}
}

// WriteUint8 writes the uint8 number to the writer
func WriteUint8(w io.Writer, num uint8) (int64, error) {
	if n, err := w.Write([]byte{byte(num)}); err != nil {
		return int64(n), err
	} else if n != 1 {
		return int64(n), ErrInvalidLength
	} else {
		return 1, nil
	}
}

// WriteBytes writes the byte array bytes with the var-length-bytes to the writer
func WriteBytes(w io.Writer, bs []byte) (int64, error) {
	var wrote int64
	if len(bs) < 254 {
		if n, err := WriteUint8(w, uint8(len(bs))); err != nil {
			return wrote, err
		} else {
			wrote += n
		}
		if n, err := w.Write(bs); err != nil {
			return wrote, err
		} else {
			wrote += int64(n)
		}
	} else if len(bs) < 65536 {
		if n, err := WriteUint8(w, 254); err != nil {
			return wrote, err
		} else {
			wrote += n
		}
		if n, err := WriteUint16(w, uint16(len(bs))); err != nil {
			return wrote, err
		} else {
			wrote += n
		}
		if n, err := w.Write(bs); err != nil {
			return wrote, err
		} else {
			wrote += int64(n)
		}
	} else {
		if n, err := WriteUint8(w, 255); err != nil {
			return wrote, err
		} else {
			wrote += n
		}
		if n, err := WriteUint32(w, uint32(len(bs))); err != nil {
			return wrote, err
		} else {
			wrote += n
		}
		if n, err := w.Write(bs); err != nil {
			return wrote, err
		} else {
			wrote += int64(n)
		}
	}
	return wrote, nil
}

// WriteString writes the string with the var-length-byte to the writer
func WriteString(w io.Writer, str string) (int64, error) {
	return WriteBytes(w, []byte(str))
}

// WriteBool writes the bool using a uint8 to the writer
func WriteBool(w io.Writer, b bool) (int64, error) {
	if b {
		return WriteUint8(w, 1)
	} else {
		return WriteUint8(w, 0)
	}
}

// ReadUint64 reads a uint64 number from the reader
func ReadUint64(r io.Reader) (uint64, int64, error) {
	BNum := make([]byte, 8)
	n, err := r.Read(BNum)
	if err != nil {
		return 0, int64(n), err
	}
	if n != 8 {
		return 0, int64(n), ErrInvalidLength
	}
	return binary.LittleEndian.Uint64(BNum), 8, nil
}

// ReadUint32 reads a uint32 number from the reader
func ReadUint32(r io.Reader) (uint32, int64, error) {
	BNum := make([]byte, 4)
	n, err := r.Read(BNum)
	if err != nil {
		return 0, int64(n), err
	}
	if n != 4 {
		return 0, int64(n), ErrInvalidLength
	}
	return binary.LittleEndian.Uint32(BNum), 4, nil
}

// ReadUint16 reads a uint16 number from the reader
func ReadUint16(r io.Reader) (uint16, int64, error) {
	BNum := make([]byte, 2)
	n, err := r.Read(BNum)
	if err != nil {
		return 0, int64(n), err
	}
	if n != 2 {
		return 0, int64(n), ErrInvalidLength
	}
	return binary.LittleEndian.Uint16(BNum), 2, nil
}

// ReadUint8 reads a uint8 number from the reader
func ReadUint8(r io.Reader) (uint8, int64, error) {
	BNum := make([]byte, 1)
	n, err := r.Read(BNum)
	if err != nil {
		return 0, int64(n), err
	}
	if n != 1 {
		return 0, int64(n), ErrInvalidLength
	}
	return uint8(BNum[0]), 1, nil
}

// ReadBytes reads a byte array from the reader
func ReadBytes(r io.Reader) ([]byte, int64, error) {
	var bs []byte
	var read int64
	if Len, n, err := ReadUint8(r); err != nil {
		return nil, read, err
	} else if Len < 254 {
		read += n
		bs = make([]byte, Len)
		if n, err := r.Read(bs); err != nil {
			return nil, read, err
		} else {
			read += int64(n)
		}
		return bs, read, nil
	} else if Len == 254 {
		if Len, n, err := ReadUint16(r); err != nil {
			return nil, read, err
		} else {
			read += n
			bs = make([]byte, Len)
			if n, err := r.Read(bs); err != nil {
				return nil, read, err
			} else {
				read += int64(n)
			}
		}
		return bs, read, nil
	} else {
		if Len, n, err := ReadUint32(r); err != nil {
			return nil, read, err
		} else {
			read += n
			bs = make([]byte, Len)
			if n, err := r.Read(bs); err != nil {
				return nil, read, err
			} else {
				read += int64(n)
			}
		}
		return bs, read, nil
	}
}

// ReadString reads a string array from the reader
func ReadString(r io.Reader) (string, int64, error) {
	if bs, n, err := ReadBytes(r); err != nil {
		return "", n, err
	} else {
		return string(bs), n, err
	}
}

// ReadBool reads a bool using a uint8 from the reader
func ReadBool(r io.Reader) (bool, int64, error) {
	if v, n, err := ReadUint8(r); err != nil {
		return false, n, err
	} else {
		return (v == 1), n, err
	}
}
