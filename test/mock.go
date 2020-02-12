package test

import (
	"errors"
	"github.com/RiverDanceGit/kdniaoGo"
	"github.com/RiverDanceGit/kdniaoGo/util"
	"strings"
)

func getConfig() (kdniaoGo.KdniaoConfig, error) {
	eBusinessId, appKey, err := getConfigValue()
	return kdniaoGo.NewKdniaoConfig(eBusinessId, appKey), err
}

func getConfigValue() (string, string, error) {
	configStr, err := util.FileGetContents("../config.txt")
	if err != nil {
		return "", "", err
	}
	configs := strings.Split(configStr, ",")
	if 2 != len(configs) {
		return "", "", errors.New("eBusinessId or appKey is empty")
	}
	return configs[0], configs[1], nil
}
