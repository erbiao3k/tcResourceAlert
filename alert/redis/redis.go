package redis

import (
	"errors"
	"fmt"
	"strings"
	"tcResourceAlert/msg_sender"
	"tcResourceAlert/profile"
	"tcResourceAlert/resource/monitor"
	"tcResourceAlert/resource/redis"
	"tcResourceAlert/utils"
)

func Alert(startTime, endTime string) error {
	redislist, err := redis.GetInstancelist()
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstancelistErrmsg, profile.NamespaceRedismem, err))
	}

	var instanceids []string
	for _, i := range *redislist {
		instanceids = append(instanceids, i.Id)
	}

	data, err := monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemCpuUtil, startTime, endTime, monitor.Instance(profile.Dimensioninstanceid, instanceids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemCpuUtil, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := redis.InstanceInfo(instance.InstanceId, redislist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "Redis实例侧，平均CPU利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemUtil, startTime, endTime, monitor.Instance(profile.Dimensioninstanceid, instanceids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemUtil, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := redis.InstanceInfo(instance.InstanceId, redislist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "Redis实例侧，平均内存利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemConnectionsUtil, startTime, endTime, monitor.Instance(profile.Dimensioninstanceid, instanceids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemConnectionsUtil, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := redis.InstanceInfo(instance.InstanceId, redislist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "Redis实例侧，平均连接数利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemLatencyAvg, startTime, endTime, monitor.Instance(profile.Dimensioninstanceid, instanceids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemLatencyAvg, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(100) {
			instanceinfo := redis.InstanceInfo(instance.InstanceId, redislist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "Redis实例侧，Proxy到RedisServer的执行平均延迟(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemLatencyRead, startTime, endTime, monitor.Instance(profile.Dimensioninstanceid, instanceids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemLatencyRead, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(100) {
			instanceinfo := redis.InstanceInfo(instance.InstanceId, redislist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "Redis实例侧，Proxy到RedisServer的读命令平均执行延迟(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemLatencyWrite, startTime, endTime, monitor.Instance(profile.Dimensioninstanceid, instanceids))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemLatencyWrite, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(100) {
			instanceinfo := redis.InstanceInfo(instance.InstanceId, redislist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "Redis实例侧，Proxy到RedisServer的写命令平均执行延迟(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemCpuUtilProxy, startTime, endTime, monitor.RedisProxyInstance(profile.Dimensioninstanceid, profile.Dimensionpnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemCpuUtilProxy, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForProxyid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "RedisProxyCPU利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemInBandwidthUtilProxy, startTime, endTime, monitor.RedisProxyInstance(profile.Dimensioninstanceid, profile.Dimensionpnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemInBandwidthUtilProxy, err))
	}

	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForProxyid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "Proxy侧，Redis内网入带宽利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemOutBandwidthUtilProxy, startTime, endTime, monitor.RedisProxyInstance(profile.Dimensioninstanceid, profile.Dimensionpnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemOutBandwidthUtilProxy, err))
	}

	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForProxyid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "Proxy侧，Redis内网出带宽利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}
	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemLatencyWriteProxy, startTime, endTime, monitor.RedisProxyInstance(profile.Dimensioninstanceid, profile.Dimensionpnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemLatencyWriteProxy, err))
	}

	for _, instance := range *data {
		if instance.MetricValue > float64(100) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForProxyid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "Proxy侧，RedisProxy到RedisServer的写命令平均执行延迟(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}
	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemLatencyAvgProxy, startTime, endTime, monitor.RedisProxyInstance(profile.Dimensioninstanceid, profile.Dimensionpnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemLatencyAvgProxy, err))
	}

	for _, instance := range *data {
		if instance.MetricValue > float64(100) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForProxyid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "Proxy侧，RedisProxy到RedisServer的所有命令平均延迟(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemLatencyReadProxy, startTime, endTime, monitor.RedisProxyInstance(profile.Dimensioninstanceid, profile.Dimensionpnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemLatencyReadProxy, err))
	}

	for _, instance := range *data {
		if instance.MetricValue > float64(100) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForProxyid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "Proxy侧，RedisProxy到RedisServer的读命令平均执行延迟(ms)/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemCpuUtilNode, startTime, endTime, monitor.RedisNodeInstance(profile.Dimensioninstanceid, profile.Dimensionrnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemCpuUtilNode, err))
	}

	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForNodeid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "Redis节点侧，平均CPU利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemConnectionsUtilNode, startTime, endTime, monitor.RedisNodeInstance(profile.Dimensioninstanceid, profile.Dimensionrnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemConnectionsUtilNode, err))
	}

	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForNodeid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "Redis节点侧，连接数利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceRedismem, profile.MetricRedismemMemUtilNode, startTime, endTime, monitor.RedisNodeInstance(profile.Dimensioninstanceid, profile.Dimensionrnodeid, redislist))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceRedismem, profile.MetricRedismemMemUtilNode, err))
	}

	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			var instanceinfo *redis.Instance
			ids := strings.Split(instance.InstanceId, " ")
			instanceinfo = redis.InstanceInfo(ids[0], redislist)
			if instanceinfo == nil {
				instanceinfo = redis.InstanceInfoForNodeid(ids[1], redislist)
			}
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), ids[0]+"/"+ids[1], instanceinfo.Name, "Redis节点侧，内存利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	return nil
}
