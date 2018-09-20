package util

import (
	"encoding/binary"
	"io"
)

// WriteUint64 TODO
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

// WriteUint32 TODO
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

// WriteUint16 TODO
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

// WriteUint8 TODO
func WriteUint8(w io.Writer, num uint8) (int64, error) {
	if n, err := w.Write([]byte{byte(num)}); err != nil {
		return int64(n), err
	} else if n != 1 {
		return int64(n), ErrInvalidLength
	} else {
		return 1, nil
	}
}

// WriteBytes8 TODO
func WriteBytes8(w io.Writer, bs []byte) (int64, error) {
	var wrote int64
	n, err := WriteUint8(w, uint8(len(bs)))
	if err != nil {
		return wrote, err
	} else {
		wrote += n
	}
	if n, err := w.Write(bs); err != nil {
		return wrote, err
	} else {
		wrote += int64(n)
	}
	return wrote, nil
}

// WriteBytes TODO
func WriteBytes(w io.Writer, bs []byte) (int64, error) {
	var wrote int64
	n, err := WriteUint32(w, uint32(len(bs)))
	if err != nil {
		return wrote, err
	} else {
		wrote += n
	}
	if n, err := w.Write(bs); err != nil {
		return wrote, err
	} else {
		wrote += int64(n)
	}
	return wrote, nil
}

// WriteString8 TODO
func WriteString8(w io.Writer, str string) (int64, error) {
	return WriteBytes8(w, []byte(str))
}

// WriteString TODO
func WriteString(w io.Writer, str string) (int64, error) {
	return WriteBytes(w, []byte(str))
}

// WriteBool TODO
func WriteBool(w io.Writer, b bool) (int64, error) {
	if b {
		return WriteUint8(w, 1)
	} else {
		return WriteUint8(w, 0)
	}
}

// ReadUint64 TODO
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

// ReadUint32 TODO
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

// ReadUint16 TODO
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

// ReadUint8 TODO
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

// ReadBytes TODO
func ReadBytes(r io.Reader) ([]byte, int64, error) {
	var bs []byte
	var read int64
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

// ReadBytes8 TODO
func ReadBytes8(r io.Reader) ([]byte, int64, error) {
	var bs []byte
	var read int64
	if Len, n, err := ReadUint8(r); err != nil {
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

// ReadString8 TODO
func ReadString8(r io.Reader) (string, int64, error) {
	if bs, n, err := ReadBytes8(r); err != nil {
		return "", n, err
	} else {
		return string(bs), n, err
	}
}

// ReadString TODO
func ReadString(r io.Reader) (string, int64, error) {
	if bs, n, err := ReadBytes(r); err != nil {
		return "", n, err
	} else {
		return string(bs), n, err
	}
}

// ReadBool TODO
func ReadBool(r io.Reader) (bool, int64, error) {
	if v, n, err := ReadUint8(r); err != nil {
		return false, n, err
	} else {
		return (v == 1), n, err
	}
}
