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
	2: "ERROR! GAE002: Error Building Client - Please check out network status and browser settings.",
	3: "ERROR! GAE003: Error Invoice Request - Please contact Aliyun support to figure out causes.",
}

type AliyunAddrMetrix struct {
	AliyunAddrProjName    string
	AliyunAddressID       int64
	AliyunAddressee       string
	AliyunDeliveryAddress string
}

var EncryptedMetrixLen int = len(CsvFilter.CsvConvert())

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

//获取单个阿里云账号的Address
func GetInvoiceAddress(AliyunID int) (response *bssopenapi.QueryCustomerAddressListResponse) {
	//入参是AliyunID,通过ID,在AliyunTokenSet这个结构体中去找正确的key和secret.
	accessKey, accessSecret := TokenDecrypt(CsvFilter.CsvConvert()[AliyunID])
	AliyunInvoiceClient, AliyunClientErr := bssopenapi.NewClientWithAccessKey("cn-shanghai", accessKey, accessSecret)
	if AliyunClientErr != nil {
		fmt.Println(GetAddressErrorCode[2])
	}
	AliyunInvoiceRequest := bssopenapi.CreateQueryCustomerAddressListRequest()
	AliyunInvoiceRequest.Scheme = "https"
	AliyunInvoiceResponse, AliyunInvoiceErr := AliyunInvoiceClient.QueryCustomerAddressList(AliyunInvoiceRequest)
	if AliyunInvoiceErr != nil {
		fmt.Println(GetAddressErrorCode[3])
	}
	return AliyunInvoiceResponse
	//返回的是一个合集，需要筛选生成结构体
}

func GenerateAddrMetrix() map[int]AliyunAddrMetrix {
	AddrMetrixRes := make(map[int]AliyunAddrMetrix)
	for AliyunID := 1; AliyunID <= EncryptedMetrixLen; AliyunID++ {
		AddrList := GetInvoiceAddress(AliyunID).Data.CustomerInvoiceAddressList.CustomerInvoiceAddress
		AddrNum := len(AddrList)
		AddrProjName := CsvFilter.CsvConvert()[AliyunID].AliyunNicknam
		for AddrID := 0; AddrID < AddrNum; AddrID++ {
			if AddrList[AddrID].Addressee == "方苇" {
				AddrMetrix := AliyunAddrMetrix{AddrProjName, AddrList[AddrID].Id, AddrList[AddrID].Addressee, AddrList[AddrID].DeliveryAddress}
				AddrMetrixRes[AliyunID] = AddrMetrix
			} else {
				continue
			}
		}
	}
	return AddrMetrixRes
}

//打包TokenDecrypt和GetInvoiceAddress到循环中，输出一个有效的的结构体，包含ID，UserName，项目名。

func main() {
	//fmt.Printf("%#v\n", GenerateAddrMetrix())
	fmt.Println(GenerateAddrMetrix())
}
