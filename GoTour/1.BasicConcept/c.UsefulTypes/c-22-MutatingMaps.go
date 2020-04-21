//Mutating Maps
//Insert or update an element in map m:
//
//m[key] = elem
//Retrieve an element:
//
//elem = m[key]
//Delete an element:
//
//delete(m, key)
//Test that a key is present with a two-value assignment:
//
//elem, ok = m[key]
//If key is in m, ok is true. If not, ok is false.
//
//If key is not in the map, then elem is the zero value for the map's element type.
//
//Note: If elem or ok have not yet been declared you could use a short declaration form:
//
//elem, ok := m[key]
package main

import "fmt"

func main() {
	Pokemon := make(map[string]int)

	Pokemon["Pikachu"] = 25
	fmt.Println("The value:", Pokemon["Pikachu"])

	Pokemon["Volcarona"] = 637
	fmt.Println("The value:", Pokemon["Volcarona"])

	delete(Pokemon, "Pikachu")
	fmt.Println("The value:", Pokemon["Pikachu"])

	poke, ok := Pokemon["Pikachu"]
	fmt.Println("The value:", poke, "Present?", ok)
}
