package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/hlinfocc/hlinfoPwh/models"
)

// Operations about WorkRemind
type WorkRemindController struct {
	BaseController
}

// @Title CreateUser
// @Description 增加
// @Param	body		body 	models.WorkRemind	true		"body for user content"
// @Success 200 {int} models.WorkRemind
// @Failure 403 body is empty
// @router / [post]
func (u *WorkRemindController) Post() {
	var wr models.WorkRemind
	json.Unmarshal(u.Ctx.Input.RequestBody, &wr)
	fmt.Println(wr.StartTime)
	fmt.Println(wr.ExpiresTime)
	err := models.WorkRemindSave(wr)
	fmt.Println(err)
	if err == nil {
		u.Success("操作成功", 1, &wr)
	} else {
		u.Error("保存失败", err.Error())
	}
}

// @Title List
// @Description 列表查询
// @Success 200 {object} models.WorkRemind
// @router /list [get]
func (that *WorkRemindController) List() {
	page, _ := that.GetInt("page", 1)
	limit, _ := that.GetInt("limit", 20)
	status, _ := that.GetInt("status", -1)
	keywords := that.GetString("keywords", "")
	var memid int64
	list, total, _ := models.QueryWrList(limit, page, memid, status, keywords)
	that.Layuiok("获取成功", total, &list)
}

// @Title Get
// @Description get WorkRemind by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.WorkRemind
// @Failure 403 :id is empty
// @router /:id [get]
func (u *WorkRemindController) Get() {
	id := u.GetString(":id")
	if id != "" {
		workRemind := models.FetchWr(id)
		u.Data["json"] = workRemind
		u.Success("获取成功", 1, &workRemind)
	}
	u.Failed("获取数据失败")
}

// @Title Update
// @Description update the WorkRemind
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.WorkRemind	true		"body for WorkRemind content"
// @Success 200 {object} models.WorkRemind
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *WorkRemindController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the WorkRemind
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *WorkRemindController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
