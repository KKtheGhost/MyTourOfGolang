//Interface values
//Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
//
//(value, type)
//An interface value holds a value of a specific underlying concrete type.
//
//Calling a method on an interface value executes the method of the same name on its underlying type.
//本质就是接口可以理解成就是一个value，可以用于传递，就这么简单。
package main

import "fmt"

type Juzi struct {
	S string
}

type Num float64

type M interface {
	Abc()
}

func (j Juzi) Abc() {
	fmt.Println(j.S)
}

func (n Num) Abc() {
	fmt.Println(n)
}

func Type(m M) {
	fmt.Printf("%v, %T\n", m, m)
}

func main() {
	var m M = &Juzi{"那是真的牛逼兄弟"}
	Type(m)
	m.Abc()

	m = Num(5.12342354235)
	Type(m)
	m.Abc()
}
