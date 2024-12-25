package profile

const (
	MetricCvmBaseCpuUsage = "BaseCpuUsage" // 基础CPU利用率通过宿主机采集上报，无须安装监控组件即可查看数据，子机高负载情况下仍可持续采集上报数据
	MetricCvmCpuUsage     = "CpuUsage"     // 服务器CPU利用率
	MetricCvmCpuLoadavg   = "CpuLoadavg"   // 服务器CPU一分钟平均负载，1分钟内正在使用和等待使用CPU的平均任务数（Windows机器无此指标）
	MetricCvmMemUsage     = "MemUsage"     // 服务器用户实际内存利用率，不包括缓冲区与系统缓存占用的内存，除去缓存、buffer和剩余，用户实际使用内存与总内存之比
	MetricCvmDiskUsage    = "CvmDiskUsage" // 磁盘已使用容量占总容量的百分比（所有磁盘)
	MetricCvmDiskIoAwait  = "DiskAwait"    // 硬盘IO等待时间
	MetricCvmDiskUtil     = "DiskUtil"     // 硬盘IO繁忙比率

	MetricRedismemCpuUtil               = "CpuUtil"               // Redis实例侧，平均CPU利用率
	MetricRedismemUtil                  = "MemUtil"               // Redis实例侧，平均内存利用率
	MetricRedismemConnectionsUtil       = "ConnectionsUtil"       // Redis实例侧，平均连接数利用率
	MetricRedismemLatencyAvg            = "LatencyAvg"            // Redis实例侧，Proxy到RedisServer的执行延迟平均值
	MetricRedismemLatencyRead           = "LatencyRead"           // Redis实例侧，Proxy到RedisServer的读命令平均执行延迟
	MetricRedismemLatencyWrite          = "LatencyWrite"          // Redis实例侧，Proxy到RedisServer的写命令平均执行延迟
	MetricRedismemCpuUtilProxy          = "CpuUtilProxy"          // RedisProxyCPU利用率
	MetricRedismemInBandwidthUtilProxy  = "InBandwidthUtilProxy"  // Proxy侧，Redis内网入流量实际使用和最大流量比
	MetricRedismemOutBandwidthUtilProxy = "OutBandwidthUtilProxy" // Proxy侧，Redis内网出流量实际使用和最大流量比
	MetricRedismemLatencyAvgProxy       = "LatencyAvgProxy"       // Proxy侧，RedisProxy到RedisServer的执行延迟平均值
	MetricRedismemLatencyReadProxy      = "LatencyReadProxy"      // Proxy侧，RedisProxy到RedisServer的读命令平均执行延迟
	MetricRedismemLatencyWriteProxy     = "LatencyWriteProxy"     // Proxy侧，RedisProxy到RedisServer的写命令平均执行延迟
	MetricRedismemCpuUtilNode           = "CpuUtilNode"           // Redis节点侧，平均CPU利用率
	MetricRedismemConnectionsUtilNode   = "ConnectionsUtilNode"   // Redis节点侧，连接数利用率
	MetricRedismemMemUtilNode           = "MemUtilNode"           // Redis节点侧，实际使用内存和申请总内存之比

	MetricMongodbDelay100               = "Delay100"               // MongoDB集群单位时间内成功请求延迟在100ms以上次数
	MetricMongodbAvgAllRequestDelay     = "AvgAllRequestDelay"     // MongoDB集群所有请求平均延迟
	MetricMongodbConnper                = "Connper"                // MongoDB集群连接数利用率
	MetricMongodbClusterDiskUsage       = "ClusterDiskUsage"       // MongoDB集群存储利用率
	MetricMongodbSlaveDelay             = "SlaveDelay"             // MongoDB主从单位时间内平均延迟
	MetricMongodbReplicaDiskUsage       = "ReplicaDiskUsage"       // MongoDB副本集容量利用率
	MetricMongodbCpuUsage               = "CpuUsage"               // MongoDB节点CPU利用率
	MetricMongodbMemUsage               = "MemUsage"               // MongoDB节点内存利用率
	MetricMongodbNodeSlavedelay         = "NodeSlavedelay"         // MongoDB从节点延迟
	MetricMongodbDiskUsage              = "DiskUsage"              // MongoDB节点磁盘利用率
	MetricMongodbNodeAvgAllRequestDelay = "NodeAvgAllRequestDelay" // MongoDB节点所有请求平均延迟
	MetricMongodbNodeAvgUpdateDelay     = "NodeAvgUpdateDelay"     // MongoDB节点更新命令延迟平均值
	MetricMongodbNodeAvgInsertDelay     = "NodeAvgInsertDelay"     // MongoDB节点插入命令延迟平均值
	MetricMongodbNodeAvgReadDelay       = "NodeAvgReadDelay"       // MongoDB节点读命令延迟平均值
	MetricMongodbNodeAvgAggregateDelay  = "NodeAvgAggregateDelay"  // MongoDB节点aggregate命令延迟平均值
	MetricMongodbNodeAvgCountDelay      = "NodeAvgCountDelay"      // MongoDB节点counts命令延迟平均值
	MetricMongodbNodeAvgGetmoreDelay    = "NodeAvgGetmoreDelay"    // MongoDB节点getmore命令延迟平均值
	MetricMongodbNodeAvgDeleteDelay     = "NodeAvgDeleteDelay"     // MongoDB节点delete请求延迟平均值
	MetricMongodbNodeAvgCommandDelay    = "NodeAvgCommandDelay"    // MongoDB节点command请求延迟平均值
	MetricMongodbNodeDelay100           = "NodeDelay100"           // MongoDB节点延迟超过100毫秒请求量

	MetricTDSQL_C_MysqlSlowQueries       = "SlowQueries"       // TDSQL-C_MySQL慢查询数
	MetricTDSQL_C_MysqlMemoryUserAte     = "MemoryUserAte"     // TDSQL-C_MySQL内存利用率
	MetricTDSQL_C_MysqlStorageUserate    = "StorageUserate"    // TDSQL-C_MySQL存储利用率
	MetricTDSQL_C_MysqlConnectionUserAte = "ConnectionUserAte" // TDSQL-C_MySQL连接数利用率
	MetricTDSQL_C_MysqlCpuUserAte        = "CpuUserAte"        // TDSQL-C_MySQLCPU利用率

	MetricNatgatewayConnsUsage           = "ConnsUsage"           // NAT网关网络连接数利用率
	MetricNatgatewayEgressbandwidthusage = "Egressbandwidthusage" // NAT网关外网出带宽利用率
	MetricNatgatewayWanInByteUsage       = "WanInByteUsage"       /// NAT网关外网入带宽利用率

	MetricVipIntraffic  = "VipIntraffic"  // 弹性公网IP入带宽，需要获取IP总带宽计算外网入带宽利用率，以及总利用率
	MetricVipOuttraffic = "VipOuttraffic" // 弹性公网IP出带宽，需要获取IP总带宽计算外网出带宽利用率，以及总利用率

	MetricPubliclbDropTotalConns      = "DropTotalConns"     // 公网CLB在统计粒度内，负载均衡或监听器上丢弃的连接数
	MetricPubliclbInDropBits          = "InDropBits"         // 公网CLB在统计粒度内，客户端通过外网访问负载均衡时丢弃的带宽
	MetricPubliclbOutDropBits         = "OutDropBits"        // 公网CLB在统计粒度内，负载均衡访问外网时丢弃的带宽
	MetricPubliclbDropQps             = "DropQps"            // 公网CLB在统计粒度内，负载均衡或监听器上丢弃的请求数
	MetricPubliclbIntrafficVipRatio   = "IntrafficVipRatio"  // 公网CLB在统计粒度内入带宽利用率
	MetricPubliclbOuttrafficVipRatio  = "OuttrafficVipRatio" // 公网CLB在统计粒度内出带宽利用率
	MetricPubliclbClbHttp5xx          = "ClbHttp5xx"         // 公网CLB在统计粒度内，负载均衡返回5xx状态码的个数（负载均衡和后端服务器返回码之和),此指标为七层监听器独有指标
	MetricPubliclbHttp5xx             = "Http5xx"            // 公网CLB在统计粒度内，后端服务器返回5xx状态码的个数,此指标为七层监听器独有指标
	MetricPubliclbConcurConnVipRatio  = "ConcurConnVipRatio" // 公网CLB并发连接数利用率
	MetricPubliclbNewConnVipRatio     = "NewConnVipRatio"    // 公网CLB新建连接数利用率
	MetricPrivatelbDropTotalConns     = "DropTotalConns"     // 内网CLB在统计粒度内，负载均衡或监听器上丢弃的连接数
	MetricPrivatelbInDropBits         = "InDropBits"         // 内网CLB丢弃流入数据包
	MetricPrivatelbOutDropBits        = "OutDropBits"        // 内网CLB丢弃流出数据包
	MetricPrivatelbDropQps            = "DropQps"            // 内网CLB在统计粒度内，负载均衡或监听器上丢弃的请求数
	MetricPrivatelbIntrafficVipRatio  = "IntrafficVipRatio"  // 内网CLB入带宽利用率
	MetricPrivatelbOuttrafficVipRatio = "OuttrafficVipRatio" // 内网CLB出带宽利用率
	MetricPrivatelbClbHttp5xx         = "ClbHttp5xx"         // 内网CLB在统计粒度内，负载均衡返回 5xx 状态码的个数（负载均衡和后端服务器返回码之和）
	MetricPrivatelbHttp5xx            = "Http5xx"            // 内网CLB在统计粒度内，后端服务器返回 5xx 状态码的个数，此指标为七层监听器独有指标。
	MetricPrivatelbConcurConnVipRatio = "ConcurConnVipRatio" // 内网CLB并发连接数利用率
	MetricPrivatelbNewConnVipRatio    = "NewConnVipRatio"    // 内网CLB新建连接数利用率

)

