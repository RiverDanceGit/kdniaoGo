package request

import (
	"encoding/base64"
	"github.com/RiverDanceGit/kdniaoGo"
	"github.com/RiverDanceGit/kdniaoGo/util"
	"net/url"
)

type KdniaoRequest struct {
	config      kdniaoGo.KdniaoConfig
	requestType string
	requestData string
	dataType    string
}

func (req KdniaoRequest) ToValues() url.Values {
	values := make(url.Values)
	values.Set("EBusinessID", req.GetConfig().GetEBusinessId())
	values.Set("RequestType", req.GetRequestType())
	values.Set("RequestData", req.GetRequestData())
	values.Set("DataSign", req.GetDataSign())
	values.Set("DataType", req.GetDataType())
	return values
}

func (req *KdniaoRequest) SetConfig(config kdniaoGo.KdniaoConfig) *KdniaoRequest {
	req.config = config
	return req
}

func (req KdniaoRequest) GetConfig() kdniaoGo.KdniaoConfig {
	return req.config
}

func (req *KdniaoRequest) SetRequestType(requestType string) *KdniaoRequest {
	req.requestType = requestType
	return req
}

func (req KdniaoRequest) GetRequestType() string {
	return req.requestType
}

func (req *KdniaoRequest) SetRequestData(requestData string) *KdniaoRequest {
	req.requestData = requestData
	return req
}

func (req KdniaoRequest) GetRequestData() string {
	return req.requestData
}

func (req *KdniaoRequest) SetDataType(dataType string) *KdniaoRequest {
	req.dataType = dataType
	return req
}

func (req KdniaoRequest) GetDataType() string {
	return req.dataType
}

func (req KdniaoRequest) GetDataSign() string {
	str := util.Md5(req.GetRequestData() + req.GetConfig().GetAppKey())
	return base64.StdEncoding.EncodeToString([]byte(str))
}
