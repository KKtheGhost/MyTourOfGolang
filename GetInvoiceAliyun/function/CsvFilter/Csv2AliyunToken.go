//package CsvFilter
package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	AliyunTokenSecretEncrypt string
}

func CsvReader() [][]string {
	//var AliyunTokenMetrix map[int]AliyunTokenSet
	CsvFileName := "AliyunToken.csv"
	CsvFilePath := "/etc/GetAliyunInvoice/"
	CsvFile := CsvFilePath + CsvFileName                                //此处最好还是绝对路径写死吧，比如需要把配置放到/etc或者/opt下面，这样比较安全。
	CsvFileRawContent, CsvFileRawContentErr := ioutil.ReadFile(CsvFile) //需要判断CSV的存在性，及时中断
	if CsvFileRawContentErr != nil {
		fmt.Println(CsvErrorCode[1])
		os.Exit(0)
	}
	CsvFileReaderContent := csv.NewReader(strings.NewReader(string(CsvFileRawContent)))
	CsvFileContent, CsvFileContentErr := CsvFileReaderContent.ReadAll()
	if CsvFileContentErr != nil {
		fmt.Println(CsvErrorCode[2])
		os.Exit(0)
	}
	//此处需要验证输出的合法性：1. 字段合规  2.加密token格式合规 3.解密格式合规 4.存在两行及以上的输出。
	//fmt.Println(CsvFileContent) //测试能否获取输出用语句
	return CsvFileContent
}

//把CsvReader中获取的[][]string结构，通过循环打到结构体AliyunTokenMetrix里面去
//func CsvConvert() {
func main() {
	AliyunTokenMetrix := make(map[int]AliyunTokenSet)
	CsvFileContent := CsvReader()
	AliyunTokenSetNum := len(CsvFileContent)
	if AliyunTokenSetNum < 2 {
		fmt.Println(CsvErrorCode[3])
		os.Exit(0)
	}
	//fmt.Println(CsvFileContent)
	for i := 1; i < AliyunTokenSetNum; i++ {
		AliyunToken := AliyunTokenSet{CsvFileContent[i][1], CsvFileContent[i][2], CsvFileContent[i][3]}
		AliyunTokenMetrix[i] = AliyunToken
	}
	fmt.Println(AliyunTokenMetrix)
}
