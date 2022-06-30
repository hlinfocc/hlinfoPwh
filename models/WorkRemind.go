package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	uuid "github.com/satori/go.uuid"
)

/*
*工作提醒表
 */
type WorkRemind struct {
	Id          string    `orm:"size(36);pk"`
	Memeid      int64     `orm:"column(memeid)" description:"用户id"`
	Title       string    `orm:"column(title);type(text)" description:"标题"`
	Status      int       `orm:"column(status);size(int);default(0)" description:"状态:0正常,1不再提醒"`
	WrType      int       `orm:"column(wr_type);size(int);default(0)" description:"资源类型:0日常事件,1服务器,2域名"`
	HostInfo    string    `orm:"column(host_info);null;type(text)" description:"主机IP或域名"`
	Detail      string    `orm:"column(detail);null;type(text)" description:"详情"`
	Remarks     string    `orm:"column(remarks);null;type(text)" description:"备注"`
	StartTime   time.Time `orm:"column(start_time);type(datetime)" description:"开始时间"`
	ExpiresTime time.Time `orm:"column(expires_time);type(datetime)" description:"到期时间"`
	Createtime  time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Updatetime  time.Time `orm:"auto_now;type(datetime)" description:"最后更新时间"`
	DeleteTime  time.Time `orm:"column(delete_time);null" description:"删除标志"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(WorkRemind))
}

// 查询主机列表，首字母大写（public）
func QueryWrList(limit int, page int, memid int64, status int, keywords string) ([]*WorkRemind, int64, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}
	offset := (page - 1) * limit

	o := orm.NewOrm()
	cnd := orm.NewCondition()
	host := new(WorkRemind)
	qs := o.QueryTable(host) // 返回 QuerySeter
	cnd2 := cnd.And("DeleteTime__isnull", false)
	if status >= 0 && status <= 2 {
		cnd2.And("Status", status)
	}
	cnd2.And("memid", memid)
	if keywords != "" {
		cnd1 := cnd.Or("Title__icontains", keywords).Or("HostInfo__icontains", keywords).Or("Detail__icontains", keywords)
		cnd1.Or("Remarks__icontains", keywords)
		cnd2 = cnd2.AndCond(cnd1)
	}
	qs = qs.SetCond(cnd2)
	qs = qs.OrderBy("-createtime")
	qs = qs.Offset(offset).Limit(limit)
	qs = qs.Limit(limit, offset)
	var lists []*WorkRemind
	_, err := qs.All(&lists)
	if err != nil {
		return lists, 0, err
	}
	total, errCount := qs.Count()
	if errCount != nil {
		return lists, 0, errCount
	}
	return lists, total, nil
}

func FetchWr(id string) (wr *WorkRemind) {
	wr.Id = id
	o := orm.NewOrm()
	o.Read(&wr)
	return wr
}

func WorkRemindSave(wr WorkRemind) (err error) {
	sysid := uuid.NewV4()
	wr.Id = sysid.String()
	_, err = orm.NewOrm().Insert(&wr)
	return
}
