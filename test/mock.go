package test

import (
	"errors"
	"github.com/RiverDanceGit/kdniaoGo"
	"os"
)

func getConfig() (kdniaoGo.KdniaoConfig, error) {
	eBusinessId := os.Getenv("APP_EBUSINESS_ID")
	appKey := os.Getenv("APP_KEY")
	if "" == eBusinessId || "" == appKey {
		return kdniaoGo.NewKdniaoConfig(eBusinessId, appKey), errors.New(".env not exists")
	}
	return kdniaoGo.NewKdniaoConfig(eBusinessId, appKey), nil
}
