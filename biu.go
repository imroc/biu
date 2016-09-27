package biu

import "encoding/binary"

// GetBinaryStringOfByte get the string of a byte in binary format.
// for example: 01100100
func GetBinaryStringOfByte(b byte) string {
	buf := make([]byte, 0, 8)
	buf = appendBinaryStringOfByte(b, buf)
	return string(buf)
}

// GetBinaryStringOfBytes get the string of a byte's slice in binary format.
// for example: [00010010 01101001 01100100]
func GetBinaryStringOfBytes(bs []byte) string {
	l := len(bs)
	bl := l*8 + l + 1
	buf := make([]byte, 0, bl)
	buf = append(buf, lsb)
	for _, b := range bs {
		buf = appendBinaryStringOfByte(b, buf)
		buf = append(buf, space)
	}
	buf[bl-1] = rsb
	return string(buf)
}

// GetBinaryStringOfUInt32 get the string of a uinit32 number in binary format.
func GetBinaryStringOfUInt32(i uint32) string {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return GetBinaryStringOfBytes(bs)
}
