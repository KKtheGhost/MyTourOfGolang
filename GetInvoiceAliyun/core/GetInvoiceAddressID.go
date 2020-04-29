package main

import (
	"fmt"
	"io/ioutil"
	"local/AesCalculate"
	"os"

	"local/CsvFilter"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

var GetAddressErrorCode = map[int]string{
	1: "ERROR! GAE001: Missing Csv FIle - Cannot Read the DecryptCode file or the file is not existing.",
}

//解密单个AliyunTokenSet的密钥
func TokenDecrypt(EncryptedMetrix CsvFilter.AliyunTokenSet) (accessKey string, accessSecret string) {
	DecryptKey, DecryptKeyErr := ioutil.ReadFile("/etc/GetAliyunInvoice/DecryptCode.key")
	//fmt.Printf("%v\n", string(DecryptKey))
	if DecryptKeyErr != nil {
		fmt.Println(GetAddressErrorCode[1])
		os.Exit(0)
	}
	OriginAccessKey := AesCalculate.AesDecrypt(EncryptedMetrix.AliyunTokenKeyEncrypt, string(DecryptKey))
	//fmt.Printf("%v\n", OriginAccessKey)
	OriginAccessSecret := AesCalculate.AesDecrypt(EncryptedMetrix.AliyunTokenSecretEncrypt, string(DecryptKey))
	//fmt.Printf("%v\n", OriginAccessSecret)
	return OriginAccessKey, OriginAccessSecret
}

var AliyunID int = 3

//获取单个阿里云账号的Address
func GetInvoiceAddress(AliyunID int) (response *bssopenapi.QueryCustomerAddressListResponse) {
	//入参是AliyunID,通过ID,在AliyunTokenSet这个结构体中去找正确的key和secret.
	accessKey, accessSecret := TokenDecrypt(CsvFilter.CsvConvert()[AliyunID])
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
	fmt.Printf("%#v\n", GetInvoiceAddress(AliyunID).Data.CustomerInvoiceAddressList.CustomerInvoiceAddress[3].Id) //.Data.CustomerInvoiceAddressList.CustomerInvoiceAddress)
	fmt.Printf("%#v\n", GetInvoiceAddress(AliyunID).Data.CustomerInvoiceAddressList.CustomerInvoiceAddress[3].Addressee)
	fmt.Printf("%#v\n", GetInvoiceAddress(AliyunID).Data.CustomerInvoiceAddressList.CustomerInvoiceAddress[3].DeliveryAddress)
}
