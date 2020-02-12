package request

import (
	"encoding/json"
	"github.com/RiverDanceGit/kdniaoGo/enum"
)

func NewMonitorSubscribeRequest() MonitorSubscribeRequest {
	req := MonitorSubscribeRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_MONITOR_SUBSCRIBE)
	req.SetDataType(enum.DATA_TYPE_JSON)
	return req
}

type MonitorSubscribeRequest struct {
	KdniaoRequest
	CallBack     string `json:"Callback,omitempty"`     // 用户自定义回调信息
	MemberId     string `json:"MemberID,omitempty"`     // 会员标识
	WareHouseId  string `json:"WareHouseID,omitempty"`  // 仓库标识
	CustomerName string `json:"CustomerName,omitempty"` // 电子面单客户号
	CustomerPwd  string `json:"CustomerPwd,omitempty"`  // 电子面单密码
	SendSite     string `json:"SendSite,omitempty"`     // 收件网点标识(名称)
	ShipperCode  string `json:"ShipperCode"`            // 快递公司编码
	LogisticCode string `json:"LogisticCode"`           // 快递单号
	OrderCode    string `json:"OrderCode,omitempty"`    // 订单编号
	MonthCode    string `json:"MonthCode,omitempty"`    // 月结编码
	PayType      string `json:"PayType,omitempty"`      // 邮费支付方式: 1-现付，2-到付，3-月结，4-第三方支付)
	ExpType      string `json:"ExpType,omitempty"`      // 快递类型：1-标准快件
	Cost         string `json:"Cost,omitempty"`         // 快递运费
	OtherCost    string `json:"OtherCost,omitempty"`    // 快递运费
	//Receiver      MonitorSubscribeRequestReceiver    `json:"Receiver,omitempty"`
	///Sender        MonitorSubscribeRequestSender      `json:"Sender,omitempty"`
	IsNotice      string `json:"IsNotice,omitempty"`  // 是否通知快递员上门揽件：0-通知；1-不通知；不填则
	StartDate     string `json:"StartDate,omitempty"` // 上门揽件时间段，格式：YYYY-MM-DD HH24:MM:SS
	EndDate       string `json:"EndDate,omitempty"`
	Weight        string `json:"Weight,omitempty"`        // 包裹总重量kg
	Quantity      string `json:"Quantity,omitempty"`      // 包裹数，一个包裹对应一个运单号，如果是大于1个包裹，返回则按照子母件的方式返回母运单号和子运单号
	Volume        string `json:"Volume,omitempty"`        // 包裹总体积m3
	Remark        string `json:"Remark,omitempty"`        // 备注
	IsSendMessage string `json:"IsSendMessage,omitempty"` // 是否订阅短信：0-不需要；1-需要
	//AddService    MonitorSubscribeRequestAddService  `json:"AddService,omitempty"`
	//Commodity     []MonitorSubscribeRequestCommodity `json:"Commodity,omitempty"`
}

type MonitorSubscribeRequestReceiver struct {
	Company      string `json:"Company,omitempty"` // 收件人公司
	Name         string `json:"Name"`              // 收件人
	Tel          string `json:"Tel,omitempty"`     // 电话与手机，必填一个
	Mobile       string `json:"Mobile,omitempty"`
	PostCode     string `json:"PostCode,omitempty"` // 收件地邮编 (ShipperCode为 EMS、YZPY 时必填)
	ProvinceName string `json:"ProvinceName"`       // 收件省 (如广东省，不要缺少“省”；如是直辖市，请直接传北京、上海等；如是自治区，请直接传广西壮族自治区等)
	CityName     string `json:"CityName"`           // 收件市 (如深圳市，不要缺少“市”)
	ExpAreaName  string `json:"ExpAreaName"`        // 收件区/县(如福田区，不要缺少“区”或“县”)
	Address      string `json:"Address"`            // 收件人详细地址
}

type MonitorSubscribeRequestSender struct {
	Company      string `json:"Company,omitempty"` // 发件人公司
	Name         string `json:"Name"`              // 发件人
	Tel          string `json:"Tel,omitempty"`     // 电话与手机，必填一个
	Mobile       string `json:"Mobile,omitempty"`
	PostCode     string `json:"PostCode,omitempty"` // 发件地邮编 (ShipperCode为 EMS、YZPY 时必填)
	ProvinceName string `json:"ProvinceName"`       // 发件省 (如广东省，不要缺少“省”；如是直辖市，请直接传北京、上海等；如是自治区，请直接传广西壮族自治区等)
	CityName     string `json:"CityName"`           // 发件市 (如深圳市，不要缺少“市”)
	ExpAreaName  string `json:"ExpAreaName"`        // 发件区/县(如福田区，不要缺少“区”或“县”)
	Address      string `json:"Address"`            // 发件详细地址
}

type MonitorSubscribeRequestAddService struct {
	Name       string `json:"Name,omitempty"`       // 增值服务名称
	Value      string `json:"Value,omitempty"`      // 增值服务值
	CustomerId string `json:"CustomerID,omitempty"` // 客户标识
}

type MonitorSubscribeRequestCommodity struct {
	GoodsName     string `json:"GoodsName"`     // 商品名称
	GoodsCode     string `json:"GoodsCode"`     // 商品编码
	Goodsquantity string `json:"Goodsquantity"` // 商品件数
	GoodsPrice    string `json:"GoodsPrice"`    // 商品价格
	GoodsWeight   string `json:"GoodsWeight"`   // 商品重量kg
	GoodsDesc     string `json:"GoodsDesc"`     // 商品描述
	GoodsVol      string `json:"GoodsVol"`      // 商品体积m3
}

func (req *MonitorSubscribeRequest) SetLogisticCode(logisticCode string) *MonitorSubscribeRequest {
	req.LogisticCode = logisticCode
	return req
}

func (req MonitorSubscribeRequest) GetLogisticCode() string {
	return req.LogisticCode
}

func (req *MonitorSubscribeRequest) SetShipperCode(shipperCode string) *MonitorSubscribeRequest {
	req.ShipperCode = shipperCode
	return req
}

func (req MonitorSubscribeRequest) GetShipperCode() string {
	return req.ShipperCode
}

func (req *MonitorSubscribeRequest) SetCustomerName(customerName string) *MonitorSubscribeRequest {
	req.CustomerName = customerName
	return req
}

func (req MonitorSubscribeRequest) GetCustomerName() string {
	return req.CustomerName
}

func (req *MonitorSubscribeRequest) UpdateRequestData() *MonitorSubscribeRequest {
	req.requestData = req.ToJson()
	return req
}

func (req MonitorSubscribeRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}
