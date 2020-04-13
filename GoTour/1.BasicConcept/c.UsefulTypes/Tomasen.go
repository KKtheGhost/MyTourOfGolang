package main

import "fmt"

func main() {
	var price float64 = 74451.41
	fmt.Println("仙侠道web：5%\n神仙道ios：41%\n天天打波利：26%\n公司分摊：28%\n")
	xxd := price * 0.05
	sxd := price * 0.41
	ttd := price * 0.26
	gsf := price * 0.28

	fmt.Printf("结合本月的总成本，具体分摊金额为：\n仙侠道：%.2f元\n", xxd)
	fmt.Printf("神仙道：%.2f元\n", sxd)
	fmt.Printf("天天打波利：%.2f元\n", ttd)
	fmt.Printf("公司分摊：%.2f元\n", gsf)
}