type MetricTke2 struct {
	// 监控指标名称
	Name string

	// 监控指标
	Metric string

	// 哪种资源的监控指标
	MetricType string

	// 单位
	Unit string

	// 操作符
	Operator string

	// 触发值
	TriggerValue float64

	// 告警提示词
	Prompt string
}

var MetricTke2Node = []MetricTke2{
	{
		Name:         "CPU利用率",
		Metric:       "K8sNodeCpuUsage",
		MetricType:   MetricTypeK8sNode,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 70.00,
		Prompt:       "请评估节点是否达到扩容水平或节点是否存在大量消耗CPU的Pod",
	},
	{
		Name:         "Node状态",
		Metric:       "K8sNodeStatusReady",
		MetricType:   MetricTypeK8sNode,
		Unit:         UnitCount,
		Operator:     OperatorEqual,
		TriggerValue: 0,
		Prompt:       "节点下线，请关注！！！",
	},
	{
		Name:         "内存利用率",
		Metric:       "K8sNodeMemUsage",
		MetricType:   MetricTypeK8sNode,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 80.00,
		Prompt:       "请评估节点是否达到扩容水平或节点是否存在大量消耗内存的Pod",
	},
	{
		Name:         "节点上Pod重启次数",
		Metric:       "K8sNodePodRestartTotal",
		MetricType:   MetricTypeK8sNode,
		Unit:         UnitCount,
		Operator:     OperatorGreaterthan,
		TriggerValue: 10.00,
		Prompt:       "节点发生异常事件，导致大量Pod异常重启，请检查资源集群状态以及资源情况",
	},
}

