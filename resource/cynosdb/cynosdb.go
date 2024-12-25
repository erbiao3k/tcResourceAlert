package cynosdb

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cynosdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cynosdb/v20190107"
	"tcResourceAlert/cloud"
	"tcResourceAlert/profile"
)

type Instance struct {
	// 实例ID
	Id string
	// 实例名称
	Name string
	// 实例状态
	State string
	// 所在区域
	Zone string
	// 实例类型：只读实例：ro，读写实例：rw
	Type string
	// 所属集群名称
	ClusterName string
	// 所属集群ID
	ClusterId string
	// 集群的proxy ID
	ClusterProxyId string
}

var client, _ = cynosdb.NewClient(cloud.NewCfg(profile.CynosdbCpfEndpoint, regions.Guangzhou))

func getInstances() (*[]Instance, error) {

	request := cynosdb.NewDescribeInstancesRequest()

	request.Limit = common.Int64Ptr(100)
	request.Offset = common.Int64Ptr(0)
	request.OrderBy = common.StringPtr("CREATETIME")
	request.OrderByType = common.StringPtr("DESC")
	request.Status = common.StringPtr("running")

	response, err := client.DescribeInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	var list []Instance
	for _, inst := range response.Response.InstanceSet {
		list = append(list, Instance{
			Id:          *inst.InstanceId,
			Name:        *inst.InstanceName,
			State:       *inst.StatusDesc,
			Zone:        *inst.Zone,
			Type:        *inst.InstanceType,
			ClusterName: *inst.ClusterName,
			ClusterId:   *inst.ClusterId,
		})
	}
	return &list, err
}

func getProxy(instance Instance) (*Instance, error) {
	request := cynosdb.NewDescribeProxiesRequest()

	request.ClusterId = common.StringPtr(instance.ClusterId)

	request.Limit = common.Int64Ptr(100)
	request.Offset = common.Int64Ptr(0)
	request.OrderBy = common.StringPtr("CREATETIME")
	request.OrderByType = common.StringPtr("DESC")

	response, err := client.DescribeProxies(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	for _, proxy := range response.Response.ProxyNodeInfos {
		instance.ClusterProxyId = *proxy.ProxyNodeId
	}
	return &instance, err
}

func GetInstancelist() (*[]Instance, error) {
	var list []Instance
	cynosdbInstances, _ := getInstances()
	for _, instance := range *cynosdbInstances {
		proxy, err := getProxy(instance)
		if err != nil {
			return nil, err
		}
		list = append(list, *proxy)
	}
	return &list, nil
}

func InstanceInfo(instanceid string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		if instance.Id == instanceid {
			return &instance
		}
	}
	return nil
}
