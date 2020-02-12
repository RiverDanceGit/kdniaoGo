package test

import (
	"github.com/RiverDanceGit/kdniaoGo"
	"github.com/RiverDanceGit/kdniaoGo/sdk"
	"testing"
)

func TestRecognise(t *testing.T) {
	config, err := getConfig()
	if err != nil {
		t.Error("err", err)
		return
	}
	logger := kdniaoGo.NewKdniaoLogger()

	apiRecogniseSdk := sdk.NewApiRecognise(config, logger)
	req := apiRecogniseSdk.GetRequest("550000609031770")
	resp, err := apiRecogniseSdk.GetResponse(req)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("resp.EBusinessId", resp.EBusinessId)
	t.Log("resp.LogisticCode", resp.LogisticCode)
	for _, shipper := range resp.Shippers {
		t.Log(shipper.ShipperCode, shipper.ShipperName)
	}
}