var MetricTke2Pod = []MetricTke2{
	{
		Name:         "CPU利用率（占limit）",
		Metric:       "K8sPodRateCpuCoreUsedLimit",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 85.00,
		Prompt:       "Pod的CPU利用率(占limit)升高，请关注！！！",
	},
	{
		Name:         "内存利用率（占节点，不包含cache）",
		Metric:       "K8sPodRateMemNoCacheNode",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 30.00,
		Prompt:       "Pod的内存利用率(占节点，不包含cache)消耗过大，请关注！！！",
	},
	{
		Name:         "内存使用量（占Request）",
		Metric:       "K8sPodRateMemUsageRequest",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 1000.00,
		Prompt:       "Pod内存利用率(包含cache)已达Request上限，请判断是否调整资源配置！！！",
	},
	{
		Name:         "CPU利用率（占节点）",
		Metric:       "K8sPodRateCpuCoreUsedNode",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 40.00,
		Prompt:       "Pod消耗过多节点CPU，请关注！！！",
	},
	{
		Name:         "内存利用率（占request，不含cache）",
		Metric:       "K8sPodRateMemNoCacheRequest",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 1000.00,
		Prompt:       "Pod内存利用率(不含cache)已达Request上限，请判断是否调整资源配置！！！",
	},
	{
		Name:         "Pod重启次数",
		Metric:       "K8sPodRestartTotal",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitCount,
		Operator:     OperatorGreaterthan,
		TriggerValue: 0.00,
		Prompt:       "Pod发生重启事件，请关注！！！",
	},
	{
		Name:         "CPU利用率（占request）",
		Metric:       "K8sPodRateCpuCoreUsedRequest",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 500.00,
		Prompt:       "PodCPU利用率已达Request上限，请判断是否调整资源配置！！！",
	},
	{
		Name:         "内存利用率（占limit,包含cache）",
		Metric:       "K8sPodRateMemUsageLimit",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 90.00,
		Prompt:       "Pod的内存利用率(占limit，含cache)升高，请关注！！！",
	},
	{
		Name:         "Pod_Ready状态",
		Metric:       "K8sPodStatusReady",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorEqual,
		TriggerValue: 0,
		Prompt:       "Pod状态异常，请关注！！！",
	},
	{
		Name:         "内存利用率(占limit，不包含cache)",
		Metric:       "K8sPodRateMemNoCacheLimit",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorEqual,
		TriggerValue: 85,
		Prompt:       "Pod的内存利用率(占limit，不含cache)升高，请关注！！！",
	},
	{
		Name:         "内存利用率（占节点）",
		Metric:       "K8sPodRateMemUsageNode",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 40,
		Prompt:       "Pod的内存利用率(占节点，包含cache)消耗过大，请关注！！！",
	},
	// CPU利用率（占Pod规格）只有超级节点上的pod有该指标
	//{
	//	Name:         "CPU使用率（占Pod规格）",
	//	Metric:       "K8sPodRateCpuCoreUsedResource",
	//	MetricType:   MetricTypeK8sPod,
	//	Unit:         UnitPercentage,
	//	Operator:     OperatorGreaterthan,
	//	TriggerValue: 80,
	//},
	// 内存利用率（占Pod规格）只有超级节点上的pod有该指标
	//{
	//	Name:         "内存利用率（占Pod规格，不包含cache）",
	//	Metric:       "K8sPodRateMemNoCacheResource",
	//	MetricType:   MetricTypeK8sPod,
	//	Unit:         UnitPercentage,
	//	Operator:     OperatorGreaterthan,
	//	TriggerValue: 80,
	//},
	{
		Name:         "rootfs空间使用率",
		Metric:       "K8sPodDiskUsage",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 80,
		Prompt:       "rootfs空间使用率消耗过大，请关注！！！",
	},
	{
		Name:         "内存利用率（working_set占limit）",
		Metric:       "K8sPodRateMemWorkingSetLimit",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 90,
		Prompt:       "Pod的内存利用率(working_set占limit)升高，请判断是否调整资源配置",
	},
	{
		Name:         "内存利用率（working_set占节点）",
		Metric:       "K8sPodRateMemWorkingSetNode",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 40.00,
		Prompt:       "Pod的内存利用率(working_set占节点)消耗过大，请关注！！！",
	},
	{
		Name:         "内存利用率（working_set占request）",
		Metric:       "K8sPodRateMemWorkingSetRequest",
		MetricType:   MetricTypeK8sPod,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 1000,
		Prompt:       "Pod的内存利用率(working_set占request)已达Request上限，请判断是否调整资源配置！！！",
	},
	// 内存利用率（working_set占Pod规格）只有超级节点上的pod有该指标
	//{
	//	Name:         "内存利用率（working_set占Pod规格）",
	//	Metric:       "K8sPodRateMemWorkingSetResource",
	//	MetricType:   MetricTypeK8sPod,
	//	Unit:         UnitPercentage,
	//	Operator:     OperatorGreaterthan,
	//	TriggerValue: 80,
	//},
}

