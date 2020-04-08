//Creating a slice with make
//Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
//
//The make function allocates a zeroed array and returns a slice that refers to that array:
//
//a := make([]int, 5)  // len(a)=5
//To specify a capacity, pass a third argument to make:
//
//b := make([]int, 0, 5) // len(b)=0, cap(b)=5
//
//b = b[:cap(b)] // len(b)=5, cap(b)=5
//b = b[1:]      // len(b)=4, cap(b)=4
package main

import "fmt"

func PrintInfo(a string, b []int) {
	fmt.Printf("%s has len=%d and cap=%d, equals to %v\n", a, len(b), cap(b), b)

}

func main() {
	int1 := make([]int, 5)
	PrintInfo("int1", int1)

	int2 := make([]int, 7, 11)
	PrintInfo("int2", int2)

	int3 := int2[:4]
	PrintInfo("int3", int3)

	int4 := int2[3:6]
	PrintInfo("int4", int4)
}
