//Range
//The range form of the for loop iterates over a slice or map.
//
//When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
//range和Python的区别在于这里返回的是两个值，第一个是下标序号，第二个是值。
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { //i返回了下标，v返回了值本身
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
