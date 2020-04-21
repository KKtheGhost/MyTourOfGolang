//Exercise: Fibonacci closure
//Let's have some fun with functions.
//
//Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y, z := 0, 1, 0
	//闭包函数，x,y,z会被外部赋予，此处的外部指的是上一次循环。
	return func() int {
		z, x, y = x, y, x+y
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
