//Zero values
//Variables declared without an explicit initial value are given their zero value.
//
//The zero value is:
//
//0 for numeric types,
//false for the boolean type, and
//"" (the empty string) for strings.
package main

import "fmt"

func main() {
	var aa int
	var bb string
	var cc bool
	var dd float64
	fmt.Printf("%v %q %v %v\n", aa, bb, cc, dd)
}
