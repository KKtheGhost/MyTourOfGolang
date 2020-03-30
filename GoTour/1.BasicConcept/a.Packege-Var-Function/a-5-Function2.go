//函数（续）
//当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
//
//在本例中，
//
//x int, y int
//被缩写为
//
//x, y int
package main

import (
	"fmt"
	"math"
)

func cmplxii(a, b, c float64) float64 {
	return a*b + math.Sqrt(b) - (c+5)*a
}

func main() {
	fmt.Printf("复杂计算的结果是:%v", cmplxii(7, 12, 67))
}
