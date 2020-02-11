package enum

const (
	REQUEST_TYPE_EORDER              = "1007"    // 电子面单API-电子面单接口
	REQUEST_TYPE_EORDER_CANCEL       = "1147"    // 电子面单API-订单取消接口
	REQUEST_TYPE_EORDER_QUERY        = "1127"    // 电子面单API-单号余量查询接口,客户号申请接口
	REQUEST_TYPE_EORDER_PUSH         = "1117"    // 电子面单API-客户号推送接口
	REQUEST_TYPE_PICKUP              = "1801"    // 上门取件API
	REQUEST_TYPE_ORDER               = "1001"    // 预约取件API
	REQUEST_TYPE_COD_REGITER_ACCOUNT = "9001"    // 代收货款API-用户注册接口
	REQUEST_TYPE_COD_UPDATE_ACCOUNT  = "CMD1002" // 代收货款API-更新用户信息
	REQUEST_TYPE_COD_QUERY_ACCOUNT   = "CMD1003" // 代收货款API-查询用户信息
	REQUEST_TYPE_COD_SUBMIT_BANK     = "CMD1009" // 代收货款API-提交返款银行信息
	REQUEST_TYPE_COD_QUERY_BANK      = "CMD1008" // 代收货款API-查询返款银行信息
	REQUEST_TYPE_COD_QUERY_AMOUNT    = "CMD1014" // 代收货款API-查询用户额度
	REQUEST_TYPE_COD_PAY_FIRST       = "CMD1004" // 代收货款API-垫付业务申请
	REQUEST_TYPE_COD_QUIT            = "CMD1005" // 代收货款API-直退业务申请
	REQUEST_TYPE_COD_APPLY           = "CMD1006" // 代收货款API-普通代收货款申请
	REQUEST_TYPE_COD_QUERY_APPLY     = "CMD1007" // 代收货款API-查询服务申请状态
	REQUEST_TYPE_COD_QUERY_ORDER     = "CMD1010" // 代收货款API-获取订单货款状态
	REQUEST_TYPE_TRACK               = "1002"    // 即时查询API
	REQUEST_TYPE_FOLLOW_SUBSCRIBE    = "1008"    // 物流跟踪API-订阅接口
	REQUEST_TYPE_FOLLOW_TRACK        = "101"     // 物流跟踪API-轨迹查询结果(推送接口,商户实现)
	REQUEST_TYPE_FOLLOW_STATE        = "107"     // 物流跟踪API-货款状态(推送接口,商户实现)
	REQUEST_TYPE_MONITOR_REALTIME    = "8001"    // 在途监控API-即时查询(增值版)接口
	REQUEST_TYPE_MONITOR_SUBSCRIBE   = "8008"    // 在途监控API-订阅(增值版)接口
	REQUEST_TYPE_MONITOR_TRACK       = "102"     // 在途监控API-轨迹查询结果(推送接口,增值版,商户实现)
	REQUEST_TYPE_MONITOR_STATE       = "107"     // 在途监控API-货款状态(推送接口,增值版,商户实现)
	REQUEST_TYPE_SAFEMAIL            = "3001"    // 隐私快递API
	REQUEST_TYPE_EXPRECOMMEND        = "2006"    // 智选物流API-智选物流
	REQUEST_TYPE_EXPRECOMMEND_TPL    = "2004"    // 智选物流API-模板导入
	REQUEST_TYPE_SMS_TPL             = "8102"    // 物流短信API-短信模版
	REQUEST_TYPE_SMS_DIY             = "8101"    // 物流短信API-自定义类短信
	REQUEST_TYPE_SMS_BLACK           = "8103"    // 物流短信API-短信黑名单
	REQUEST_TYPE_RECOGNISE           = "2002"    // 单号识别API
	REQUEST_TYPE_EVALUATE_SUBMIT     = "1011"    // 物流评价API-提交投诉
	REQUEST_TYPE_EVALUATE_COMPANY    = "1013"    // 物流评价API-获取企业列表
	REQUEST_TYPE_EVALUATE_SCORE      = "1012"    // 物流评价API-获取企业平均分
	REQUEST_TYPE_VERIFIED            = "1021"    // 实名寄递API
)
