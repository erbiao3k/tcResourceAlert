package tke2

import (
	"errors"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tke2moniotor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	"tcResourceAlert/msg_sender"
	"tcResourceAlert/profile"
	"tcResourceAlert/resource/monitor"
	"tcResourceAlert/resource/tke2"
	"tcResourceAlert/utils"
)

func Alert(startTime, endTime string) error {
	clusterlist, err := tke2.GetInstancelist()
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstancelistErrmsg, profile.NamespaceTke2, err))
	}

	var clusterids []string
	for _, clusterid := range *clusterlist {
		clusterids = append(clusterids, clusterid.Id)
	}

	podMetrics := getMetric(profile.MetricTke2Pod)
	containerMetrics := getMetric(profile.MetricTke2Container)
	for _, cluster := range *clusterlist {
		var nodeids []string
		for _, nodeid := range *cluster.NodePool {
			nodeids = append(nodeids, nodeid.Id)
		}
		workloads, err := tke2PodMonitor(nodeids, podMetrics, cluster.Id, startTime, endTime, profile.MetricTke2Pod)
		if err != nil {
			return err
		}

		err = tke2ContainerMonitor(nodeids, containerMetrics, workloads, cluster.Id, startTime, endTime, profile.MetricTke2Container)
		if err != nil {
			return err
		}
	}

	nodeMetrics := getMetric(profile.MetricTke2Node)
	err = tke2Monitor(clusterids, nodeMetrics, startTime, endTime, profile.MetricTke2Node)
	if err != nil {
		return err
	}

	workloadMetrics := getMetric(profile.MetricTke2Workload)
	err = tke2Monitor(clusterids, workloadMetrics, startTime, endTime, profile.MetricTke2Workload)
	if err != nil {
		return err
	}

	clusterMetrics := getMetric(profile.MetricTke2Cluster)
	err = tke2Monitor(clusterids, clusterMetrics, startTime, endTime, profile.MetricTke2Cluster)
	if err != nil {
		return err
	}

	pvcMetrics := getMetric(profile.MetricTke2Pvc)
	err = tke2Monitor(clusterids, pvcMetrics, startTime, endTime, profile.MetricTke2Pvc)
	if err != nil {
		return err
	}

	return nil
}

func trigger(data *[]monitor.StatisticData, metriclist []profile.MetricTke2) {
	for _, d := range *data {
		for _, m := range metriclist {
			var msgBool bool
			if d.MetricName == m.Metric {
				switch m.Operator {
				case profile.OperatorGreaterthan:
					if d.MetricValue > m.TriggerValue {
						msgBool = true
					}
				case profile.OperatorLessthan:
					if d.MetricValue < m.TriggerValue {
						msgBool = true
					}
				case profile.OperatorEqual:
					if d.MetricValue == m.TriggerValue {
						msgBool = true
					}
				}
			}
			if msgBool {
				d.Prompt = m.Prompt
				//log.Println(utils.ReplaceMsg(utils.Json(d)))
				msg_sender.MsgSender(utils.ReplaceMsg(utils.Json(d)))
			}
		}
	}
}

func tke2Monitor(clusterids, metrics []string, startTime, endTime string, metriclist []profile.MetricTke2) error {
	Condition := []*tke2moniotor.MidQueryCondition{
		{
			Key:      common.StringPtr(profile.Dimensiontke_cluster_instance_id),
			Operator: common.StringPtr("in"),
			Value:    common.StringPtrs(clusterids),
		},
	}

	data, err := monitor.DescStatisticData(profile.NamespaceTke2, startTime, endTime, metrics, Condition)
	if err != nil {
		return err
	}
	trigger(data, metriclist)

	return nil
}

func tke2PodMonitor(nodeids, metrics []string, clusterid, startTime, endTime string, metriclist []profile.MetricTke2) ([]string, error) {
	Condition := []*tke2moniotor.MidQueryCondition{
		{
			Key:      common.StringPtr(profile.Dimensiontke_cluster_instance_id),
			Operator: common.StringPtr("="),
			Value:    common.StringPtrs([]string{clusterid}),
		},
		{
			Key:      common.StringPtr(profile.Dimensionun_instance_id),
			Operator: common.StringPtr("in"),
			Value:    common.StringPtrs(nodeids),
		},
	}

	data, err := monitor.DescStatisticData(profile.NamespaceTke2, startTime, endTime, metrics, Condition)
	if err != nil {
		return nil, err
	}
	trigger(data, metriclist)

	var workloads []string
	for _, wl := range *data {
		if !utils.Inslice(workloads, wl.WorkloadName) {
			workloads = append(workloads, wl.WorkloadName)
		}
	}
	return workloads, nil
}

func getMetric(metriclist []profile.MetricTke2) []string {
	var metrics []string
	for _, m := range metriclist {
		metrics = append(metrics, m.Metric)
	}
	return metrics
}

func tke2ContainerMonitor(nodeids, metrics, workload []string, clusterid, startTime, endTime string, metriclist []profile.MetricTke2) error {
	Condition := []*tke2moniotor.MidQueryCondition{
		{
			Key:      common.StringPtr(profile.Dimensiontke_cluster_instance_id),
			Operator: common.StringPtr("="),
			Value:    common.StringPtrs([]string{clusterid}),
		},
		{
			Key:      common.StringPtr(profile.Dimensionun_instance_id),
			Operator: common.StringPtr("in"),
			Value:    common.StringPtrs(nodeids),
		},
		{
			Key:      common.StringPtr(profile.Dimensionworkload_name),
			Operator: common.StringPtr("in"),
			Value:    common.StringPtrs(workload),
		},
	}

	data, err := monitor.DescStatisticData(profile.NamespaceTke2, startTime, endTime, metrics, Condition)
	if err != nil {
		return err
	}
	trigger(data, metriclist)

	return nil
}
