package main

import (
	"log"
	"tcResourceAlert/alert/cynosdb"
	"tcResourceAlert/utils"
	"time"
)

func main() {
	for {
		var err error

		startTime, endTime := utils.RangeTime(2, 4)
		log.Println("开始获取资源的监控数据，达到阈值会有告警哦～")
		log.Printf("当前获取的监控数据时间范围是：%s~%s", startTime, endTime)

		err = cynosdb.Alert(startTime, endTime)
		if err != nil {
			log.Println("【TDSQL-C MYSQL资源告警异常】", err)
		}

		log.Println("本周期监控数据已获取完成，请注意是否有告警信息哦～")

		time.Sleep(time.Minute * 5)

		utils.SlowRun()

	}
}
