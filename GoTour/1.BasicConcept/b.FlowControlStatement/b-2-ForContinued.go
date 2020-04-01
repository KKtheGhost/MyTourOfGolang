//For continued
//The init and post statements are optional.
package main

import "fmt"

func Fibonacci(lim int) int {
	a, b, c := 1, 1, 0
	for a < lim {
		fmt.Println(a)
		c = a
		a += b
		b = c
	}
	return a
}

func main() {
	lim1, lim2 := 1000, 10000
	fmt.Println(Fibonacci(lim1))
	fmt.Println(Fibonacci(lim2))
}