var MetricTke2Workload = []MetricTke2{
	{
		Name:         "工作负载异常",
		Metric:       "K8sWorkloadAbnormal",
		MetricType:   MetricTypeK8sWorkload,
		Unit:         UnitCount,
		Operator:     OperatorEqual,
		TriggerValue: 1,
		Prompt:       "工作负载状态异常，请关注！！！",
	},
	{
		Name:         "CPU利用率",
		Metric:       "K8sWorkloadRateCpuCoreUsedCluster",
		MetricType:   MetricTypeK8sWorkload,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 85,
		Prompt:       "工作负载CPU利用率升高,请关注！！！",
	},
	{
		Name:         "内存利用率（不含cache）",
		Metric:       "K8sWorkloadRateMemNoCacheCluster",
		MetricType:   MetricTypeK8sWorkload,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 85.00,
		Prompt:       "工作负载内存利用率(不含cache)升高,请关注！！！",
	},
	{
		Name:         "Pod重启次数",
		Metric:       "K8sWorkloadPodRestartTotal",
		MetricType:   MetricTypeK8sWorkload,
		Unit:         UnitCount,
		Operator:     OperatorGreaterthan,
		TriggerValue: 1.00,
		Prompt:       "工作负载异常重启,请关注！！！",
	},
	{
		Name:         "内存利用率",
		Metric:       "K8sWorkloadRateMemUsageBytesCluster",
		MetricType:   MetricTypeK8sWorkload,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 85.00,
		Prompt:       "工作负载内存利用率升高，请关注！！！",
	},
}

