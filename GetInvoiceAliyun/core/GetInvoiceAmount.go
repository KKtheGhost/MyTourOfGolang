package AliyunInvoiceCore

//本工具目的是,面对日益增多的阿里云账号,通过工具自动每月月初获取发票,从而大幅度减少无意义的重复劳动.

//需要导入的无非就是aliyunbssopenapi
//以及处理json的开源插件gojsonq
//fmt仅用于调试,之后可以去掉.
import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

//本函数的报错代码
var GetAmountErrorCode = map[int]string{
	1: "ERROR! GAE001: Missing Csv FIle - Cannot Read the DecryptCode file or the file is not existing.",
	2: "ERROR! GAE002: Error Building Client - Please check out network status and browser settings.",
	3: "ERROR! GAE003: Error Invoice Request - Please contact Aliyun support to figure out causes.",
}

//设定项目-账单结构体
type AliyunAmountMetrix struct {
	AliyunProjName      string
	AliyunInvoiceAmount float64
	AliyunInvoiceId     []string
}

//设定日期格式用于下面获取日期
const DATE_FORMAT = "20060102"

//本变量重复申明，最后打包main的时候需要删除。
//var CsvConvertRes = CsvFilter.CsvConvert()
//var EncryptedMetrixLen int = len(CsvConvertRes)

//本函数为重复函数，最后打包main的时候需要删除。core本身也应该作为一个包去使用
//func TokenDecrypt(EncryptedMetrix CsvFilter.AliyunTokenSet) (accessKey string, accessSecret string) {
//	DecryptKey, DecryptKeyErr := ioutil.ReadFile("/etc/GetAliyunInvoice/DecryptCode.key")
//	if DecryptKeyErr != nil {
//		fmt.Println(GetAmountErrorCode[1])
//		os.Exit(0)
//	}
//	//解密token，获取原始的tokenKey和tokenSecret
//	OriginAccessKey := AesCalculate.AesDecrypt(EncryptedMetrix.AliyunTokenKeyEncrypt, string(DecryptKey))
//	OriginAccessSecret := AesCalculate.AesDecrypt(EncryptedMetrix.AliyunTokenSecretEncrypt, string(DecryptKey))
//	return OriginAccessKey, OriginAccessSecret
//}

//本函数获取本月可申请账单额度
func GetInvoiceAmount(LastMonth string, AliyunID int) (response *bssopenapi.QueryEvaluateListResponse) {
	AmountProjName := CsvConvertRes[AliyunID].AliyunNicknam
	accessKey, accessSecret := TokenDecrypt(CsvConvertRes[AliyunID])
	AliyunInvoiceClient, AliyunClientErr := bssopenapi.NewClientWithAccessKey("cn-shanghai", accessKey, accessSecret)
	if AliyunClientErr != nil {
		fmt.Println(GetAmountErrorCode[2])
		os.Exit(0)
	}
	AliyunInvoiceInfo := bssopenapi.CreateQueryEvaluateListRequest()
	AliyunInvoiceInfo.Scheme = "https"
	AliyunInvoiceInfo.BillCycle = LastMonth
	AliyunInvoiceInfo.BizTypeList = &[]string{"ALIYUN"}

	AliyunInvoiceResponse, AliyunInvoiceErr := AliyunInvoiceClient.QueryEvaluateList(AliyunInvoiceInfo)
	if AliyunInvoiceErr != nil {
		fmt.Println(GetAmountErrorCode[3], "\n", AmountProjName, "has an incorrect token!")
		os.Exit(0)
	}
	return AliyunInvoiceResponse
	//fmt.Printf("%#v\n", AliyunInvoiceResponse.Data.TotalInvoiceAmount)
}

//通过循环打包，输出一个结构体合集，包含项目名+可开票金额
func GenerateAmountMetrix() map[int]AliyunAmountMetrix {
	//获取上个月月份
	Year, Month, _ := time.Now().Date()
	ThisMonth := time.Date(Year, Month, 1, 0, 0, 0, 0, time.Local)
	LastMonth := ThisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT)
	LastMonth = LastMonth[:6]
	//=======================
	AmountMetrixRes := make(map[int]AliyunAmountMetrix)
	//此处上界为大于等于，因为Aliyun起始是1
	for AliyunID := 1; AliyunID <= EncryptedMetrixLen; AliyunID++ {
		GetInvoiceAmountRes := GetInvoiceAmount(LastMonth, AliyunID).Data
		RawAmount := GetInvoiceAmountRes.TotalUnAppliedInvoiceAmount
		//获取完整的invoice ID表
		EvaluateLen := len(GetInvoiceAmountRes.EvaluateList.Evaluate)
		var EvaluateIds []string
		for EvaluateId := 0; EvaluateId < EvaluateLen; EvaluateId++ {
			//fmt.Println(GetInvoiceAmountRes.EvaluateList.Evaluate[EvaluateId].Id)
			EvaluateIds = append(EvaluateIds, strconv.FormatInt(GetInvoiceAmountRes.EvaluateList.Evaluate[EvaluateId].Id, 10))
		}
		TrueAmount := float64(RawAmount) / 100
		AmountProjName := CsvConvertRes[AliyunID].AliyunNicknam
		AmountMetrix := AliyunAmountMetrix{AmountProjName, TrueAmount, EvaluateIds}
		AmountMetrixRes[AliyunID] = AmountMetrix
	}
	return AmountMetrixRes
}

//main仅供测试使用。
//func main() {
//	TestMetrix := GenerateAmountMetrix()
//	for i := 1; i <= EncryptedMetrixLen; i++ {
//		fmt.Printf("%v, ", TestMetrix[i].AliyunProjName)
//		fmt.Printf("%.2f, ", TestMetrix[i].AliyunInvoiceAmount)
//		fmt.Printf("%v\n", TestMetrix[i].AliyunInvoiceId)
//	}
//}
