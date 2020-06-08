//Methods and pointer indirection
//Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:
//
//var v Vertex
//ScaleFunc(v, 5)  // Compile error!
//ScaleFunc(&v, 5) // OK
//while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
//
//var v Vertex
//v.Scale(5)  // OK
//p := &v
//p.Scale(10) // OK
//For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
//带指针参数的函数必须接受一个指针, 以指针为接收者的方法被调用时，接收者既能为值又能为指针
//Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。
package main

import (
	"fmt"
)

type Metrix struct {
	len float64
	wid float64
}

func Sqr(m *Metrix, x float64) {
	m.len = m.len * x
	m.wid = m.wid * x
}

func (m *Metrix) SqrII(x float64) {
	m.len = m.len * x
	m.wid = m.wid * x
}

func main() {
	aa := Metrix{5, 6}
	pp := &Metrix{6, 5}
	aa.SqrII(2)
	pp.SqrII(3)
	Sqr(&aa, 5)
	Sqr(pp, 5)
	fmt.Println(aa, pp)
}
