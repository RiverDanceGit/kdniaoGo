package request

import (
	"encoding/json"
	"github.com/RiverDanceGit/kdniaoGo/enum"
)

func NewRecogniseRequest() RecogniseRequest {
	req := RecogniseRequest{}
	req.SetRequestType(enum.REQUEST_TYPE_RECOGNISE)
	req.SetDataType(enum.DATA_TYPE)
	return req
}

type RecogniseRequest struct {
	KdniaoRequest
	LogisticCode string `json:"LogisticCode"`
}

func (req *RecogniseRequest) SetLogisticCode(logisticCode string) *RecogniseRequest {
	req.LogisticCode = logisticCode
	return req
}

func (req RecogniseRequest) GetLogisticCode() string {
	return req.LogisticCode
}

func (req *RecogniseRequest) UpdateRequestData() *RecogniseRequest {
	req.requestData = req.ToJson()
	return req
}

func (req RecogniseRequest) ToJson() string {
	str, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(str)
}
