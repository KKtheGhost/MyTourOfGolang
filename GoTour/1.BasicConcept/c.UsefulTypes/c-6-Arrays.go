//Arrays
//The type [n]T is an array of n values of type T.
//
//The expression
//
//var a [10]int
//declares a variable a as an array of ten integers.
//
//An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
package main

import "fmt"

var (
	allint    [3]int
	allstring [2]string
)

func main() {
	allint[0] = 22
	allint[1] = 14
	allstring[0] = "Fuck"
	allstring[1] = "Me"

	fmt.Println(allint, allstring)
}
