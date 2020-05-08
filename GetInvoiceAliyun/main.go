package main

import (
	"fmt"
	"local/AliyunInvoiceCore"
	"local/CsvFilter"
)

var CsvConvertRes = CsvFilter.CsvConvert()
var EncryptedMetrixLen int = len(CsvConvertRes)

func main() {
	TestMetrix := AliyunInvoiceCore.GenerateAmountMetrix()
	for i := 1; i <= EncryptedMetrixLen; i++ {
		fmt.Printf("%v, ", TestMetrix[i].AliyunProjName)
		fmt.Printf("%.2f, ", TestMetrix[i].AliyunInvoiceAmount)
		fmt.Printf("%v\n", TestMetrix[i].AliyunInvoiceId)
	}
	GenerateAddrMetrixRes := AliyunInvoiceCore.GenerateAddrMetrix()
	for i := 1; i <= EncryptedMetrixLen; i++ {
		fmt.Printf("%v,%v,%v\n", GenerateAddrMetrixRes[i].AliyunAddrProjName, GenerateAddrMetrixRes[i].AliyunAddressID, GenerateAddrMetrixRes[i].AliyunAddressee)
	}

}
