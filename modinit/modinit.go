package modinit

import (
	"os"
	"flag"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/dbv/wen/utils/conf"
	"github.com/dbv/wen/utils/log"
)

type SetJson struct {
	SerName  string `json:"server_name"`
	UseMysql string `json:"use_mysql"`
	MysqlUrl string `json:"mysql_url"`
	LogPath  string `json:"log_path"`
	FilePath string `json:"file_path"`
}

//所有数据初始化
const version = "0.1"

var XormInstance *xorm.Engine = nil

var (
	currPath, _ = os.Getwd()
	configPath  = flag.String("c", currPath+"/conf/set.ini", "Path to a configuration file")
	logPath     = flag.String("o", currPath+"/log/", "log path")
	CfgData     = SetJson{}
)

func init() {
	ConfgInit()
	if CfgData.UseMysql > "0" {
		MysqlInit()
	}
}

func ConfgInit() {
	//配置加载
	flag.Parse()
	//log初始化
	log.NewLogger(*logPath, log.LoggerLevelDebug)
	log.Debug("currPath:", currPath, " configPath:", *configPath)
	if len(*configPath) <= 1 {
		log.Error("配置未找到")
	}
	cfg, err := conf.LoadConfigFile(*configPath)
	if err != nil {
		log.Error("读取配置文件失败[conf.ini]")
		return
	}
	CfgData.SerName, _ = cfg.GetValue("app", "server_name")
	CfgData.UseMysql, _ = cfg.GetValue("app", "use_mysql")
	CfgData.MysqlUrl, _ = cfg.GetValue("app", "mysql_url")
	CfgData.LogPath, _ = cfg.GetValue("app", "log_path")
	CfgData.FilePath, _ = cfg.GetValue("app", "file_path")
	log.Debug("配置加载成功:", CfgData)
}

//mysql初始化
func MysqlInit() {
	//建立mysql连接
	if XormInstance == nil {
		var err error
		XormInstance, err = xorm.NewEngine("mysql", CfgData.MysqlUrl)
		if err != nil {
			log.Error("初始化mysql失败")
		} else {
			log.Debug("初始化mysql连接成功")
		}
	}
}

//获取资源
func GetMysqlHandler() *xorm.Engine {
	for XormInstance == nil {
		MysqlInit()
	}
	return XormInstance
}
