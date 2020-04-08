//Slice literals
//A slice literal is like an array literal without the length.
//
//This is an array literal:
//
//[3]bool{true, true, false}
//And this creates the same array as above, then builds a slice that references it:
//
//[]bool{true, true, false}
package main

import "fmt"

func main() {
	array1 := [4]string{"anal", "amputee", "futanari", "whore"}
	array2 := []string{"mutiple breasts", "gender bender", "pantyhose", "ahegao"}
	fmt.Println(array1, array2)

	love := []struct {
		a int
		b string
	}{
		{13, "loli"},
		{14, "JK"},
		{20, "Daigakusei"},
		{24, "OL"},
	}
	fmt.Println(love)
}