var MetricTke2Cluster = []MetricTke2{
	{
		Name:         "集群CPU分配率",
		Metric:       "K8sClusterRateCpuCoreRequestCluster",
		MetricType:   MetricTypeK8sCluster,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 300.00,
		Prompt:       "集群CPU分配率过高，请关注！！！",
	},
	{
		Name:         "集群内存分配率",
		Metric:       "K8sClusterRateMemRequestBytesCluster",
		MetricType:   MetricTypeK8sCluster,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 300.00,
		Prompt:       "集群内存分配率过高，请关注！！！",
	},
	{
		Name:         "集群CPU利用率",
		Metric:       "K8sClusterRateCpuCoreUsedCluster",
		MetricType:   MetricTypeK8sCluster,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 70.00,
		Prompt:       "请判断集群是否需要扩容或集群是否存在大量消耗CPU的Pod",
	},
	{
		Name:         "集群内存利用率",
		Metric:       "K8sClusterRateMemUsageBytesCluster",
		MetricType:   MetricTypeK8sCluster,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 70.00,
		Prompt:       "请判断集群是否需要扩容或集群是否存在大量消耗内存的Pod，集群内存利用率升高",
	},
	{
		Name:         "集群内存利用率（不含cache）",
		Metric:       "K8sClusterRateMemNoCacheBytesCluster",
		MetricType:   MetricTypeK8sCluster,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 85.00,
		Prompt:       "请判断集群是否需要扩容或集群是否存在大量消耗内存的Pod，集群内存利用率(不含cache)升高",
	},
	{
		Name:         "集群可分配的Pod数量",
		Metric:       "K8sClusterAllocatablePodsTotal",
		MetricType:   MetricTypeK8sCluster,
		Unit:         UnitCount,
		Operator:     OperatorLessthan,
		TriggerValue: 50,
		Prompt:       "集群可分配Pod数偏低，请考虑集群是否需要扩容或是否存在过多副本的应用",
	},
	{
		Name:         "内存使用率（working_set）",
		Metric:       "K8sClusterRateMemWorkingSetCluster",
		MetricType:   MetricTypeK8sCluster,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 70,
		Prompt:       "集群内存使用率(working_set)升高请考虑集群是否需要扩容或是否存在过多副本的应用",
	},
}

var MetricTke2Pvc = []MetricTke2{
	{
		Name:         "PVC云盘利用率",
		Metric:       "K8sPvcDiskUsage",
		MetricType:   MetricTypeK8sPvc,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 80.00,
		Prompt:       "PVC云盘利用率升高，请关注！！！",
	},
}

