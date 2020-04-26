package main

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func main() {
	AliyunInvoiceClient, err := bssopenapi.NewClientWithAccessKey("cn-shanghai", "accessKey", "accessSecret")

	AliyunInvoiceRequest := bssopenapi.CreateQueryCustomerAddressListRequest()
	AliyunInvoiceRequest.Scheme = "https"

	AliyunInvoiceResponse, AliyunInvoiceErr := AliyunInvoiceClient.QueryCustomerAddressList(AliyunInvoiceRequest)
	if AliyunInvoiceErr != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("%#v\n", AliyunInvoiceResponse.Data.CustomerInvoiceAddressList)
}
