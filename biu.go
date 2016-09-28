package biu

import "encoding/binary"

// ToBinaryString get string in binary format according to input data.
// The input data can be diffrent kinds of basic data type.
func ToBinaryString(v interface{}) (s string) {

	switch v := v.(type) {
	case []byte:
		s = BytesToBinaryString(v)
	case int8:
		s = ByteToBinaryString(uint8(v))
	case uint8: // byte
		s = ByteToBinaryString(v)
	case int16:
		s = Uint16ToBinaryString(uint16(v))
	case uint16:
		s = Uint16ToBinaryString(v)
	case int32:
		s = Uint32ToBinaryString(uint32(v))
	case uint32:
		s = Uint32ToBinaryString(v)
	case uint:
		s = Uint64ToBinaryString(uint64(v))
	case int:
		s = Uint64ToBinaryString(uint64(v))
	case int64:
		s = Uint64ToBinaryString(uint64(v))
	case uint64:
		s = Uint64ToBinaryString(v)
		//TODO add float number support
	default:
		panic(ErrTypeUnsupport)

	}
	return

}

// ReadBinaryString read the string in binary format into input data.
func ReadBinaryString(s string, data interface{}) (err error) {

	bs := BinaryStringToBytes(s)
	switch data := data.(type) {

	case *int8:
		*data = int8(bs[0])
	case *uint8:
		*data = bs[0]
	case *int16:
		*data = int16(bytesToUint16(bs))
	case *uint16:
		*data = bytesToUint16(bs)
	case *int32:
		*data = int32(bytesToUint32(bs))
	case *uint32:
		*data = bytesToUint32(bs)
	case *int64:
		*data = int64(bytesToUint64(bs))
	case *uint64:
		*data = bytesToUint64(bs)
		//TODO add float number support
	default:
		err = ErrTypeUnsupport
	}
	return
}

func bytesToUint16(bs []byte) uint16 {
	bs = fillBytes(bs, 2)
	return binary.BigEndian.Uint16(bs)
}

func bytesToUint32(bs []byte) uint32 {
	bs = fillBytes(bs, 4)
	return binary.BigEndian.Uint32(bs)
}

func bytesToUint64(bs []byte) uint64 {
	bs = fillBytes(bs, 8)
	return binary.BigEndian.Uint64(bs)
}

// fillBytes fills byte slice with zero bytes ahead when its'
// length is not greater than n.
func fillBytes(bs []byte, n int) []byte {

	l := len(bs)
	if l >= n {
		return bs
	}

	nbs := make([]byte, n)
	n -= l // n zero bytes need to fill

	for i := 0; i < n; i++ {
		nbs[i] = byte(0)
	}

	copy(nbs[n:], bs)

	return nbs
}

// ToHexString get string in Hexadecimal format according to input data.
// The input data can be diffrent kinds of basic data type.
func ToHexString(v interface{}) (s string) {
	//TODO implements it

	return
}

// ToOctalString get string in octal format according to input data.
// The input data can be diffrent kinds of basic data type.
func ToOctalString(v interface{}) (s string) {
	//TODO implements it

	return
}
