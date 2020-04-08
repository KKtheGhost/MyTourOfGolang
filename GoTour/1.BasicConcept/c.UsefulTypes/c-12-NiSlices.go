//Nil slices
//The zero value of a slice is nil.
//
//A nil slice has a length and capacity of 0 and has no underlying array.
package main

import "fmt"

func main() {
	var nilarray []string
	fmt.Println(nilarray, len(nilarray), cap(nilarray))
	if nilarray == nil {
		fmt.Println("Empty!")
	}
}
