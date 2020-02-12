package response

type RecogniseResponse struct {
	Success      bool                        `json:"Success"`
	Code         string                      `json:"Code"`
	Reason       string                      `json:"Reason"`
	EBusinessId  string                      `json:"EBusinessID"`
	LogisticCode string                      `json:"LogisticCode"` // 物流单号
	Shippers     []RecogniseResponseShippers `json:"Shippers"`
}

type RecogniseResponseShippers struct {
	ShipperCode string `json:"ShipperCode"` // 快递公司编码
	ShipperName string `json:"ShipperName"` // 快递公司名称
}

func (resp *RecogniseResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *RecogniseResponse) GetError() string {
	msg := ""
	if "" != resp.Code {
		msg = "[" + resp.Code + "]"
	}
	if "" != resp.Reason {
		msg = msg + resp.Reason
	}
	return msg
}
