//
//Function values
//Functions are values too. They can be passed around just like other values.
//
//Function values may be used as function arguments and return values.
//
package main

import (
	"fmt"
	"math"
)

//这里这么解释，fn被定义为一个函数，这个函数呢需要的格式是，输入一个(float64, float64)的数组，然后返回一个float64的值。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	//此处为(5*5 + 12*12)^(1/2)=13
	fmt.Println(compute(hypot))
	//此处因为hypot的入参被compute强制定义为(3, 4), 所以结果是(3*3 + 4*4)^(1/2)=5
	fmt.Println(compute(math.Pow))
	//math.Pow入参也是(float64, float64)格式，合法。入参强制为(3, 4)，输出结果为3**4=81
}
