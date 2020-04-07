//Slices
//An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
//
//The type []T is a slice with elements of type T.
//
//A slice is formed by specifying two indices, a low and high bound, separated by a colon:
//
//a[low : high]
//This selects a half-open range which includes the first element, but excludes the last one.
//
//The following expression creates a slice which includes elements 1 through 3 of a:
//
//a[1:4]
package main

import "fmt"

var (
	C7int    = [6]int{1, 2, 3, 12, 4, 5}
	C7string = "Aglathrax hig' thrixa"
)

func main() {
	C7intSlice := C7int[1:3]
	C7string := C7string[3:9]
	fmt.Println(C7intSlice)
	fmt.Println(C7string)
}
