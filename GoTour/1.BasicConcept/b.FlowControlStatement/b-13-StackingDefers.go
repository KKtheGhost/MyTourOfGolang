//Stacking defers
//Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
//
//To learn more about defer statements read this blog post.
package main

import "fmt"

type Fish struct {
	Name   string
	Price  int
	Locate string
	Get    bool
}

var GetFish = Fish{"WhaleShark", 15000, "Ocean", true}

func main() {
	defer fmt.Printf("这种东西名字叫%v\n", GetFish.Name)
	defer fmt.Printf("你有没有钓到这种东西呢？%v\n", GetFish.Get)

	fmt.Printf("这种鱼的价格是%v,可以钓到的位置是%v\n", GetFish.Price, GetFish.Locate)
}