var MetricTke2Apiserver = []MetricTke2{
	{
		Name:         "Apiserver CPU利用率",
		Metric:       "K8sComponentApiserverRateCpuCoreUsageLimit",
		MetricType:   MetricTypeK8sApiserver,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 80.00,
		Prompt:       "ApiserverCPU利用率升高，请关注！！！",
	},
	{
		Name:         "Apiserver 内存利用率",
		Metric:       "K8sComponentApiserverRateMemUsageLimit",
		MetricType:   MetricTypeK8sApiserver,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 80.00,
		Prompt:       "Apiserver内存利用率升高，请关注！！！",
	},
	{
		Name:         "Apiserver List请求的延迟P99",
		Metric:       "K8sComponentApiserverRequestListDurationSecondsP99",
		MetricType:   MetricTypeK8sApiserver,
		Unit:         UnitCount,
		Operator:     OperatorGreaterthan,
		TriggerValue: 10.00,
		Prompt:       "ApiserverList请求延迟(s)过大，请关注！！！",
	},
	{
		Name:         "Apiserver 写请求延迟P99",
		Metric:       "K8sComponentApiserverRequestMutatingDurationSecondsP99",
		MetricType:   MetricTypeK8sApiserver,
		Unit:         UnitCount,
		Operator:     OperatorGreaterthan,
		TriggerValue: 10.00,
		Prompt:       "Apiserver写请求延迟(s)过大，请关注！！！",
	},
	{
		Name:         "Apiserver Get请求延迟P99",
		Metric:       "K8sComponentApiserverRequestNonListDurationSecondsP99",
		MetricType:   MetricTypeK8sApiserver,
		Unit:         UnitCount,
		Operator:     OperatorGreaterthan,
		TriggerValue: 10.00,
		Prompt:       "ApiserverGet请求延迟(s)过大，请关注！！！",
	},
}

var MetricTke2Container = []MetricTke2{
	{
		Name:         "ContainerCPU利用率（占节点）",
		Metric:       "K8sContainerRateCpuCoreUsedNode",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 30.00,
		Prompt:       "容器消耗节点过多CPU，请关注！！！",
	},
	{
		Name:         "Container内存利用率（占节点,不包含cache）",
		Metric:       "K8sContainerRateMemNoCacheNode",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 40.00,
		Prompt:       "容器消耗节点过多内存(不包含cache)，请关注！！！",
	},
	{
		Name:         "Container内存利用率（占节点）",
		Metric:       "K8sContainerRateMemUsageNode",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 30.00,
		Prompt:       "容器消耗节点过多内存，请关注！！！",
	},
	{
		Name:         "ContainerCPU利用率（占Request）",
		Metric:       "K8sContainerRateCpuCoreUsedRequest",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 500.00,
		Prompt:       "容器消耗的CPU达到request上限，请考虑调整资源配置！！！",
	},
	{
		Name:         "Container内存利用率（占Request，不含cache）",
		Metric:       "K8sContainerRateMemNoCacheRequest",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 1000.00,
		Prompt:       "容器消耗的内存(不含cache)达到request上限，请考虑程序异常或调整资源配置！！！",
	},
	{
		Name:         "Container内存利用率（占Request）",
		Metric:       "K8sContainerRateMemUsageRequest",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 1000.00,
		Prompt:       "容器消耗的内存(含cache)达到request上限，请考虑调整资源配置！！！",
	},
	{
		Name:         "ContainerCPU利用率（占limit）",
		Metric:       "K8sContainerRateCpuCoreUsedLimit",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 85.00,
		Prompt:       "容器消耗的CPU即将达到limit上限，请考虑程序异常或调整资源配置！！！",
	},
	{
		Name:         "Container内存利用率（占Limit，不含cache）",
		Metric:       "K8sContainerRateMemNoCacheLimit",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 85.00,
		Prompt:       "容器消耗的内存(不含cache)即将达到limit上限，请考虑程序异常或调整资源配置！！！",
	},
	{
		Name:         "Container内存利用率（占limit）",
		Metric:       "K8sContainerRateMemUsageLimit",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 90.00,
		Prompt:       "容器消耗的内存(含cache)即将达到limit上限，请考虑程序异常或调整资源配置！！！",
	},
	{
		Name:         "Container内存利用率（working_set占limit）",
		Metric:       "K8sContainerRateMemWorkingSetLimit",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 90.00,
		Prompt:       "容器(working_set占limit)消耗的内存即将达到limit上限，请考虑程序异常或调整资源配置！！！",
	},
	{
		Name:         "Container内存利用率（working_set占节点）",
		Metric:       "K8sContainerRateMemWorkingSetNode",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 30.00,
		Prompt:       "容器(working_set占节点)消耗节点过多内存，请考虑程序异常或调整资源配置！！！",
	},
	{
		Name:         "Container内存利用率（working_set占request）",
		Metric:       "K8sContainerRateMemWorkingSetRequest",
		MetricType:   MetricTypeK8sContainer,
		Unit:         UnitPercentage,
		Operator:     OperatorGreaterthan,
		TriggerValue: 1000.00,
		Prompt:       "容器(working_set占request)消耗的内存达到request上限，请考虑调整资源配置！！！",
	},
}
