//Constants
//Constants are declared like variables, but with the const keyword.
//
//Constants can be character, string, boolean, or numeric values.
//
//Constants cannot be declared using the := syntax.
package main

import "fmt"

const GameName, GameNameEN string = "集合啦！动物森友会", "Animal Crossing! New Horizons"

func main() {
	const price int = 325
	fmt.Println("现在最火的游戏是", GameName, "\n")
	fmt.Println("The most popular game in 2020 is ", GameNameEN, "\n")
	fmt.Println("现在的价格是：", price)
	fmt.Println("Today's price is ", price)
}
