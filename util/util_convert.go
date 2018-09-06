package util

import (
	"encoding/binary"
)

// Uint64ToBytes TODO
func Uint64ToBytes(v uint64) []byte {
	BNum := make([]byte, 8)
	binary.LittleEndian.PutUint64(BNum, v)
	return BNum
}

func Uint32ToBytes(v uint32) []byte {
	BNum := make([]byte, 4)
	binary.LittleEndian.PutUint32(BNum, v)
	return BNum
}

func Uint16ToBytes(v uint16) []byte {
	BNum := make([]byte, 2)
	binary.LittleEndian.PutUint16(BNum, v)
	return BNum
}

func Uint48ToBytes(a uint32, b uint16) []byte {
	BNum := make([]byte, 6)
	binary.LittleEndian.PutUint32(BNum, a)
	binary.LittleEndian.PutUint16(BNum[4:], b)
	return BNum
}

func BytesToUint16(v []byte) uint16 {
	return binary.LittleEndian.Uint16(v)
}

func BytesToUint32(v []byte) uint32 {
	return binary.LittleEndian.Uint32(v)
}

func BytesToUint64(v []byte) uint64 {
	return binary.LittleEndian.Uint64(v)
}
