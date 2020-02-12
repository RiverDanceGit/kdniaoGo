# kdniaoGo
快递鸟 SDK Go 语言版

## Requirements
Go 1.12 or above.

## Install
```shell
go get -u -x -v -insecure github.com/RiverDanceGit/kdniaoGo
```

### Configure Parameter
EBusinessID 和 AppKey 去 [快递鸟用户管理后台](http://kdniao.com/UserCenter/UserHome.aspx) 查看

### API Document
* [API 文档首页](http://www.kdniao.com/api-all)
* 快递公司编码 [《2019快递鸟接口支持快递公司编码.xlsx》](http://www.kdniao.com/file/2019%E5%BF%AB%E9%80%92%E9%B8%9F%E6%8E%A5%E5%8F%A3%E6%94%AF%E6%8C%81%E5%BF%AB%E9%80%92%E5%85%AC%E5%8F%B8%E7%BC%96%E7%A0%81.xlsx)

#### 下单类接口
* [电子面单API](http://www.kdniao.com/api-eorder)
* [上门取件API](http://www.kdniao.com/api-pickup)
* [预约取件API](http://www.kdniao.com/api-order)
* [代收货款API](http://www.kdniao.com/api-cod)

#### 查询类接口
* [即时查询API](http://www.kdniao.com/api-track)
* [物流跟踪API](http://www.kdniao.com/api-follow)
* [在途监控API](http://www.kdniao.com/api-monitor) (完成 SUBSCRIBE 和 REALTIME 接口)

#### 增值类接口
* [隐私快递API](http://www.kdniao.com/api-safemail)
* [智选物流API](http://www.kdniao.com/api-exprecommend)
* [物流短信API](http://www.kdniao.com/api-sms)
* [单号识别API](http://www.kdniao.com/api-recognise) (完成)
* [物流评价API](http://www.kdniao.com/api-evaluate)
* [实名寄递API](http://www.kdniao.com/api-verified)
