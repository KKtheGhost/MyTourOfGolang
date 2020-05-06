package main

import (
	"fmt"
	"local/AesCalculate"
)

func main() {
	accessKey, accessSecret, AliyunDecryptKey := "LTAI4GFtXeJ1kCHgJh899DNf", "rfPNccrRo4T3zSA4pJho30CNqSzAjl", "sMok6A63cUigEE6P"
	AliyunTokenKeyEncrypt, AliyunTokenSecretEncrypt := AesCalculate.AesEncrypt(accessKey, AliyunDecryptKey), AesCalculate.AesEncrypt(accessSecret, AliyunDecryptKey)
	fmt.Printf("%v\n", AliyunTokenKeyEncrypt)
	fmt.Printf("%v\n", AliyunTokenSecretEncrypt)
}
