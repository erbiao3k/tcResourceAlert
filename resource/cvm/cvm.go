package cvm

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sdkerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"tcResourceAlert/cloud"
	"tcResourceAlert/profile"
)

type Instance struct {
	// 实例名称
	Name string `json:"name"`

	// 实例ID
	Id string `json:"instance_id"`

	// 实例所在的位置
	Zone string `json:"zone"`

	// 实例运行状态
	State string `json:"status"`
}

func GetInstancelist() (*[]Instance, error) {
	client, _ := cvm.NewClient(cloud.NewCfg(profile.CvmCpfEndpoint, regions.Guangzhou))
	request := cvm.NewDescribeInstancesRequest()

	request.Filters = []*cvm.Filter{
		{
			Name:   common.StringPtr("instance-state"),
			Values: common.StringPtrs([]string{"RUNNING"}),
		},
	}
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(100)

	response, err := client.DescribeInstances(request)

	if _, ok := err.(*sdkerr.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var CVMs []Instance
	for _, instance := range response.Response.InstanceSet {
		CVMs = append(CVMs, Instance{
			Name:  *instance.InstanceName,
			Id:    *instance.InstanceId,
			Zone:  *instance.Placement.Zone,
			State: *instance.InstanceState,
		})
	}
	return &CVMs, nil
}

func InstanceInfo(instanceid string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		if instance.Id == instanceid {
			return &instance
		}
	}
	return nil
}
