package test

import (
	"github.com/RiverDanceGit/kdniaoGo"
	"github.com/RiverDanceGit/kdniaoGo/sdk"
	"testing"
)

func TestMonitorRealtime(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error("err", err)
		return
	}
	logger := kdniaoGo.NewKdniaoLogger()

	apiMonitorRealtimeSdk := sdk.NewApiMonitorRealtime(config, logger)
	req := apiMonitorRealtimeSdk.GetRequest("4303618027892", "YD")
	resp, err := apiMonitorRealtimeSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessId)
	t.Log("resp.OrderCode", resp.OrderCode)
	t.Log("resp.ShipperCode", resp.ShipperCode)
	t.Log("resp.LogisticCode", resp.LogisticCode)
	t.Log("resp.State", resp.State)
	t.Log("resp.StateEx", resp.StateEx)
	t.Log("resp.Location", resp.Location)
	for _, trace := range resp.Traces {
		t.Log(trace.AcceptTime, "|", trace.AcceptStation, "|", trace.Location, "|", trace.Action, "|", trace.Remark)
	}
}
