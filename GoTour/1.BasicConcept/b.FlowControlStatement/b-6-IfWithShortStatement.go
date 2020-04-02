//If with a short statement
//Like for, the if statement can start with a short statement to execute before the condition.
//
//Variables declared by the statement are only in scope until the end of the if.
//
//(Try using v in the last return statement.)
package main

import (
	"fmt"
	"math"
)

func ToSell(a, b, PriceIn float64) bool {
	if avg := math.Sqrt(a * b); PriceIn < avg {
		return true
	}
	return false
}

func main() {
	fmt.Println(
		ToSell(87, 94, 95),
		ToSell(88, 120, 110),
	)
}
