//Basic types
//Go's basic types are
//
//bool
//
//string
//
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//
//byte // alias for uint8
//
//rune // alias for int32
//     // represents a Unicode code point
//
//float32 float64
//
//complex64 complex128
//The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
//
//The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	aa int        = 7
	bb int8       = 15
	cc int16      = 23
	dd int32      = 99
	ee int64      = 18
	ff uint       = 2
	gg uint8      = 44
	hh uint16     = 56
	ii uint32     = 654
	jj uint64     = 213
	kk uintptr    = 16
	ll byte       = 7
	mm float32    = 21321.56756
	nn float64    = 51134213213.2132
	oo complex128 = cmplx.Sqrt(112i - 18)
	pp bool       = true
	qq string     = "牛逼啊"
)

func main() {
	fmt.Printf("Type is: %T , value is %+v\n", aa, aa)
	fmt.Printf("Type is: %T , value is %+v\n", bb, bb)
	fmt.Printf("Type is: %T , value is %+v\n", cc, cc)
	fmt.Printf("Type is: %T , value is %+v\n", dd, dd)
	fmt.Printf("Type is: %T , value is %+v\n", ee, ee)
	fmt.Printf("Type is: %T , value is %+v\n", ff, ff)
	fmt.Printf("Type is: %T , value is %+v\n", gg, gg)
	fmt.Printf("Type is: %T , value is %+v\n", hh, hh)
	fmt.Printf("Type is: %T , value is %+v\n", ii, ii)
	fmt.Printf("Type is: %T , value is %+v\n", jj, jj)
	fmt.Printf("Type is: %T , value is %+v\n", kk, kk)
	fmt.Printf("Type is: %T , value is %+v\n", ll, ll)
	fmt.Printf("Type is: %T , value is %+v\n", mm, mm)
	fmt.Printf("Type is: %T , value is %+v\n", nn, nn)
	fmt.Printf("Type is: %T , value is %+v\n", oo, oo)
	fmt.Printf("Type is: %T , value is %+v\n", pp, pp)
	fmt.Printf("Type is: %T , value is %+v\n", qq, qq)
}
