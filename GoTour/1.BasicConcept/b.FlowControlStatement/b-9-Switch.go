//Switch
//A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.
//
//Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
package main

import "fmt"

func DaTouCai(price int) {
	fmt.Println("来看看大头菜的价格\n")
	switch a := price; a {
	case 100:
		fmt.Println("价格不行")
	case 200:
		fmt.Println("还有得赚")
	case 400:
		fmt.Println("大头菜赢家！")
	default:
		fmt.Println("土豪今晚来你岛")
	}
}

func main() {
	var a int
	fmt.Println("你岛上大头菜的价格是：")
	fmt.Scanln(&a)
	DaTouCai(a)
}
