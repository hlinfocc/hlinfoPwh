package api

import (
	"bytes"
	"errors"
	"log"
	"time"

	"github.com/hlinfocc/hlinfoPoh/models"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	userAgent "github.com/wenlng/go-user-agent"
	"golang.org/x/crypto/bcrypt"
)

type AccountController struct {
	BaseController
}

type Resp struct {
	Code    int
	Success bool
	Msg     string
}

type LoginParam struct {
	UserName string
	Password string
}

func (that *AccountController) Get() {
	ua := that.Ctx.Input.UserAgent()
	o := orm.NewOrm()
	weblogs := new(models.WebVisitLog)
	weblogs.VisitIp = that.Ctx.Input.IP()
	weblogs.VisitDatetime = time.Now()
	weblogs.VisitBrowser = userAgent.GetBrowserName(ua)
	weblogs.VisitContent = "进入登录页面"
	weblogs.VisitIpinfo = ""
	weblogs.VisitOs = userAgent.GetOsName(ua)
	weblogs.VisitUsercode = ""

	id, err := o.Insert(weblogs)
	if err == nil {
		log.Println(id)
	}

	that.Data["Website"] = "beego.me"
	that.Data["Email"] = "astaxie@gmail.com"
	that.Data["CopyYear"] = time.Now().Year()
	that.TplName = "login.html"
}

func (that *AccountController) Post() {
	var res Resp
	res.Code = 200
	res.Success = false

	param := LoginParam{}
	if err := that.ParseForm(&param); err != nil {
		//handle error
		res.Msg = "解析参数失败"
		that.Data["json"] = res
		that.ServeJSON()
		return
	}
	valid := validation.Validation{}
	valid.Required(param.UserName, "UserName").Message("用户名不能为空")
	valid.Required(param.Password, "Password").Message("密码不能为空")
	valid.MinSize(param.Password, 8, "Password").Message("密码至少8位以上")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		var buffer bytes.Buffer
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			if buffer.Len() > 0 {
				buffer.WriteString("，")
			}
			buffer.WriteString(err.Message)
		}
		res.Msg = buffer.String()
		that.Data["json"] = res
		that.ServeJSON()
		return
	}

	log.Println(param)
	//读库验证

	hashPwd := "11111" //string(passwordbyte)
	isOk, _ := ValidatePassword(param.Password, hashPwd)
	if !isOk {
		res.Msg = "用户名或密码错误"
		that.Data["json"] = res
		that.ServeJSON()
		return
	}
	now := time.Now()
	var nowTime = now.Format("2006-01-02 15:04:05")
	log.Println(nowTime)
	res.Success = true
	res.Msg = "登录成功"
	that.Data["json"] = res
	that.ServeJSON()
}

func (that *AccountController) Exit() {

}

//GeneratePassword 给密码就行加密操作
func GeneratePassword(userPassword string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return string(byte), err
}

//ValidatePassword 密码比对
func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil

}
