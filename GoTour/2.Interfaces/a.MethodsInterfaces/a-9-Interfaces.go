//Interfaces
//An interface type is defined as a set of method signatures.
//
//A value of interface type can hold any value that implements those methods.
//
//Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
//Very Important part of language Golang
package main

import (
	"fmt"
)

type Metrix struct {
	X float64
	Y float64
}

type Response struct {
	Status int64
	Result string
	Value  float64
}

type Inter int

type Face interface {
	Abc() Response
}

func (m Metrix) MyFloat() float64 {
	if m.X < 0 {
		return float64(-m.X + m.Y)
	}
	return float64(m.X + m.Y)
}

func (num Inter) Abc() Response {
	var res string
	var status int64
	if num < 0 {
		res = "Small"
		status = 302
	} else {
		res = "Big"
		status = 307
	}
	resp := Response{status, res, float64(num)}
	return resp
}

func (m *Metrix) Abc() Response {
	var res string
	var status int64
	m.X = m.X * m.X
	m.Y = m.Y * m.Y
	if m.X > 49 {
		res = "success"
		status = 200
	} else {
		res = "fail"
		status = 504
	}
	resp := Response{status, res, m.Y * m.X}
	return resp
}

func main() {
	var fac Face
	f := Metrix{1.414, 1.723}
	v := Inter(31)

	fac = &f
	fmt.Println(fac.Abc())
	fac = v

	fmt.Println(fac.Abc())
}
