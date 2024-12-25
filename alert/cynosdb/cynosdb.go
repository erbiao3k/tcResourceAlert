package cynosdb

import (
	"errors"
	"fmt"
	"tcResourceAlert/msg_sender"
	"tcResourceAlert/profile"
	"tcResourceAlert/resource/cynosdb"
	"tcResourceAlert/resource/monitor"
	"tcResourceAlert/utils"
)

func Alert(startTime, endTime string) error {
	dblist, err := cynosdb.GetInstancelist()
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstancelistErrmsg, profile.NamespaceCynosdbmysql, err))
	}

	var dbIds []string
	for _, i := range *dblist {
		dbIds = append(dbIds, i.Id)
	}

	data, err := monitor.GetMonitorData(profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlSlowQueries, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, dbIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlSlowQueries, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(100) {
			instanceinfo := cynosdb.InstanceInfo(instance.InstanceId, dblist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "TDSQL-C_MySQL慢查询个数/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlMemoryUserAte, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, dbIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlMemoryUserAte, err))
	}
	for _, instance := range *data {
		if instance.MetricValue > float64(90) {
			instanceinfo := cynosdb.InstanceInfo(instance.InstanceId, dblist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "TDSQL-C_MySQL内存利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlStorageUserate, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, dbIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlStorageUserate, err))
	}

	for _, instance := range *data {
		if instance.MetricValue >= float64(90) {
			instanceinfo := cynosdb.InstanceInfo(instance.InstanceId, dblist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "TDSQL-C_MySQL存储利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlConnectionUserAte, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, dbIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlConnectionUserAte, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(90) {
			instanceinfo := cynosdb.InstanceInfo(instance.InstanceId, dblist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "TDSQL-C_MySQL连接数利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	data, err = monitor.GetMonitorData(profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlCpuUserAte, startTime, endTime, monitor.Instance(profile.DimensionInstanceId, dbIds))
	if err != nil {
		return errors.New(fmt.Sprintf(profile.InstanceMetricErrmsg, profile.NamespaceCynosdbmysql, profile.MetricTDSQL_C_MysqlCpuUserAte, err))
	}
	for _, instance := range *data {
		if instance.MetricValue >= float64(90) {
			instanceinfo := cynosdb.InstanceInfo(instance.InstanceId, dblist)
			msg := fmt.Sprintf(profile.AlertBasicMsg, utils.Nowtime(), instanceinfo.Id, instanceinfo.Name, "TDSQL-C_MySQLCPU利用率/"+instance.MetricName, utils.Float64Str(instance.MetricValue))
			msg_sender.MsgSender(msg)
		}
	}

	return nil
}
