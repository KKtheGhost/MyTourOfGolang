package main

import (
	"fmt"
	"math"
)

func TellMe(a, b float64) (x, y int) {
	x = int(math.Floor(a*b)) * 5
	y = int(math.Pow(a, b))
	return
}

func main() {
	fmt.Println(TellMe(3, 4))
}
