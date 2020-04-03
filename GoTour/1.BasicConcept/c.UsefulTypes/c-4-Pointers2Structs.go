//Pointers to structs
//Struct fields can be accessed through a struct pointer.
//
//To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
package main

import "fmt"

//前面提到过*p.a这样的结构是不可以用的，那么这里是一个补充，就是关于p.a这样的格式的形式。其实是一种简写
//本质其实是(*p).a的写法，也就是说用括号括起来的解耦指针也可以用来实现结构体的取值方法。

type dm struct {
	aa int
	bb int
	cc int
}

func main() {
	d1 := dm{3, 4, 5}
	d2 := dm{0, 0, 0}
	p1, p2 := &d1, &d2
	p2.aa = p1.aa * p1.aa
	(*p2).bb = p1.bb * p1.bb
	p2.cc = (*p1).cc * (*p1).cc
	fmt.Println(*p2)
	fmt.Printf("数组1的值之和为%v", (*p1).aa+(*p1).bb+(*p1).cc)
}
