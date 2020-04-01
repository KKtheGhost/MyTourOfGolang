//Short variable declarations
//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
//
//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
package main

import "fmt"

func main() {
	a, b, c := 1, 2, "集合啦！"
	d, e := "动物森友会", true
	fmt.Println(a, b, c, d, e)
}
