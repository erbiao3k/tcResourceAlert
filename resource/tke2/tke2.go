package tke2

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"
	"tcResourceAlert/cloud"
	"tcResourceAlert/profile"
)

var client, _ = tke.NewClient(cloud.NewCfg(profile.TkeCpfEndpoint, regions.Guangzhou))

type TkeCluster struct {
	// kubernetes集群ID
	Id string
	// kubernetes集群名称
	Name string

	// kubernetes集群名称
	NodePool *[]Node
}

type Node struct {
	// 实例ID
	Id string

	// 节点类型：MASTER, WORKER, ETCD, MASTER_ETCD
	Type string

	// 所在节点池
	NodePool string
}

func GetInstancelist() (*[]TkeCluster, error) {

	request := tke.NewDescribeClustersRequest()

	response, err := client.DescribeClusters(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var tc []TkeCluster
	for _, i := range response.Response.Clusters {

		property, err := getNodeProperty(*i.ClusterId)
		if err != nil {
			return nil, err
		}

		tkeCluster := TkeCluster{
			Id:       *i.ClusterId,
			Name:     *i.ClusterName,
			NodePool: property,
		}

		tc = append(tc, tkeCluster)
	}
	return &tc, nil
}

func getNodeProperty(clusterid string) (*[]Node, error) {
	request := tke.NewDescribeClusterInstancesRequest()

	request.ClusterId = common.StringPtr(clusterid)
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(100)
	request.InstanceRole = common.StringPtr("ALL")

	response, err := client.DescribeClusterInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	var nodelist []Node
	for _, instance := range response.Response.InstanceSet {
		nodelist = append(nodelist, Node{
			Id:       *instance.InstanceId,
			Type:     *instance.InstanceRole,
			NodePool: *instance.NodePoolId,
		})
	}
	return &nodelist, err
}
