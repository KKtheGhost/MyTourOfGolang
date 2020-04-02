//If and else
//Variables declared inside an if short statement are also available inside any of the else blocks.
//
//(Both calls to pow return their results before the call to fmt.Println in main begins.)
package main

import "fmt"

func ShallISell(a, b, c, price int) bool {
	if total := a + b + c; price*3 < total {
		return true
	} else {
		fmt.Printf("累加总和%v还不如三倍初值%v，给我爬\n", total, price)
	}
	return false
}

func main() {
	fmt.Println(
		ShallISell(15, 16, 17, 18),
		ShallISell(15, 16, 17, 12),
	)
}
