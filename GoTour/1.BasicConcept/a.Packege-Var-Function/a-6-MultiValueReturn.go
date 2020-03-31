//多值返回
//函数可以返回任意数量的返回值。
//
//swap 函数返回了两个字符串。
package main

import (
	"fmt"
)

func muti(a, b string) (string, string) {
	return b, a
}

func main() {
	a1, a2 := muti("黑社会啊", "你是真没见过")
	b1, b2 := muti("比划比划", "小兔崽子我们")
	fmt.Printf("%v %v\n", a1, a2)
	fmt.Println(b1, b2)
}
