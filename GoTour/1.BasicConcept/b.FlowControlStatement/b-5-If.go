//If
//Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
package main

import "fmt"

func IfTest(x int) int {
	i, sum := 0, 0
	for i < x {
		if i%2 == 0 {
			sum += i
		}
		if i%2 != 0 {
			sum -= i
		}
		i += 1
	}
	return sum
}

func main() {
	fmt.Println(IfTest(13))
	fmt.Println(IfTest(51))
}
