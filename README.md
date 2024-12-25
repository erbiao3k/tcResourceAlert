# tcResourceAlert
腾讯云常用资源的监控程序，解决腾讯云可观测无法完全自动化的问题，数据从腾讯云端获取。

## 已支持的资源
- cbs
- cvm
- cynosdb
- mongodb
- redis
- tke2

## 配置
- 环境变量`SecretIdKey`，id在前，key在后，中间用英文`,`隔开
- 环境变量`WecomRobot`，多个企业微信群机器人使用英文`,`隔开
