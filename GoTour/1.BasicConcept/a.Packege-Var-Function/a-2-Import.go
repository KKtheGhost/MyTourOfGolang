package main

import (
	"fmt"
	"math"
)

var size float64 = 81

func main() {
	fmt.Printf("the length of a %v-size square is %v", size, math.Sqrt(size))
}
