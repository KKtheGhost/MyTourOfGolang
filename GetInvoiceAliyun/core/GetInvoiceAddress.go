package main

import (
	"fmt"
	"local/AesCalculate"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

type AliyunTokenSet struct {
	TokenIndex               int
	AliyunNicknam            string
	AliyunDecryptKey         string
	AliyunTokenKeyEncrypt    string
	AliyunTokenSecretEncrypt string
}

//解密Token的func
func DecryptToken(AliyunTokenKeyEncrypt, AliyunTokenSecretEncrypt string) (accessKey, accessSecret string) {

}

func GetInvoiceAddress(AliyunID string) (response *QueryCustomerAddressListResponse) {
	//入参是AliyunID,通过ID,在AliyunTokenSet这个结构体中去找正确的key和secret.
	AliyunInvoiceClient, AliyunClientErr := bssopenapi.NewClientWithAccessKey("cn-shanghai", accessKey, accessSecret)
	if AliyunClientErr != nil {
		fmt.Print(AliyunClientErr.Error())
	}
	AliyunInvoiceRequest := bssopenapi.CreateQueryCustomerAddressListRequest()
	AliyunInvoiceRequest.Scheme = "https"

	AliyunInvoiceResponse, AliyunInvoiceErr := AliyunInvoiceClient.QueryCustomerAddressList(AliyunInvoiceRequest)
	if AliyunInvoiceErr != nil {
		fmt.Print(AliyunInvoiceErr.Error())
	}
	return AliyunInvoiceResponse
}

func main() {
	AliyunID := 1
	AliyunDecryptKey := "sMok6A63cUigEE6P"
	accessKey, accessSecret := "", ""
	AliyunTokenKeyEncrypt, AliyunTokenSecretEncrypt := AesCalculate.AesEncrypt(accessKey, AliyunDecryptKey), AesCalculate.AesEncrypt(accessSecret, AliyunDecryptKey)
	fmt.Printf("AliyunTokenKeyEncrypt is %v\n", AliyunTokenKeyEncrypt)
	fmt.Printf("AliyunTokenSecretEncrypt is %v\n", AliyunTokenSecretEncrypt)
	fmt.Printf("%#v\n", GetInvoiceAddress(AliyunID).Data.CustomerInvoiceAddressList)
}
