package test

import (
	"github.com/RiverDanceGit/kdniaoGo/data"
	"testing"
)

func TestExpressGetList(t *testing.T) {
	list1 := data.GetData()
	expressCompany := data.NewExpressCompany()
	list2 := expressCompany.GetList()
	if len(list1) != len(list2) {
		t.Error("total not equal", len(list1), len(list2))
		return
	}
	t.Log(len(list1))
}

func TestExpressSearch(t *testing.T) {
	words := []string{"中通", "百世", "圆通"}

	expressCompany := data.NewExpressCompany()
	for _, word := range words {
		list := expressCompany.Search(word)
		if len(list) == 0 {
			t.Error(word, "not found")
			return
		} else {
			t.Log(word, len(list))
		}
	}
}

func TestExpressGetExpressNameByCode(t *testing.T) {
	words := []string{"SF", "JD", "YTO"}

	expressCompany := data.NewExpressCompany()
	for _, word := range words {
		name := expressCompany.GetExpressNameByCode(word)
		if "" == name {
			t.Error(word, "not found")
			return
		} else {
			t.Log(word, name)
		}
	}
}

func TestExpressGetExpressCodeByName(t *testing.T) {
	words := []string{"顺丰速运", "京东快递", "圆通速递"}

	expressCompany := data.NewExpressCompany()
	for _, word := range words {
		code := expressCompany.GetExpressCodeByName(word)
		if "" == code {
			t.Error(word, "not found")
			return
		} else {
			t.Log(word, code)
		}
	}
}

func TestExpressGetInfoByExpressCode(t *testing.T) {
	words := []string{"SF", "JD", "YTO"}

	expressCompany := data.NewExpressCompany()
	for _, word := range words {
		info, has := expressCompany.GetInfoByExpressCode(word)
		if !has {
			t.Error(word, "not found")
			return
		} else {
			t.Log(word, info.ExpressCode, info.ExpressName)
		}
	}
}

func TestExpressGetInfoByExpressName(t *testing.T) {
	words := []string{"顺丰速运", "京东快递", "圆通速递"}

	expressCompany := data.NewExpressCompany()
	for _, word := range words {
		info, has := expressCompany.GetInfoByExpressName(word)
		if !has {
			t.Error(word, "not found")
			return
		} else {
			t.Log(word, info.ExpressCode, info.ExpressName)
		}
	}
}
