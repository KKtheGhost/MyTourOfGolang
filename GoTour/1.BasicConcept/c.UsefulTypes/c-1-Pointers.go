//Pointers
//Go has pointers. A pointer holds the memory address of a value.
//
//The type *T is a pointer to a T value. Its zero value is nil.
//
//var p *int
//The & operator generates a pointer to its operand.
//
//i := 42
//p = &i
//The * operator denotes the pointer's underlying value.
//
//fmt.Println(*p) // read i through the pointer p
//*p = 21         // set i through the pointer p
//This is known as "dereferencing" or "indirecting".
//
//Unlike C, Go has no pointer arithmetic.
package main

import "fmt"

type metrix struct {
	a int
	b int
	c int
	d int
}

func main() {
	var data1 = metrix{12, 55, 23, 81}
	var data2 = metrix{36, 20, 4, 37}
	var p1 = &data1
	var p2 = &data2
	fmt.Printf("数组一为：%v\n", *p1)
	fmt.Printf("数组二为：%v\n", *p2)
	fmt.Printf("数组一1,2相加为%v\n", p1.a+p1.b)
	fmt.Printf("数组一全部元素乘积为：%v\n", p1.a*p1.b*p1.c*p1.d)
	fmt.Printf("两个数组最后一个元素的差为：%v\n", p1.a-p2.a)
}
