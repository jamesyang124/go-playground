package basic

import "time"

/*

bool
string

Numeric types:

uint      either 32 or 64 bits
int       same size as uint
uintptr   an unsigned integer large enough to store the
          uninterpreted bits of a pointer value
uint8     the set of all unsigned 8-bit integers (0 to 255)
uint16    the set of all unsigned 16-bit integers (0 to 65535)
uint32    the set of all unsigned 32-bit integers (0 to 4294967295)
uint64    the set of all unsigned 64-bit integers (0 to
          18446744073709551615)

int8      the set of all signed 8-bit integers (-128 to 127)
int16     the set of all signed 16-bit integers (-32768 to 32767)
int32     the set of all signed 32-bit integers (-2147483648 to
          2147483647)
int64     the set of all signed 64-bit integers
          (-9223372036854775808 to 9223372036854775807)

float32   the set of all IEEE-754 32-bit floating-point numbers
float64   the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and
            imaginary parts
complex128  the set of all complex numbers with float64 real and
            imaginary parts

byte        alias for uint8
rune        alias for int32 (represents a Unicode code point)

*/

// TypeConversion (x string)
// The expression T(v) converts the value v to the type T
func TypeConversion(x uint) (y uint32) {
	y = uint32(x)
	return
}

// TypeAssertion (a interface{})
// A type assertion takes a value
// and tries to create another version in the specified explicit type.
func TypeAssertion(a interface{}) (z map[string]interface{}) {
	foo := map[string]interface{}{
		"Matt": 42,
	}
	a = foo

	// https://golang.org/doc/effective_go.html#interface_conversions
	// type assertion => a.(map[string]interface{})
	z, ok := a.(map[string]interface{})

	// map[string]interface{}
	// indicate each map value can be arbitrary type ex:
	// {
	//    updagted_at: 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
	//    Matt: 42
	//    Ivy: stringVal
	// }

	if ok {
		z["updated_at"] = time.Now()
		z["Matt"] = 42
	}
	return
}

// TypeSwitch (a interface{})
// interface{} is act as aritrary type
// as an empty interface with zero methods
// similar idea as AnyVal | Any in scala
// run-time decided
func TypeSwitch(a interface{}) string {
	// a.(type) with keyword type
	switch z := a.(type) {
	case string:
		return z
	case uint32:
		return "is an integer"
	}

	return "none"
}
