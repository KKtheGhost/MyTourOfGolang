//For
//Go has only one looping construct, the for loop.
//
//The basic for loop has three components separated by semicolons:
//
//the init statement: executed before the first iteration
//the condition expression: evaluated before every iteration
//the post statement: executed at the end of every iteration
//The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.
//
//The loop will stop iterating once the boolean condition evaluates to false.
//
//Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
package main

import "fmt"

func Factorial(x int) int {
	var sum = 1
	for i := 1; i <= x; i++ {
		sum *= i
	}
	return sum
}

func main() {
	a, b := 5, 6
	fmt.Println(Factorial(a), Factorial(b))
}
