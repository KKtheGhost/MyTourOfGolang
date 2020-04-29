package main

//本工具目的是,面对日益增多的阿里云账号,通过工具自动每月月初获取发票,从而大幅度减少无意义的重复劳动.

//需要导入的无非就是aliyunbssopenapi
//以及处理json的开源插件gojsonq
//fmt仅用于调试,之后可以去掉.
import (
	"fmt"
	"io/ioutil"
	"local/AesCalculate"
	"local/CsvFilter"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

var GetAmountErrorCode = map[int]string{
	1: "ERROR! GAE001: Missing Csv FIle - Cannot Read the DecryptCode file or the file is not existing.",
	2: "ERROR! GAE002: Error Building Client - Please check out network status and browser settings.",
	3: "ERROR! GAE003: Error Invoice Request - Please contact Aliyun support to figure out causes.",
}

type AliyunAmountMetrix struct {
	AliyunProjName      string
	AliyunInvoiceAmount int64
}

var EncryptedMetrixLen int = len(CsvFilter.CsvConvert())

func TokenDecrypt(EncryptedMetrix CsvFilter.AliyunTokenSet) (accessKey string, accessSecret string) {
	DecryptKey, DecryptKeyErr := ioutil.ReadFile("/etc/GetAliyunInvoice/DecryptCode.key")
	if DecryptKeyErr != nil {
		fmt.Println(GetAmountErrorCode[1])
		os.Exit(0)
	}
	//解密token，获取原始的tokenKey和tokenSecret
	OriginAccessKey := AesCalculate.AesDecrypt(EncryptedMetrix.AliyunTokenKeyEncrypt, string(DecryptKey))
	OriginAccessSecret := AesCalculate.AesDecrypt(EncryptedMetrix.AliyunTokenSecretEncrypt, string(DecryptKey))
	return OriginAccessKey, OriginAccessSecret
}

func GetInvoiceAmount(AliyunID int) (response *bssopenapi.QueryEvaluateListResponse) {
	accessKey, accessSecret := TokenDecrypt(CsvFilter.CsvConvert()[1])
	AliyunInvoiceClient, AliyunClientErr := bssopenapi.NewClientWithAccessKey("cn-shanghai", accessKey, accessSecret)
	if AliyunClientErr != nil {
		fmt.Print(AliyunClientErr.Error())
	}
	AliyunInvoiceInfo := bssopenapi.CreateQueryEvaluateListRequest()
	AliyunInvoiceInfo.Scheme = "https"
	AliyunInvoiceInfo.BillCycle = "ALIYUN"
	AliyunInvoiceResponse, AliyunInvoiceErr := AliyunInvoiceClient.QueryEvaluateList(AliyunInvoiceInfo)
	if AliyunInvoiceErr != nil {
		fmt.Print(AliyunInvoiceErr.Error())
	}
	return AliyunInvoiceResponse
	//fmt.Printf("%#v\n", AliyunInvoiceResponse.Data.TotalInvoiceAmount)
}

func GenerateAmountMetrix() map[int]AliyunAmountMetrix {
	AmountMetrixRes := make(map[int]AliyunAmountMetrix)
	//此处上界为大于等于，因为Aliyun起始是1
	for AliyunID := 1; AliyunID <= EncryptedMetrixLen; AliyunID++ {
		Amount := GetInvoiceAmount(AliyunID).Data.TotalInvoiceAmount
		AmountProjName := CsvFilter.CsvConvert()[AliyunID].AliyunNicknam
		AmountMetrix := AliyunAmountMetrix{AmountProjName, Amount}
		AmountMetrixRes[AliyunID] = AmountMetrix
	}
	return AmountMetrixRes
}

func main() {
	fmt.Println(GenerateAmountMetrix())
}
