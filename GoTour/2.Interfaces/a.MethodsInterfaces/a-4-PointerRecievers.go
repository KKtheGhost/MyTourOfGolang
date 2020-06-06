//Pointer receivers
//You can declare methods with pointer receivers.
//
//This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
//
//For example, the Scale method here is defined on *Vertex.
//
//Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
//
//Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.
//
//With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

//若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。（对于函数的其它参数也是如此。）Scale 方法必须用指针接受者来更改 main 函数中声明的 Vertex 的值。
package main

import "fmt"

type Metrix struct {
	a int
	b int
}

func (m *Metrix) Multi(x int) { //result is 135, cause the origin value has been modified(edit pointer)
	//func (m Metrix) Multi(x int) {		//result is 27
	m.a = m.a * x
	m.b = m.b * x
}

func (m Metrix) Comb() int {
	return m.a*10 + m.b
}

func main() {
	m := Metrix{2, 7}
	m.Multi(5)
	fmt.Printf("现在的结果是%#v.\n", m.Comb())
}
