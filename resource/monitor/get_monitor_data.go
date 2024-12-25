package monitor

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"log"
	"strings"
	"tcResourceAlert/cloud"
	"tcResourceAlert/profile"
	"tcResourceAlert/resource/redis"
	"tcResourceAlert/utils"
)

var client, _ = monitor.NewClient(cloud.NewCfg(profile.MonitorCpfEndpoint, regions.Guangzhou))

type MetricData struct {
	// 实例ID
	InstanceId string

	// 实例名称
	InstanceName string

	// 实例类型
	InstanceType string

	// 监控指标
	MetricName string

	// 数据值
	MetricValue float64
}

func GetMonitorData(namespace, metric, startTime, endTime string, Instance [][]*monitor.Instance) (*[]MetricData, error) {
	request := monitor.NewGetMonitorDataRequest()
	request.Namespace = common.StringPtr(namespace)
	request.MetricName = common.StringPtr(metric)
	request.Period = common.Uint64Ptr(60)
	request.StartTime = common.StringPtr(startTime)
	request.EndTime = common.StringPtr(endTime)

	var list []MetricData

	for _, instances := range Instance {
		request.Instances = instances
		response, err := client.GetMonitorData(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		for _, i := range response.Response.DataPoints {
			var dimension []string
			for _, d := range i.Dimensions {
				dimension = append(dimension, *d.Value)
			}
			id := strings.Join(dimension, " ")
			list = append(list, MetricData{
				InstanceId:   id,
				InstanceType: namespace,
				MetricName:   metric,
				MetricValue:  utils.Average(i.Values),
			})
		}
		if len(*response.Response.Msg) > 0 {
			log.Printf("【调用限制】Response：%v，", response.ToJsonString())
		}
	}
	return &list, nil
}

func Instance(dimensionKey string, instanceIds []string) [][]*monitor.Instance {
	var instancesSlice [][]*monitor.Instance
	for _, ids := range utils.SplitIntoChunks(instanceIds, 10) {
		var instances []*monitor.Instance
		for _, instanceid := range ids {
			dimension := &monitor.Dimension{
				Name:  common.StringPtr(dimensionKey),
				Value: common.StringPtr(instanceid),
			}
			instance := &monitor.Instance{
				Dimensions: []*monitor.Dimension{dimension},
			}
			instances = append(instances, instance)
		}
		instancesSlice = append(instancesSlice, instances)
	}
	return instancesSlice
}

func RedisProxyInstance(Dimensioninstanceid string, Dimensionpnodeid string, instancelist *[]redis.Instance) [][]*monitor.Instance {
	var instancesSlice [][]*monitor.Instance
	var instanceIds []string
	for _, i := range *instancelist {
		instanceIds = append(instanceIds, i.Id)
	}

	for _, instanceids := range utils.SplitIntoChunks(instanceIds, 1) {
		var instances []*monitor.Instance
		for _, id := range instanceids {
			info := redis.InstanceInfo(id, instancelist)
			for _, proxyid := range info.ProxyId {
				dimension := &monitor.Dimension{
					Name:  common.StringPtr(Dimensioninstanceid),
					Value: common.StringPtr(id),
				}
				dimension2 := &monitor.Dimension{
					Name:  common.StringPtr(Dimensionpnodeid),
					Value: common.StringPtr(proxyid),
				}
				instance := &monitor.Instance{
					Dimensions: []*monitor.Dimension{dimension, dimension2},
				}
				instances = append(instances, instance)
			}
		}
		instancesSlice = append(instancesSlice, instances)
	}
	return instancesSlice
}

func RedisNodeInstance(Dimensioninstanceid string, Dimensionrnodeid string, instancelist *[]redis.Instance) [][]*monitor.Instance {
	var instancesSlice [][]*monitor.Instance
	var instanceIds []string
	for _, i := range *instancelist {
		instanceIds = append(instanceIds, i.Id)
	}

	for _, instanceids := range utils.SplitIntoChunks(instanceIds, 1) {
		var instances []*monitor.Instance
		for _, id := range instanceids {
			info := redis.InstanceInfo(id, instancelist)
			for _, nodeid := range info.NodeId {
				dimension := &monitor.Dimension{
					Name:  common.StringPtr(Dimensioninstanceid),
					Value: common.StringPtr(id),
				}
				dimension2 := &monitor.Dimension{
					Name:  common.StringPtr(Dimensionrnodeid),
					Value: common.StringPtr(nodeid),
				}
				instance := &monitor.Instance{
					Dimensions: []*monitor.Dimension{dimension, dimension2},
				}
				instances = append(instances, instance)
			}
		}
		instancesSlice = append(instancesSlice, instances)
	}
	return instancesSlice
}
