package request

import (
	"encoding/json"
	"github.com/RiverDanceGit/kdniaoGo/enum"
)

func NewMonitorRealtimeRequest() MonitorRealtimeRequest {
	req := MonitorRealtimeRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_MONITOR_REALTIME)
	req.SetDataType(enum.DATA_TYPE_JSON)
	return req
}

type MonitorRealtimeRequest struct {
	KdniaoRequest
	LogisticCode string `json:"LogisticCode"`        // 物流单号
	ShipperCode  string `json:"ShipperCode"`         // 快递公司编码
	OrderCode    string `json:"OrderCode,omitempty"` // 订单编号
}

func (req *MonitorRealtimeRequest) SetLogisticCode(logisticCode string) *MonitorRealtimeRequest {
	req.LogisticCode = logisticCode
	return req
}

func (req MonitorRealtimeRequest) GetLogisticCode() string {
	return req.LogisticCode
}

func (req *MonitorRealtimeRequest) SetShipperCode(shipperCode string) *MonitorRealtimeRequest {
	req.ShipperCode = shipperCode
	return req
}

func (req MonitorRealtimeRequest) GetShipperCode() string {
	return req.ShipperCode
}

func (req *MonitorRealtimeRequest) SetOrderCode(orderCode string) *MonitorRealtimeRequest {
	req.OrderCode = orderCode
	return req
}

func (req MonitorRealtimeRequest) GetOrderCode() string {
	return req.OrderCode
}

func (req *MonitorRealtimeRequest) UpdateRequestData() *MonitorRealtimeRequest {
	req.requestData = req.ToJson()
	return req
}

func (req MonitorRealtimeRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}
