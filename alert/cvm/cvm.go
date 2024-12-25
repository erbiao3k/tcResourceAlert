package cvm

import (
	"errors"
	"fmt"
	"tcResourceAlert/msg_sender"
	"tcResourceAlert/profile"
	"tcResourceAlert/resource/cvm"
	"tcResourceAlert/resource/monitor"
	"tcResourceAlert/utils"
)

func Alert(startTime, endTime string) error {
	cvmlist, err := cvm.GetInstancelist()
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstancelistErrmsg, profile.NamespaceCvm, err))
	}

	var cvmInstanceIds []string
	for _, i := range *cvmlist {
		cvmInstanceIds = append(cvmInstanceIds, i.Id)
	}

	data, err := monitor.GetMonitorData(profile.NamespaceCvm, profile.MetricCvmCpuUsage, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, cvmInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCvm, profile.MetricCvmCpuUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := cvm.InstanceInfo(instance.InstanceId, cvmlist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "服务器CPU利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceCvm, profile.MetricCvmBaseCpuUsage, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, cvmInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCvm, profile.MetricCvmBaseCpuUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := cvm.InstanceInfo(instance.InstanceId, cvmlist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "服务器宿主机上报的CPU利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceCvm, profile.MetricCvmCpuLoadavg, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, cvmInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCvm, profile.MetricCvmCpuLoadavg, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(10) {
			instanceinfo := cvm.InstanceInfo(instance.InstanceId, cvmlist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "服务器CPU一分钟平均负载(计数)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceCvm, profile.MetricCvmMemUsage, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, cvmInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCvm, profile.MetricCvmMemUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := cvm.InstanceInfo(instance.InstanceId, cvmlist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "服务器内存利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}
	data, err = monitor.GetMonitorData(profile.NamespaceCvm, profile.MetricCvmDiskUsage, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, cvmInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCvm, profile.MetricCvmDiskUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := cvm.InstanceInfo(instance.InstanceId, cvmlist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "服务器所有磁盘已使用容量占总容量的百分比(%)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}
	return nil
}
