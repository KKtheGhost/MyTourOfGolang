//Interfaces are implemented implicitly
//A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
//
//Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
package main

import "fmt"

type Single int64

type Binary struct {
	X int64
	Y int64
}

type NumbDeal interface {
	Deal()
}

func (s Single) Deal() {
	s = s*s + 2*s
	fmt.Println(s)
}

func (b Binary) Deal() {
	b.X = b.X * b.Y
	b.Y = b.X + b.Y
	res := b.X * b.Y
	fmt.Println(res)
}

func main() {
	var num NumbDeal
	s := Single(15)
	b := Binary{2, 6}

	num = &s
	num.Deal()

	num = &b
	num.Deal()
}
