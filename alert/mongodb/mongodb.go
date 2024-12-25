package mongodb

import (
	"errors"
	"fmt"
	"tcResourceAlert/msg_sender"
	"tcResourceAlert/profile"
	"tcResourceAlert/resource/mongodb"
	"tcResourceAlert/resource/monitor"
	"tcResourceAlert/utils"
)

func Alert(startTime, endTime string) error {
	mongolist, err := mongodb.GetInstancelist()
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstancelistErrmsg, profile.NamespaceMongo, err))
	}

	var mongoInstanceIds []string
	for _, i := range *mongolist {
		mongoInstanceIds = append(mongoInstanceIds, i.Id)
	}

	var monggoRsids []string
	for _, i := range *mongolist {
		for _, rs := range *i.ReplicaSet {
			monggoRsids = append(monggoRsids, rs.ReplicaSetId)
		}
	}

	var mongoNodeIds []string
	for _, i := range *mongolist {
		for _, rs := range *i.ReplicaSet {
			mongoNodeIds = append(mongoNodeIds, rs.ReplicaSetNodeId...)
		}
	}

	var mongosIds []string
	for _, i := range *mongolist {
		mongosIds = append(mongosIds, i.MongosId...)
	}

	data, err := monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbDelay100, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbDelay100, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(10000) {
			instanceinfo := mongodb.InstanceInfo(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB集群单位时间内成功请求延迟在100ms以上次数/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbAvgAllRequestDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbAvgAllRequestDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(10000) {
			instanceinfo := mongodb.InstanceInfo(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB集群所有请求平均延迟(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbConnper, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbConnper, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := mongodb.InstanceInfo(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB集群连接数利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbClusterDiskUsage, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoInstanceIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbClusterDiskUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := mongodb.InstanceInfo(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB集群存储利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbSlaveDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, monggoRsids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbSlaveDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(10) {
			instanceinfo := mongodb.InstanceInfoForRsid(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB主从单位时间内平均延迟(s)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbCpuUsage, startTime, endTime, monitor.Instance(profile.Dimensiontarget, monggoRsids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbCpuUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(90) {
			instanceinfo := mongodb.InstanceInfoForRsid(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点CPU利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbReplicaDiskUsage, startTime, endTime, monitor.Instance(profile.Dimensiontarget, monggoRsids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbReplicaDiskUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(90) {
			instanceinfo := mongodb.InstanceInfoForRsid(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB副本集存储利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbMemUsage, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbMemUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(90) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点内存利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeSlavedelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeSlavedelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(10) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB从节点延迟(s)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbDiskUsage, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbDiskUsage, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(90) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点磁盘利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgAllRequestDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgAllRequestDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(100) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点所有请求延迟平均值(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgUpdateDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgUpdateDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(1000) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点更新命令延迟平均值(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgInsertDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgInsertDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(1000) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点插入命令延迟平均值(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgReadDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgReadDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(1000) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点读命令延迟平均值(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgAggregateDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgAggregateDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(1000) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点aggregate命令延迟平均值(ms)"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgCountDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgCountDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(1000) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点counts命令延迟平均值(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgGetmoreDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgGetmoreDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(1000) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点getmore命令延迟平均值(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgDeleteDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgDeleteDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(1000) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点delete请求延迟平均值(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeAvgCommandDelay, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeAvgCommandDelay, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(1000) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点command请求延迟平均值(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceMongo, profile.MetricMongodbNodeDelay100, startTime, endTime, monitor.Instance(profile.Dimensiontarget, mongoNodeIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceMongo, profile.MetricMongodbNodeDelay100, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(100) {
			instanceinfo := mongodb.InstanceInfoForNodeids(instance.InstanceId, mongolist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "MongoDB节点延迟超过100ms请求数/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	return nil
}
