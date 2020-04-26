package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
)

func main() {
	client, err := bssopenapi.NewClientWithAccessKey("cn-shanghai", "LTAI4GGpoj3JDwzDH7PcpiHH", "YzEY1PwEfs3VLbpfzMihP7V5wgi7VR")

	request := bssopenapi.CreateQueryEvaluateListRequest()
	request.Scheme = "https"

	request.BillCycle = "ALIYUN"

	response, err := client.QueryEvaluateList(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
