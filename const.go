package biu

import "errors"

const (
	zero  = byte('0')
	one   = byte('1')
	lsb   = byte('[') // left square brackets
	rsb   = byte(']') // right square brackets
	space = byte(' ')
)

// ErrBadString represents a error of input string's format is illegal .
var ErrBadString = errors.New("bad string format")
