package data

import (
	"strings"
)

func NewExpressCompany() ExpressCompany {
	expressCompany := ExpressCompany{}
	expressCompany.Init()
	return expressCompany
}

type ExpressCompany struct {
	List []ExpressCompanyItem
}

type ExpressCompanyItem struct {
	ExpressName string
	ExpressCode string
}

func (ec *ExpressCompany) Init() {
	ec.List = GetData()
}

// 获取所有快递公司列表(提供原始数据以便自行扩展查找的方法)
func (ec ExpressCompany) GetList() []ExpressCompanyItem {
	return ec.List
}

// 关键字模糊查找快递公司
func (ec ExpressCompany) Search(word string) []ExpressCompanyItem {
	items := []ExpressCompanyItem{}
	if "" == word {
		return items
	}
	for _, item := range ec.GetList() {
		if strings.Contains(item.ExpressName, word) {
			items = append(items, item)
		} else if strings.Contains(word, item.ExpressName) {
			items = append(items, item)
		} else if strings.Contains(item.ExpressCode, word) {
			items = append(items, item)
		}
	}
	return items
}

// 获取快递公司名称
func (ec ExpressCompany) GetExpressNameByCode(expressCode string) string {
	item, has := ec.GetInfoByExpressCode(expressCode)
	if has {
		return item.ExpressName
	}
	return ""
}

// 获取快递公司编号
func (ec ExpressCompany) GetExpressCodeByName(expressName string) string {
	item, has := ec.GetInfoByExpressName(expressName)
	if has {
		return item.ExpressCode
	}
	return ""
}

// 获取快递公司
func (ec ExpressCompany) GetInfoByExpressCode(expressCode string) (ExpressCompanyItem, bool) {
	var express ExpressCompanyItem
	if "" == expressCode {
		return express, false
	}
	for _, item := range ec.GetList() {
		if expressCode == item.ExpressCode {
			return item, true
		}
	}
	return express, false
}

// 获取快递公司
func (ec ExpressCompany) GetInfoByExpressName(expressName string) (ExpressCompanyItem, bool) {
	var express ExpressCompanyItem
	if "" == expressName {
		return express, false
	}
	for _, item := range ec.GetList() {
		if expressName == item.ExpressName {
			return item, true
		}
	}
	return express, false
}
