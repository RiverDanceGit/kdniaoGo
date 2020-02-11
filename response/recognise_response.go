package response

type RecogniseResponse struct {
	Success      bool                        `json:"Success"`
	Code         string                      `json:"Code"`
	Reason       string                      `json:"Reason"`
	LogisticCode string                      `json:"LogisticCode"`
	Shippers     []RecogniseResponseShippers `json:"Shippers"`
}

type RecogniseResponseShippers struct {
	ShipperCode string `json:"ShipperCode"`
	ShipperName string `json:"ShipperName"`
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
