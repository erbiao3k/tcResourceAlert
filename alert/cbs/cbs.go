package cbs

import (
	"errors"
	"fmt"
	"tcResourceAlert/msg_sender"
	"tcResourceAlert/profile"
	"tcResourceAlert/resource/cbs"
	"tcResourceAlert/resource/monitor"
	"tcResourceAlert/utils"
)

func Alert(startTime, endTime string) error {
	disklist, unmountDisk, err := cbs.GetInstancelist()
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstancelistErrmsg, profile.NamespaceCbs, err))
	}

	if len(unmountDisk) != 0 {
		for _, instanceid := range unmountDisk {
			instanceinfo := cbs.InstanceInfo(instanceid, disklist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "磁盘未挂载使用", "磁盘未挂载使用")
			msg_sender.MsgSender(msg)
		}
	}

	var diskIds []string
	for _, id := range *disklist {
		diskIds = append(diskIds, id.Id)
	}

	data, err := monitor.GetMonitorData(profile.NamespaceCbs, profile.MetricCvmDiskIoAwait, startTime, endTime, monitor.Instance(profile.DimensionDiskId, diskIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCbs, profile.MetricCvmDiskIoAwait, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(20) {
			instanceinfo := cbs.InstanceInfo(instance.InstanceId, disklist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "硬盘IO等待时间(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceCbs, profile.MetricCvmDiskUtil, startTime, endTime, monitor.Instance(profile.DimensionDiskId, diskIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCbs, profile.MetricCvmDiskUtil, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(80) {
			instanceinfo := cbs.InstanceInfo(instance.InstanceId, disklist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "硬盘IO繁忙比率(%)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	return nil
}
