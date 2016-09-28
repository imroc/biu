package biu

import "encoding/binary"

// Uint16ToBinaryString get the string of a uint16 number in binary format.
func Uint16ToBinaryString(i uint16) string {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return BytesToBinaryString(bs)
}

// Uint32ToBinaryString get the string of a uint32 number in binary format.
func Uint32ToBinaryString(i uint32) string {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return BytesToBinaryString(bs)
}

// Uint64ToBinaryString get the string of a uint64 number in binary format.
func Uint64ToBinaryString(i uint64) string {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, i)
	return BytesToBinaryString(bs)
}
