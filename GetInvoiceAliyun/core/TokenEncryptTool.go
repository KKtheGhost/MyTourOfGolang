package main

import (
	"fmt"
	"local/AesCalculate"
)

func main() {
	accessKey, accessSecret, AliyunDecryptKey := "", "", "sMok6A63cUigEE6P"
	AliyunTokenKeyEncrypt, AliyunTokenSecretEncrypt := AesCalculate.AesEncrypt(accessKey, AliyunDecryptKey), AesCalculate.AesEncrypt(accessSecret, AliyunDecryptKey)
	fmt.Printf("%v\n", AliyunTokenKeyEncrypt)
	fmt.Printf("%v\n", AliyunTokenSecretEncrypt)
}
