package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type WebVisitLog struct {
	Id            int64     `orm:"auto;pk"`
	VisitIp       string    `orm:"column(visit_ip)"`
	VisitDatetime time.Time `orm:"column(visit_datetime);null;type(text)"`
	VisitContent  string    `orm:"column(visit_content);null;type(text)"`
	VisitBrowser  string    `orm:"column(visit_browser);null;type(text)"`
	VisitIpinfo   string    `orm:"column(visit_ipinfo);null;type(text)"`
	VisitUsercode string    `orm:"column(visit_usercode);null;type(text)"`
	VisitOs       string    `orm:"column(visit_os);null;size(100)"`
}

type AdminList struct {
	Id            int       `orm:"auto;pk"`
	RealName      string    `orm:"column(real_name);type(text)"`
	Account       string    `orm:"column(account);type(text)"`
	AdminLoginPwd string    `orm:"column(admin_login_pwd);type(text)"`
	AdminLevel    int       `orm:"column(admin_level);default(0)"`
	LoginTimes    int       `orm:"column(login_times);default(0)"`
	Enabled       int       `orm:"column(enabled);default(0)"`
	LastLoginTime time.Time `orm:"column(last_login_time)"`
	Lastip        string    `orm:"column(lastip)"`
	ThatLoginTime time.Time `orm:"column(that_login_time)"`
	Thatip        string    `orm:"column(thatip)"`
	Createtime    time.Time `orm:"auto_now_add;type(datetime)"`
	Updatetime    time.Time `orm:"auto_now;type(datetime)"`
	Delete_time   time.Time `orm:"column(delete_time);null"`
}

type OperateLogs struct {
	Id        int64     `orm:"auto;pk"`
	Account   string    `orm:"column(account);type(text)"`
	Opname    string    `orm:"column(opname);type(text)"`
	Optype    int       `orm:"column(optype);default(0)"`
	OpAppType int       `orm:"column(op_app_type);default(2)"`
	Opdetail  string    `orm:"column(opdetail);type(text)"`
	Opip      string    `orm:"column(opip)"`
	Optime    time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(WebVisitLog), new(AdminList), new(OperateLogs))
}

func (that *AdminList) Fetch(id int) (wr *AdminList) {
	wr.Id = id
	o := orm.NewOrm()
	o.Read(&wr)
	return wr
}
