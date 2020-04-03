//Switch evaluation order
//Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
//
//(For example,
//
//switch i {
//case 0:
//case f():
//}
//does not call f if i==0.)
//
//Note: Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader.
package main

import (
	"fmt"
)

func HouseLoan(level int) int {
	switch level {
	case 1:
		return 98000
	case 2:
		return 198000
	case 3:
		return 348000
	case 4:
		return 498000
	case 5:
		return 998000
	default:
		return 0
	}
}

func main() {
	var a int
	fmt.Println("你的房子升级到几级了啊")
	fmt.Scanln(&a)
	a = HouseLoan(a)
	switch a {
	case 0:
		fmt.Println("云玩家给我爬")
	default:
		fmt.Printf("哈哈哈哈穷狗你要还%v铃钱，滚去钓鱼吧", a)
	}
}
