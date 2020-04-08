//Slices of slices
//Slices can contain any type, including other slices.
package main

import (
	"fmt"
	"strings"
)

var chess = [][]string{
	[]string{"0", "0", "0", "0", "0"},
	[]string{"0", "0", "0", "0", "0"},
	[]string{"0", "0", "0", "0", "0"},
	[]string{"0", "0", "0", "0", "0"},
	[]string{"0", "0", "0", "0", "0"},
}

func main() {
	for x := 0; x < len(chess); x++ {
		chess[x][x] = "*"
	}
	for i := 0; i < len(chess); i++ {
		fmt.Printf("%s\n", strings.Join(chess[i], "-"))
	}
}
