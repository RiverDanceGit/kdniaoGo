package test

import (
	"github.com/RiverDanceGit/kdniaoGo"
	"github.com/RiverDanceGit/kdniaoGo/sdk"
	"testing"
)

func TestMonitorSubscribe(t *testing.T) {
	config := getConfig()
	logger := kdniaoGo.NewKdniaoLogger()

	apiMonitorSubscribeSdk := sdk.NewApiMonitorSubscribe(config, logger)
	req := apiMonitorSubscribeSdk.GetRequest("L288474879009752702", "ZTO")
	resp, err := apiMonitorSubscribeSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessId)
	t.Log("resp.UpdateTime", resp.UpdateTime)
}
