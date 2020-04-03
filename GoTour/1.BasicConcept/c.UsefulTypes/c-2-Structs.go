//Structs
//A struct is a collection of fields.
package main

import "fmt"

type XDmember struct {
	Name   string
	Age    int
	Level  int
	Dept   string
	Active bool
}

var FangWei = XDmember{"方苇", 26, 1, "CoreSystem", true}

func main() {
	var p = &FangWei
	fmt.Println(*p) //*本身只是对指针代表的内存位置的一次解耦，相当于还原出了那个内存地址所代表的值。所以可以作为一个整体直接输出
	//但是*p本身并不是结构体，所以无法使用类似*p.a这样的格式去获取结构体中的值。需要注意。
}
