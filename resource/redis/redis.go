package redis

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	redis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
	"tcResourceAlert/cloud"
	"tcResourceAlert/profile"
	"tcResourceAlert/utils"
)

type Instance struct {
	Id      string
	Name    string
	State   string
	Zone    string
	ProxyId []string
	NodeId  []string
}

var client, _ = redis.NewClient(cloud.NewCfg(profile.RedisCpfEndpoint, regions.Guangzhou))

func GetInstancelist() (*[]Instance, error) {
	request := redis.NewDescribeInstancesRequest()

	request.Limit = common.Uint64Ptr(100)
	request.Offset = common.Uint64Ptr(0)
	request.OrderBy = common.StringPtr("createtime")
	request.OrderType = common.Int64Ptr(1)
	request.Status = common.Int64Ptrs([]int64{2})

	response, err := client.DescribeInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var list []Instance
	for _, inst := range response.Response.InstanceSet {
		instance := Instance{
			Id:    *inst.InstanceId,
			Name:  *inst.InstanceName,
			State: *inst.InstanceTitle,
			Zone:  *inst.Region,
		}

		instances, err := getNodelist(instance)
		if err != nil {
			return nil, err
		}
		list = append(list, *instances)
	}
	return &list, err
}

func getNodelist(instance Instance) (*Instance, error) {
	request := redis.NewDescribeInstanceNodeInfoRequest()

	request.InstanceId = common.StringPtr(instance.Id)
	request.Limit = common.Int64Ptr(100)
	request.Offset = common.Int64Ptr(0)

	response, err := client.DescribeInstanceNodeInfo(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var proxys []string
	for _, p := range response.Response.Proxy {
		proxys = append(proxys, *p.NodeId)
	}
	var nodes []string
	for _, r := range response.Response.Redis {
		nodes = append(nodes, *r.NodeId)
	}
	instance.ProxyId = proxys
	instance.NodeId = nodes
	return &instance, nil
}

func InstanceInfo(instanceid string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		if instance.Id == instanceid {
			return &instance
		}
	}
	return nil
}

func InstanceInfoForProxyid(proxyid string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		if utils.Inslice(instance.ProxyId, proxyid) {
			return &instance
		}
	}
	return nil
}

func InstanceInfoForNodeid(nodeid string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		if utils.Inslice(instance.NodeId, nodeid) {
			return &instance
		}
	}
	return nil
}
