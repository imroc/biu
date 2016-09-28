package biu

import (
	"encoding/binary"
	"regexp"
)

// ByteToBinaryString get the string of a byte in binary format.
// for example: 01100100
func ByteToBinaryString(b byte) string {
	buf := make([]byte, 0, 8)
	buf = appendBinaryString(b, buf)
	return string(buf)
}

// BytesToBinaryString get the string of a byte's slice in binary format.
// for example: [00010010 01101001 01100100]
func BytesToBinaryString(bs []byte) string {
	l := len(bs)
	bl := l*8 + l + 1
	buf := make([]byte, 0, bl)
	buf = append(buf, lsb)
	for _, b := range bs {
		buf = appendBinaryString(b, buf)
		buf = append(buf, space)
	}
	buf[bl-1] = rsb
	return string(buf)
}

var rbDel = regexp.MustCompile(`[^01]`) // regex for delete useless string of binary format

// BinaryStringToBytes get the binary bytes according to the
// input string of binary format.
// for example, inupt string is "[00001010 10101010 01001010]",
// returns byte slice,whose length is 3 and if you print its'
// value,it may be like this:[10, 170, 174]
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

// ToBinaryString convert a lot of basic data type to string in binary format.
func ToBinaryString(v interface{}) string {

	return ""

}

// UInt32ToBinaryString get the string of a uinit32 number in binary format.
func UInt32ToBinaryString(i uint32) string {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return BytesToBinaryString(bs)
}
