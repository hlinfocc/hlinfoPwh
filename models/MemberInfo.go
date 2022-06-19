package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type MemberInfo struct {
	Id            int64     `orm:"auto;pk"`
	RealName      string    `orm:"column(real_name);null;type(text);" description:"真实姓名"`
	NickName      string    `orm:"column(nick_name);null;type(text)" description:"昵称"`
	Account       string    `orm:"column(account);unique;type(text)" description:"账号"`
	Password      string    `orm:"column(password);type(text)" description:"密码"`
	Description   string    `orm:"column(description);null;type(text)" description:"个人简介"`
	Email         string    `orm:"column(email);null;size(150)" description:"E-mail"`
	Tel           string    `orm:"column(tel);null;size(15)" description:"电话"`
	Avatar        string    `orm:"column(avatar);null;type(text)" description:"头像"`
	Role          int       `orm:"column(role);size(int);default(1)" description:"角色:0管理员,1普通用户"`
	RoleName      string    `orm:"column(role_name);null;size(255)" description:"角色名称"`
	Openid        string    `orm:"column(openid);null;type(text)" description:"openID"`
	Status        int       `orm:"column(status);size(int);default(1)" description:"状态:0未审核,1正常,2拒绝,3禁用"`
	JobName       string    `orm:"column(job_name);null;type(text)" description:"职位"`
	LastLoginTime time.Time `orm:"column(last_login_time)" description:"上一次登录时间"`
	Lastip        string    `orm:"column(lastip)" description:"上一次登录IP"`
	ThatLoginTime time.Time `orm:"column(that_login_time)" description:"本次登录时间"`
	Thatip        string    `orm:"column(thatip)" description:"本次登录IP"`
	Createtime    time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Updatetime    time.Time `orm:"auto_now;type(datetime)" description:"最后更新时间"`
	Delete_time   time.Time `orm:"column(delete_time);null" description:"删除标志"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(MemberInfo))
}
