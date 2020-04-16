//Maps
//A map maps keys to values.
//
//The zero value of a map is nil. A nil map has no keys, nor can keys be added.
//
//The make function returns a map of the given type, initialized and ready for use.
package main

import "fmt"

type PokemonVI struct {
	HP, Atk, Def, SpA, SpD, Spd int
}

var Pokemon map[string]PokemonVI

func main() {
	Pokemon = make(map[string]PokemonVI)
	Pokemon["Serebii"] = PokemonVI{
		100, 100, 100, 100, 100, 100,
	}
	fmt.Print("Serebii的个体值为", Pokemon["Serebii"])
}
