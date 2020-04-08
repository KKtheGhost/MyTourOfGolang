//Slice length and capacity
//A slice has both a length and a capacity.
//
//The length of a slice is the number of elements it contains.
//
//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
//
//The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
//
//You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
//
//我们可以把容量当做成总长度减去左指针走过的元素值
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
