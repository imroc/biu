package biu

import (
	"encoding/binary"
	"regexp"
)

// ByteToBinaryString get the string in binary format of a byte.
func ByteToBinaryString(b byte) string {
	buf := make([]byte, 0, 8)
	buf = appendBinaryString(buf, b)
	return string(buf)
}

// BytesToBinaryString get the string in binary format of a byte's slice.
func BytesToBinaryString(bs []byte) string {
	l := len(bs)
	bl := l*8 + l + 1
	buf := make([]byte, 0, bl)
	buf = append(buf, lsb)
	for _, b := range bs {
		buf = appendBinaryString(buf, b)
		buf = append(buf, space)
	}
	buf[bl-1] = rsb
	return string(buf)
}

// regex for delete useless string which is going to be in binary format.
var rbDel = regexp.MustCompile(`[^01]`)

// BinaryStringToBytes get the binary bytes according to the
// input string which is in binary format.
func BinaryStringToBytes(s string) (bs []byte) {
	if len(s) == 0 {
		panic(ErrEmptyString)
	}

	s = rbDel.ReplaceAllString(s, "")
	l := len(s)
	if l == 0 || l%8 != 0 {
		panic(ErrBadStringFormat)
	}

	bs = make([]byte, 0, l/8)

	var n uint8
	for i, b := range []byte(s) {
		m := i % 8
		switch b {
		case one:
			n += uint8arr[m]
		}
		if m == 7 {
			bs = append(bs, n)
			n = 0
		}
	}
	return
}

// UInt32ToBinaryString get the string of a uinit32 number in binary format.
func UInt32ToBinaryString(i uint32) string {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return BytesToBinaryString(bs)
}
