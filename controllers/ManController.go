package controllers

import (
	"encoding/json"
)

// MainController struct
type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (that *MainController) NavMenu() {
	menuJson := []byte(`
	{
		"homeInfo": {
			"title": "首页",
			"href": "page/welcome"
		},
		"logoInfo": {
			"title": "hlinfoNcm",
			"image": "static/img/logo.jpg",
			"href": ""
		},
		"menuInfo": [
			{
				"title": "常规管理",
				"icon": "fa fa-address-book",
				"href": "",
				"target": "_self",
				"child": [
					{
						"title": "主页",
						"href": "page/welcome",
						"icon": "fa fa-home",
						"target": "_self"
					},
					{
						"title": "主机管理",
						"href": "page/hostlist",
						"icon": "fa fa-server",
						"target": "_self"
					},
					{
						"title": "负载均衡管理",
						"href": "page/slblist",
						"icon": "fa fa-sliders",
						"target": "_self"
					},
					{
						"title": "TCP代理",
						"href": "page/tcpProxy",
						"icon": "fa fa-tumblr",
						"target": "_self"
					},
					{
						"title": "服务管理",
						"href": "page/terminal",
						"icon": "fa fa-terminal",
						"target": "_self"
					},
					{
						"title": "系统设置",
						"href": "page/setting",
						"icon": "fa fa-gears",
						"target": "_self"
					}
				]
			},
			{
				"title": "DNS云解析",
				"icon": "fa fa-address-book",
				"href": "",
				"target": "_self",
				"child": [
					{
						"title": "概况",
						"href": "page/dnswelcome",
						"icon": "fa fa-home",
						"target": "_self"
					},
					{
						"title": "公网解析",
						"href": "",
						"icon": "fa fa-home",
						"target": "_self",
						"child": [
							{
								"title": "域名绑定",
								"href": "page/domainBind",
								"icon": "fa fa-terminal",
								"target": "_self"
							},
							{
								"title": "域名解析",
								"href": "page/publicDns",
								"icon": "fa fa-terminal",
								"target": "_self"
							}
						]
					},
					{
						"title": "内网解析",
						"href": "",
						"icon": "fa fa-home",
						"target": "_self",
						"child": [
							{
								"title": "内网域名管理",
								"href": "page/privateDomain",
								"icon": "fa fa-terminal",
								"target": "_self"
							},
							{
								"title": "内网域名解析",
								"href": "page/privateDns",
								"icon": "fa fa-terminal",
								"target": "_self"
							}
						]
					},
					{
						"id": "1024",
						"title": "DNS服务管理",
						"href": "page/dnsService",
						"icon": "fa fa-terminal",
						"target": "_self"
					}
				]
			}
		]
	}
	`)
	var rsJsonData map[string]interface{}
	_ = json.Unmarshal(menuJson, &rsJsonData)
	that.Data["json"] = rsJsonData
	that.ServeJSON()
}

func (that *MainController) ClearCache() {
	tmp := []byte(`{
		"code": 1,
		"msg": "服务端清理缓存成功"
	  }`)
	var rsJsonData map[string]interface{}
	_ = json.Unmarshal(tmp, &rsJsonData)
	that.Data["json"] = rsJsonData
	that.ServeJSON()
}
