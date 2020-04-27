//package CsvFilter
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/csv"
)

//本package的ERROR CODE表。
var CsvErrorCode = map[int]string{
	1: "ERROR! CSV001: Missing Csv FIle - Cannot Read the CSV file or the file is not existing.",
	2: "ERROR! CSV002: The csv file is invalid, please check file content.",
	3: "ERROR! CSV003: The csv file has no valid data, please check and fill the file.",
}

//本函数目的就是将CSV文件转化为下面的结构体, 并且输出到core功能中, 通过AES解密后, 给Client使用.下面的数据都需要传递.
type AliyunTokenSet struct {
	AliyunNicknam            string
	AliyunDecryptKey         string
	AliyunTokenKeyEncrypt    string
	AliyunTokenSecretEncrypt string
}

func CsvReader() bool{
	//var AliyunTokenMetrix map[int]AliyunTokenSet
	CsvFileName := "AliyunToken.csv"
	CsvFilePath := "/etc/GetAliyunInvoice/"
	CsvFile := CsvFilePath + CsvFileName	//此处最好还是绝对路径写死吧，比如需要把配置放到/etc或者/opt下面，这样比较安全。
	CsvFileRawContent, CsvFileRawContentErr := ioutil.ReadFile(CsvFile)	//需要判断CSV的存在性，及时中断
	if CsvFileRawContentErr != nil {
		fmt.Println(CsvErrorCode[1])
		return false
	}
	CsvFileReaderContent := csv.NewReader(strings.NewReader(string(CsvFileRawContent)))
	CsvFileContent, CsvFileContentErr := CsvFileReaderContent.ReadAll()	
	if CsvFileContentErr != nil {
		fmt.Println(CsvErrorCode[2])
		return false
	}
	AliyunTokenSetNum := len(CsvFileContent)
	if AliyunTokenSetNum < 2 {
		fmt.Println(CsvErrorCode[3])
		return false
	}
	
	//此处需要验证输出的合法性：1. 字段合规  2.加密token格式合规 3.解密格式合规 4.存在两行及以上的输出。
	fmt.Println(CsvFileContent)	//测试能否获取输出用语句
	return true
}

//func CsvConvert() {
func main() {
	//var AliyunTokenMetrix map[int]AliyunTokenSet
	CsvReader()
}