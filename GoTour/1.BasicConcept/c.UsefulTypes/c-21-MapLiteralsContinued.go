//Map literals continued
//If the top-level type is just a type name, you can omit it from the elements of the literal.
//若顶级类型只是一个类型名，你可以在文法的元素中省略它。
package main

import "fmt"

type PokemonVI struct {
	HP, Atk, Def, SpA, SpD, Spd int
}

//所有var赋值结构体的时候，任何一个被赋值单位必须有逗号结尾，需要注意
var Pokemon = map[string]PokemonVI{
	"Serebii": {
		100, 100, 100, 100, 100, 100,
	},
	"Dracovish": {
		90, 90, 100, 70, 80, 75,
	},
}

func main() {
	fmt.Println(Pokemon)
}
