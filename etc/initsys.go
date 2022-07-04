package etc

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 初始化数据库连接
func RegisterDataBase() {
	//获取数据库类型，用于今后兼容多个数据库用
	dbtype := beego.AppConfig.DefaultString("datasource::type", "sqlite")
	host := beego.AppConfig.DefaultString("datasource::host", "127.0.0.1")
	port := beego.AppConfig.DefaultInt("datasource::port", 5432)
	username := beego.AppConfig.DefaultString("datasource::username", "postgres")
	dbname := beego.AppConfig.String("datasource::dbname")
	password := beego.AppConfig.String("datasource::password")
	runmode := beego.AppConfig.String("runmode")

	if dbtype == "" || dbtype == "pgsql" || dbtype == "postgresql" {
		orm.RegisterDriver("postgres", orm.DRPostgres)
		dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
		orm.RegisterDataBase("default", "postgres", dataSource)
		if runmode == "dev" || runmode == "test" {
			orm.Debug = true
		}
	} else if dbtype == "sqlite" {
		orm.DefaultTimeLoc = time.UTC
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", "./data/pwh.db")
		if runmode == "dev" || runmode == "test" {
			orm.Debug = true
		}
	} else {
		//mysql等数据库以后在扩展
		errors.New("暂不支持PostgreSQL及sqlite以外的数据库")
	}
}

// 获取本机网卡IP
func getLocalIp() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return "0.0.0.0"
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return "0.0.0.0"
}

// 初始化配置
func RegisterInitConfig() {
	// 应用名称
	beego.BConfig.AppName = "hlinfoPwh"
	//Flash 数据设置时 Cookie 的名称，
	beego.BConfig.WebConfig.FlashName = "HLINFOPOH_FLASH"
	// Flash 数据的分隔符
	beego.BConfig.WebConfig.FlashSeparator = "HLINFOPOHFLASH"
	// 模板左标签，默认值是{{
	beego.BConfig.WebConfig.TemplateLeft = "@{"
	// 模板右标签，默认值是}}
	beego.BConfig.WebConfig.TemplateRight = "}@"
	// XSRF 的 key 信息，默认值是 beegoxsrf。 EnableXSRF＝true 才有效
	beego.BConfig.WebConfig.XSRFKey = "ncmxsrf"

	//session 是否开启，默认是 false。
	beego.BConfig.WebConfig.Session.SessionOn = true
	// 存在客户端的 cookie 名称，默认值是 beegosessionID。
	beego.BConfig.WebConfig.Session.SessionName = "ncmSessionID"
	// 是否允许在 HTTP 请求时，返回原始请求体数据字节，默认为 false （GET or HEAD or 上传文件请求除外）。
	beego.BConfig.CopyRequestBody = true
	// 是否开启 gzip 支持
	beego.BConfig.EnableGzip = true
	//beego 服务器默认在请求的时候输出 server 名称。
	beego.BConfig.ServerName = "hlinfoWS"
	//设置默认端口
	sysHttpPort := beego.AppConfig.DefaultInt("httpport", 1088)
	beego.BConfig.Listen.HTTPPort = sysHttpPort
	//设置监听IP
	sysHttpAddr := beego.AppConfig.DefaultString("httpaddr", getLocalIp())
	beego.BConfig.Listen.HTTPAddr = sysHttpAddr
}

// 初始化数据库
func InitDatabases() {
	orm.RunSyncdb("default", false, true)
}

func InitSys() {
	RegisterInitConfig()
	RegisterDataBase()
	InitDatabases()
}
