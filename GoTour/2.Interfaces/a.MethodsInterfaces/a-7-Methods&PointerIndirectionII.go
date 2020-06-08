//Methods and pointer indirection (2)
//The equivalent thing happens in the reverse direction.
//
//Functions that take a value argument must take a value of that specific type:
//
//var v Vertex
//fmt.Println(AbsFunc(v))  // OK
//fmt.Println(AbsFunc(&v)) // Compile error!
//while methods with value receivers take either a value or a pointer as the receiver when they are called:
//
//var v Vertex
//fmt.Println(v.Abs()) // OK
//p := &v
//fmt.Println(p.Abs()) // OK
//In this case, the method call p.Abs() is interpreted as (*p).Abs().
//接受一个值作为参数的函数必须接受一个指定类型的值, 以值为接收者的方法被调用时，接收者既能为值又能为指针
//方法调用 p.Abs() 会被解释为 (*p).Abs()
package main

import "fmt"

type Metrix struct {
	len float64
	wid float64
}

func Sqr(m Metrix, x float64) float64 {
	m.len = m.len * x
	m.wid = m.wid * x
	return m.len - m.wid
}

func (m Metrix) SqrII(x float64) float64 {
	m.len = m.len * x
	m.wid = m.wid * x
	return m.len - m.wid
}

func main() {
	aa := Metrix{5, 6}
	pp := &Metrix{6, 5}
	aa.SqrII(2)
	pp.SqrII(3)

	fmt.Println(Sqr(aa, 5), Sqr(*pp, 5))
}
