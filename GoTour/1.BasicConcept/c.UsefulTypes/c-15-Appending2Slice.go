//Appending to a slice
//It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.
//
//func append(s []T, vs ...T) []T
//The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
//
//The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
//
//If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
//
//(To learn more about slices, read the Slices: usage and internals article.)
package main

import "fmt"

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var array []string
	printSlice(array)

	array = append(array, "Animal")
	printSlice(array)

	array = append(array, " Crossing!")
	printSlice(array)

	array = append(array, "New", " Horizons")
	printSlice(array)
}
