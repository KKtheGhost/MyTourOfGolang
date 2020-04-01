//Numeric Constants
//Numeric constants are high-precision values.
//
//An untyped constant takes the type needed by its context.
//
//Try printing needInt(Big) too.
//
//(An int can store at maximum a 64-bit integer, and sometimes less.)
package main

import "fmt"

const (
	LingQian = 1 << 50
	LvCheng  = LingQian >> 48
)

func GetInt(x int) int           { return x*20 + 5 }
func GetFloat(x float64) float64 { return x*0.01 - 12 }

func main() {
	fmt.Println(GetFloat(LingQian))
	fmt.Println(GetFloat(LvCheng))
	fmt.Println(GetInt(LvCheng))
}
