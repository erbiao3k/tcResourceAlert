package mongodb

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	mongodb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20190725"
	"tcResourceAlert/cloud"
	"tcResourceAlert/profile"
	"tcResourceAlert/utils"
	"time"
)

type Instance struct {
	// 实例ID
	Id string

	// 实例名称
	Name string

	// 实例所在区域
	Zone string

	// 实例类型，副本集：0，分片副本集：1
	ClusterType int64

	// 副本集
	ReplicaSet *[]ReplicaSet

	// 副本集IP
	ReplicaSetVip string

	// mongosId
	MongosId []string

	// 运行状态
	State string
}

type ReplicaSet struct {
	// 副本集ID
	ReplicaSetId string

	// 副本集名称
	// 副本集ID
	ReplicaSetName string

	// 副本集节点ID
	ReplicaSetNodeId []string
}

var client, _ = mongodb.NewClient(cloud.NewCfg(profile.MongodbCpfEndpoint, regions.Guangzhou))

func GetInstancelist() (*[]Instance, error) {

	request := mongodb.NewDescribeDBInstancesRequest()

	request.ClusterType = common.Int64Ptr(-1)
	request.Limit = common.Uint64Ptr(100)
	request.Offset = common.Uint64Ptr(0)
	request.OrderBy = common.StringPtr("CreateTime")
	request.OrderByType = common.StringPtr("DESC")

	response, err := client.DescribeDBInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {

		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var list []Instance

	for _, inst := range response.Response.InstanceDetails {
		//var rs []ReplicaSet
		//for _, shard := range inst.ReplicaSets {
		//	rs = append(rs, ReplicaSet{
		//		ReplicaSetId:   *shard.ReplicaSetId,
		//		ReplicaSetName: *shard.ReplicaSetName,
		//	})
		//}

		mongos, replicaset, err := getNodeProperty(*inst.InstanceId)
		if err != nil {
			return nil, err
		}

		list = append(list, Instance{
			Id:            *inst.InstanceId,
			Name:          *inst.InstanceName,
			Zone:          *inst.Zone,
			ClusterType:   int64(*inst.ClusterType),
			ReplicaSet:    &replicaset,
			ReplicaSetVip: *inst.Vip,
			MongosId:      mongos,
			State:         *inst.InstanceStatusDesc,
		})
	}
	return &list, err
}

func getNodeProperty(instanceid string) (mongosId []string, ReplicaSets []ReplicaSet, err error) {
	time.Sleep(time.Duration(0.4 * float64(time.Second)))
	request := mongodb.NewDescribeDBInstanceNodePropertyRequest()
	request.InstanceId = common.StringPtr(instanceid)
	response, err := client.DescribeDBInstanceNodeProperty(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	for _, i := range response.Response.Mongos {
		mongosId = append(mongosId, *i.NodeName)
	}

	for _, rss := range response.Response.ReplicateSets {
		var nodeids []string
		var rsId string
		for _, i := range rss.Nodes {
			nodeids = append(nodeids, *i.NodeName)
			rsId = *i.ReplicateSetId
		}
		ReplicaSets = append(ReplicaSets, ReplicaSet{
			ReplicaSetId:     rsId,
			ReplicaSetName:   rsId,
			ReplicaSetNodeId: nodeids,
		})
	}
	return
}

func InstanceInfo(instanceid string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		if instance.Id == instanceid {
			return &instance
		}
	}
	return nil
}

func InstanceInfoForRsid(rsId string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		for _, rs := range *instance.ReplicaSet {
			if rs.ReplicaSetId == rsId {
				return &instance
			}
		}
	}
	return nil
}

func InstanceInfoForNodeids(nodeId string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		for _, rs := range *instance.ReplicaSet {
			if utils.Inslice(rs.ReplicaSetNodeId, nodeId) {
				return &instance
			}
		}
	}
	return nil
}

func InstanceInfoForMongosids(Mongosid string, instances *[]Instance) *Instance {
	for _, instance := range *instances {
		if utils.Inslice(instance.MongosId, Mongosid) {
			return &instance
		}
	}
	return nil
}
