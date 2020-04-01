//For is Go's "while"
//At that point you can drop the semicolons: C's while is spelled for in Go.
package main

import "fmt"

func LeiJia(x int) int {
	i := 1
	for i < x {
		i += i
	}
	return i
}

func main() {
	fmt.Println(LeiJia(1000))
	fmt.Println(LeiJia(12345))
}
