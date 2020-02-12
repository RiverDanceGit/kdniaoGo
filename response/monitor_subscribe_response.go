package response

type MonitorSubscribeResponse struct {
	Success     bool   `json:"Success"`
	Reason      string `json:"Reason"`
	EBusinessId string `json:"EBusinessID"`
	UpdateTime  string `json:"UpdateTime"`
}

func (resp *MonitorSubscribeResponse) IsSuccess() bool {
	return resp.Success
}

func (resp *MonitorSubscribeResponse) GetError() string {
	return resp.Reason
}
