//Methods are functions
//Remember: a method is just a function with a receiver argument.
//
//Here's Abs written as a regular function with no change in functionality.
package main

import "fmt"

type Tri struct {
	blen float64
	hlen float64
}

func Triangle(tri Tri) float64 {
	return tri.blen * tri.hlen / 2
}

func main() {
	Tri := Tri{5, 6}
	fmt.Printf("Triangle size is %#v.\n", Triangle(Tri))
}
