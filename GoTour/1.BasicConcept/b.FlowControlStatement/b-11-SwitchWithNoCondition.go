//Switch with no condition
//Switch without a condition is the same as switch true.
//
//This construct can be a clean way to write long if-then-else chains.
package main

import "fmt"

func WhichFish(price int) {
	switch {
	case price < 100:
		fmt.Println("蝌蚪")
	case price < 500:
		fmt.Println("溪哥")
	case price < 1000:
		fmt.Println("黑鲈鱼")
	case price < 2000:
		fmt.Println("鳄龟")
	case price < 3000:
		fmt.Println("花羔红点鲑")
	case price < 5000:
		fmt.Println("恩氏多鳍鱼")
	case price < 10000:
		fmt.Println("金鳟")
	default:
		fmt.Println("远东哲罗鱼")
	}
}

func main() {
	var a int
	fmt.Println("今天钓鱼卖了多少")
	fmt.Scanln(&a)
	WhichFish(a)
}
