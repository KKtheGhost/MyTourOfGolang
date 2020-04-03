//Struct Literals
//A struct literal denotes a newly allocated struct value by listing the values of its fields.
//
//You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
//
//The special prefix & returns a pointer to the struct value.
package main

import "fmt"

type C5 struct {
	aaa string
	bbb int
}

var (
	C51 = C5{aaa: "老铁们"}
	C52 = C5{bbb: 8}
	C53 = C5{}
	p   = &C5{"oh", 900}
)

func main() {
	fmt.Println(C51, C52, C53, (*p), p)
}
