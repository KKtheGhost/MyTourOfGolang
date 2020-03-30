//包
//每个 Go 程序都是由包构成的。
//
//程序从 main 包开始运行。
//
//本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。
//
//按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。
//
//*注意：* 此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。 （要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed。 练习场中的时间为常量，因此你需要用其它的值作为种子数。）
//
package main

import (
	"fmt"
)

func main() {
	num := 5
	fmt.Printf("My Env Seed is %+v", num)
}
