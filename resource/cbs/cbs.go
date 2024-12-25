package cbs

import (
	cbs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs/v20170312"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"tcResourceAlert/cloud"
	"tcResourceAlert/profile"
)

type Disk struct {
	// 名称
	Name string `json:"name"`

	// ID
	Id string

	// 绑定的CVM实例ID
	InstanceId string `json:"instance_id"`

	// 位置
	Zone string `json:"location"`

	// 云硬盘类型。取值范围：SYSTEM_DISK：系统盘DATA_DISK：数据盘。
	Usage string `json:"usage"`
}

func GetInstancelist() (*[]Disk, []string, error) {

	client, _ := cbs.NewClient(cloud.NewCfg(profile.CbsCpfEndpoint, regions.Guangzhou))

	request := cbs.NewDescribeDisksRequest()

	request.Limit = common.Uint64Ptr(100)
	request.OrderField = common.StringPtr("CREATE_TIME")
	request.Offset = common.Uint64Ptr(0)
	request.Order = common.StringPtr("DESC")

	response, err := client.DescribeDisks(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	var CBSs []Disk
	var unmountCBS []string
	for _, disk := range response.Response.DiskSet {
		CBSs = append(CBSs, Disk{
			Id:         *disk.DiskId,
			Name:       *disk.DiskName,
			InstanceId: *disk.InstanceId,
			Zone:       *disk.Placement.Zone,
			Usage:      *disk.DiskUsage,
		})
		if disk.InstanceId == nil {
			unmountCBS = append(unmountCBS, *disk.DiskId)
		}
	}

	return &CBSs, unmountCBS, nil
}

func InstanceInfo(instanceid string, instances *[]Disk) *Disk {
	for _, instance := range *instances {
		if instance.Id == instanceid {
			return &instance
		}
	}
	return nil
}
