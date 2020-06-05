//Methods
//Go does not have classes. However, you can define methods on types.
//
//A method is a function with a special receiver argument.
//
//The receiver appears in its own argument list between the func keyword and the method name.
//
//In this example, the Abs method has a receiver of type Vertex named v.
package main

import "fmt"

type Tri struct {
	blen float64
	hlen float64
}

func (tri Tri) Triangle() float64 {
	return tri.blen*tri.hlen/2
}

func main() {
	Tri := Tri{5, 6}
	fmt.Printf("Triangle size is %#v.\n", Tri.Triangle())
}