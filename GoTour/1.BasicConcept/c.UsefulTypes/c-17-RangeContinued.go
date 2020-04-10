//Range continued
//You can skip the index or value by assigning to _.
//
//for i, _ := range pow
//for _, value := range pow
//If you only want the index, you can omit the second variable.
//
//for i := range pow
//这里解释一个很常用的下划线的作用，本质就是ignore一个值。在range这个功能里可以是index，也可以是value
package main

import "fmt"

func GetValue(dic []int) []int {
	for i, _ := range dic {
		dic[i] = i << uint(i)
	}
	return dic
}

func GetIndex(dic []int) []int {
	idic := make([]int, len(dic))
	for i, _ := range dic {
		idic[i] = i
	}
	return idic
}

func main() {
	array := make([]int, 7)
	fmt.Println(GetIndex(array))
	fmt.Println(GetValue(array))
}
