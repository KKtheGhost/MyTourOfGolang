//Map literals
//Map literals are like struct literals, but the keys are required.
package main

import "fmt"

type PokemonVI struct {
	HP, Atk, Def, SpA, SpD, Spd int
}

var Pokemon = map[string]PokemonVI{
	"Serebii": PokemonVI{
		100, 100, 100, 100, 100, 100,
	},
	"Dracovish": PokemonVI{
		90, 90, 100, 70, 80, 75,
	},
}

func main() {
	fmt.Println(Pokemon)
}
