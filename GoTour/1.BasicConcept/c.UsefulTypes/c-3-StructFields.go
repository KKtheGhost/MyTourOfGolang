//Struct Fields
//Struct fields are accessed using a dot.
package main

import "fmt"

type Macbook struct {
	MacType   string
	SerialNum string
	Year      int
}

func main() {
	Air := Macbook{"Macbook Air", "2102310YJV10500897", 2019}
	Air.Year = 2020
	fmt.Println(Air)
}
