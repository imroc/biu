Biu
==============
biu is a golang binary util.
you can convert bytes to different string format(binary,hex,octal. support binary format only right now),
even can convert your data(basic data type) to different string format directly.
And you can also convert a string in binary,hex or octal format to a data with data type whatever you wanted.


## Features
 * high performance
 * easy to use
 * rich data type support

## Quick Start
##### Install
``` sh
$ go get github.com/imroc/biu
```
##### Bytes to string (binary format)
``` go
bs := []byte{1, 2, 3}
s := biu.BytesToBinaryString(bs)
fmt.Println(s) //[00000001 00000010 00000011]
fmt.Println(biu.ByteToBinaryString(byte(3))) //00000011
```
##### Any data type to string (binary format)
``` go
fmt.Println(biu.ToBinaryString(3)) //[00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011]
fmt.Println(biu.ToBinaryString(int8(4))) //00000100
fmt.Println(biu.ToBinaryString(uint16(2))) //[00000000 00000010]
fmt.Println(biu.ToBinaryString([]byte{1, 2, 3})) //[00000001 00000010 00000011]
```
##### String (binary format) to bytes
``` go
s := "[00000011 10000000]"
bs := biu.BinaryStringToBytes(s)
fmt.Printf("%#v\n", bs) //[]byte{0x3, 0x80}
```
##### String (binary format) to any data type
``` go
var a uint8
biu.ReadBinaryString("00000010", &a)
fmt.Println(a) //2

var b int16
biu.ReadBinaryString("[1 0000 0011]", &b)
fmt.Println(b) //259
```
## LICENSE
biu is is distributed under the terms of the MIT License.
