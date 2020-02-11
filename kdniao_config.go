package kdniaoGo

func NewKdniaoConfig(eBusinessId, appKey string) KdniaoConfig {
	kdniaoConfig := KdniaoConfig{}
	kdniaoConfig.SetEBusinessId(eBusinessId)
	kdniaoConfig.SetAppKey(appKey)
	return kdniaoConfig
}

type KdniaoConfig struct {
	eBusinessId string
	appKey      string
}

func (req *KdniaoConfig) SetEBusinessId(eBusinessId string) *KdniaoConfig {
	req.eBusinessId = eBusinessId
	return req
}

func (req KdniaoConfig) GetEBusinessId() string {
	return req.eBusinessId
}

func (req *KdniaoConfig) SetAppKey(appKey string) *KdniaoConfig {
	req.appKey = appKey
	return req
}

func (req KdniaoConfig) GetAppKey() string {
	return req.appKey
}
