//Function closures
//Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
//
//For example, the adder function returns a closure. Each closure is bound to its own sum variable.
package main

import "fmt"

func adder() func(int) int {
	//此处的意思其实是，返回的是一个函数的【值】，而这个函数的格式是func(int) int
	sum := 0
	return func(x int) int {
		//此处对函数func进行了定义
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i > -10; i-- {
		fmt.Println(
			pos(i),
			neg(-2*i+i),
		)
	}
}
