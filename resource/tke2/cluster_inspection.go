package tke2

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
	"log"
)

func Inspection() error {
	request := tke.NewListClusterInspectionResultsRequest()

	response, err := client.ListClusterInspectionResults(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
	}
	if err != nil {
		return err
	}
	for _, cluster := range response.Response.InspectionResults {
		for _, diagnostic := range cluster.Diagnostics {
			for _, result := range diagnostic.Results {
				log.Println(*result.Desc)
			}
		}
		break
	}
	return nil
}
