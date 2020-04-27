package main

//本工具目的是,面对日益增多的阿里云账号,通过工具自动每月月初获取发票,从而大幅度减少无意义的重复劳动.

//需要导入的无非就是aliyunbssopenapi
//以及处理json的开源插件gojsonq
//fmt仅用于调试,之后可以去掉.
import (
	"fmt"
	"local/AesCalculate"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

//核心就是模块化.
//location无所谓,只代表Client的地域,不影响Invoice的获取.
//主要后面的accessKey和accessSecret需要通过aliyun-config.yml的配置来获取.
//以及AliyunInvoiceInfo.BillCycle也需要处理一下.虽然多数阿里云账号没有MARKET消费,但是香肠派对和主账号是有的
func main() {
	AliyunDecryptKey := "sMok6A63cUigEE6P"
	accessKey, accessSecret := "", ""
	AliyunTokenKeyEncrypt, AliyunTokenSecretEncrypt := AesCalculate.AesEncrypt(accessKey, AliyunDecryptKey), AesCalculate.AesEncrypt(accessSecret, AliyunDecryptKey)
	fmt.Printf("AliyunTokenKeyEncrypt is %v\n", AliyunTokenKeyEncrypt)
	fmt.Printf("AliyunTokenSecretEncrypt is %v\n", AliyunTokenSecretEncrypt)
	AliyunInvoiceClient, AliyunClientErr := bssopenapi.NewClientWithAccessKey("cn-shanghai", accessKey, accessSecret)
	if AliyunClientErr != nil {
		fmt.Print(AliyunClientErr.Error())
	}
	//构建APIclient用于后续获取源格式
	//AliyunInvoiceClient, err := bssopenapi.NewClientWithAccessKey("cn-shanghai", "accessKey", "accessSecret")
	AliyunInvoiceInfo := bssopenapi.CreateQueryEvaluateListRequest()
	AliyunInvoiceInfo.Scheme = "https"
	AliyunInvoiceInfo.BillCycle = "ALIYUN"

	AliyunInvoiceResponse, AliyunInvoiceErr := AliyunInvoiceClient.QueryEvaluateList(AliyunInvoiceInfo)
	if AliyunInvoiceErr != nil {
		fmt.Print(AliyunInvoiceErr.Error())
	}
	fmt.Printf("%#v\n", AliyunInvoiceResponse.Data.TotalInvoiceAmount)
	//size := gojsonq.New().JSONString(AliyunInvoiceResponse).From("Data.PageSize")
	//fmt.Println(size.(string))
}
