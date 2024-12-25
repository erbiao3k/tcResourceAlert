package monitor

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"tcResourceAlert/profile"
	"tcResourceAlert/utils"
)

type StatisticData struct {
	// 集群名称
	ClusterName string `json:"集群名称,omitempty"`

	// 集群ID
	ClusterId string `json:"集群ID,omitempty"`

	// 节点ID
	NodeId string `json:"节点ID,omitempty"`

	// 节点类型
	NodeRole string `json:"节点类型,omitempty"`

	// 节点IP
	NodeIP string `json:"节点IP,omitempty"`

	//命名空间
	Namespace string `json:"命名空间,omitempty"`

	// 工作负载名称
	WorkloadName string `json:"工作负载名称,omitempty"`

	// 工作负载类型
	WorkloadKind string `json:"工作负载类型,omitempty"`

	// POD名称
	PodName string `json:"POD名称,omitempty"`

	// Container名称
	ContainerName string `json:"Container名称,omitempty"`

	// ContainerID
	ContainerId string `json:"containerID,omitempty"`

	// 监控指标
	MetricName string `json:"监控指标,omitempty"`

	// 数据值
	MetricValue float64 `json:"数据值,omitempty"`

	// 告警提示
	Prompt string `json:"告警提示,omitempty"`
}

func DescStatisticData(namespace, startTime, endTime string, metric []string, conditions []*monitor.MidQueryCondition) (*[]StatisticData, error) {
	request := monitor.NewDescribeStatisticDataRequest()
	request.Module = common.StringPtr("monitor")
	request.Namespace = common.StringPtr(namespace)
	request.MetricNames = common.StringPtrs(metric)
	request.Conditions = conditions
	request.Period = common.Uint64Ptr(60)
	request.StartTime = common.StringPtr(startTime)
	request.EndTime = common.StringPtr(endTime)
	response, err := client.DescribeStatisticData(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var statisticData []StatisticData
	for _, i := range response.Response.Data {
		for _, p := range i.Points {
			sData := StatisticData{
				MetricName: *i.MetricName,
			}
			for _, d := range p.Dimensions {
				switch *d.Name {
				case profile.Dimensionnode:
					sData.NodeIP = *d.Value
				case profile.Dimensionnode_role:
					sData.NodeRole = *d.Value
				case profile.Dimensiontke_cluster_instance_id:
					sData.ClusterId = *d.Value
				case profile.Dimensionun_instance_id:
					sData.NodeId = *d.Value
				case profile.Dimensionworkload_name:
					sData.WorkloadName = *d.Value
				case profile.Dimensionpod_name:
					sData.PodName = *d.Value
				case profile.Dimensionnamespace:
					sData.Namespace = *d.Value
				case profile.Dimensionworkload_kind:
					sData.WorkloadKind = *d.Value
				}
			}

			var values []*float64
			for _, v := range p.Values {
				values = append(values, v.Value)
			}
			sData.MetricValue = utils.Float(utils.Average(values))
			statisticData = append(statisticData, sData)
		}
	}
	return &statisticData, err
}
