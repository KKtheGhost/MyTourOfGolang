package main

import (
	"fmt"
	"io/ioutil"
	"local/AesCalculate"
	"os"

	"local/CsvFilter"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

//本功能的异常抛出列表。
var GetAddressErrorCode = map[int]string{
	1: "ERROR! GAE001: Missing Csv FIle - Cannot Read the DecryptCode file or the file is not existing.",
	2: "ERROR! GAE002: Error Building Client - Please check out network status and browser settings.",
	3: "ERROR! GAE003: Error Invoice Request - Please contact Aliyun support to figure out causes.",
}

//本功能的输出结构体（用户 地址 地址id）对照表。
type AliyunAddrMetrix struct {
	AliyunAddrProjName    string
	AliyunAddressID       int64
	AliyunAddressee       string
	AliyunDeliveryAddress string
}

//获取Csv输出长度，用于后续循环。
var EncryptedMetrixLen int = len(CsvFilter.CsvConvert())

//解密单个AliyunTokenSet的密钥
func TokenDecrypt(EncryptedMetrix CsvFilter.AliyunTokenSet) (accessKey string, accessSecret string) {
	DecryptKey, DecryptKeyErr := ioutil.ReadFile("/etc/GetAliyunInvoice/DecryptCode.key")
	if DecryptKeyErr != nil {
		fmt.Println(GetAddressErrorCode[1])
		os.Exit(0)
	}
	//解密token，获取原始的tokenKey和tokenSecret
	OriginAccessKey := AesCalculate.AesDecrypt(EncryptedMetrix.AliyunTokenKeyEncrypt, string(DecryptKey))
	OriginAccessSecret := AesCalculate.AesDecrypt(EncryptedMetrix.AliyunTokenSecretEncrypt, string(DecryptKey))
	return OriginAccessKey, OriginAccessSecret
}

//利用aliyun-go-sdk获取原始的阿里云发票申请地址信息，用于后续数据处理。
func GetInvoiceAddress(AliyunID int) (response *bssopenapi.QueryCustomerAddressListResponse) {
	//入参是AliyunID,通过ID,在AliyunTokenSet这个结构体中去找正确的key和secret.
	//AliyunID的好处是，从CsvFilter开始，就保持了index的稳定和一致性，便于后期排查问题。此处的AliyunID从1开始。在维护时需要注意。
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

//打包TokenDecrypt和GetInvoiceAddress到循环中，输出一个有效的的结构体，包含项目名, 地址ID，申请人姓名，具体地址。
func GenerateAddrMetrix() map[int]AliyunAddrMetrix {
	AddrMetrixRes := make(map[int]AliyunAddrMetrix)
	//此处上界为大于等于，因为Aliyun起始是1
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

//main仅作为测试用，最后这个要打成一个包的
func main() {
	fmt.Println(GenerateAddrMetrix())
}
