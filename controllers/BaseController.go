package controllers

import (
	"bytes"

	"github.com/astaxie/beego"
)

// BaseController struct
type BaseController struct {
	beego.Controller
}

// ResJson 响应 json 结果
func (that *BaseController) ResJson(code int, msg string, count int64, data ...interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg
	jsonData["count"] = count

	if len(data) > 0 {
		jsonData["data"] = data
	}

	that.Data["json"] = &jsonData
	that.ServeJSON()

	// returnJSON, err := json.Marshal(jsonData)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// that.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	// //禁用缓存
	// //that.Ctx.ResponseWriter.Header().Set("Cache-Control", "no-cache, no-store")
	// //使用gzip原始，json数据会只有原本数据的10分之一左右
	// if strings.Contains(strings.ToLower(that.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
	// 	that.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
	// 	//gzip压缩
	// 	w := gzip.NewWriter(that.Ctx.ResponseWriter)
	// 	defer w.Close()
	// 	w.Write(returnJSON)
	// 	w.Flush()
	// } else {
	// 	io.WriteString(that.Ctx.ResponseWriter, string(returnJSON))
	// }
	// that.StopRun()

}

// layuiok 返回Layui 数据表格的数据格式
func (that *BaseController) Layuiok(msg string, count int64, data interface{}) {
	jsonData := make(map[string]interface{}, 3)
	jsonData["code"] = 0
	jsonData["msg"] = msg
	jsonData["count"] = count

	if data != nil {
		jsonData["data"] = data
	}
	that.Data["json"] = &jsonData
	that.ServeJSON()
}

// 成功
func (that *BaseController) Success(msg string, count int64, data interface{}) {
	that.ResJson(200, msg, count, data)
}

func (that *BaseController) Error(msg string, data interface{}) {
	that.ResJson(403, msg, 0, data)
}

func (that *BaseController) Failed(msg string) {
	that.ResJson(403, msg, 0, nil)
}

// ExecuteViewPathTemplate 执行指定的模板并返回执行结果.
func (that *BaseController) ExecuteViewPathTemplate(tplName string, data interface{}) (string, error) {
	var buf bytes.Buffer

	viewPath := that.ViewPath

	if that.ViewPath == "" {
		viewPath = beego.BConfig.WebConfig.ViewsPath

	}

	if err := beego.ExecuteViewPathTemplate(&buf, tplName, viewPath, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
