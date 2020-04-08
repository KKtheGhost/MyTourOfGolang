//Slices are like references to arrays
//A slice does not store any data, it just describes a section of an underlying array.
//
//Changing the elements of a slice modifies the corresponding elements of its underlying array.
//
//Other slices that share the same underlying array will see those changes.
package main

import "fmt"

var pkm = [5]string{"Mew", "Mewtwo", "Mashato", "Serebii", "Diancy"}

func main() {
	ft := pkm[0:2]
	bk := pkm[3:5]
	fmt.Println(ft, bk)

	bk[0] = "Jirachi"
	fmt.Println(bk)
}
