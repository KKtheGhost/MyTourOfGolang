//Methods continued
//You can declare a method on non-struct types, too.
//
//In this example we see a numeric type MyFloat with an Abs method.
//
//You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
//就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法
package main

import "fmt"

type Body struct {
	Weight float64
	Height float64
}

func (man Body) BMI() float64 {
	bmi := man.Weight / (man.Height * man.Height)
	return bmi
}

func main() {
	ManA := Body{98, 1.83}
	fmt.Printf("BMI of this man is %#.2f\n", ManA.BMI())
}
