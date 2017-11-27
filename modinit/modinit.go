package modinit

import (
	"os"
	"flag"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/dbv/wen/utils/conf"
	"github.com/dbv/wen/utils/log"
	"gopkg.in/redis.v5"
)

type SetJson struct {
	SerName  string `json:"server_name"`
	UseMysql string `json:"use_mysql"`
	MysqlUrl string `json:"mysql_url"`
	LogPath  string `json:"log_path"`
	FilePath string `json:"file_path"`
	UseRedis string `json:"use_redis"`
	RedisUrl string `json:"redis_url"`
}

//所有数据初始化
const version = "0.1"

var XormInstance *xorm.Engine = nil
var RedisInstance *redis.Client = nil

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
	if RedisInstance == nil {
		RedisInit()
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
	CfgData.UseRedis, _ = cfg.GetValue("app", "use_redis")
	CfgData.RedisUrl, _ = cfg.GetValue("app", "redis_url")
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

//redis初始化
func RedisInit() {
	if RedisInstance == nil {
		//单实例
		RedisInstance = redis.NewClient(&redis.Options{
			Addr:     CfgData.RedisUrl,
			Password: "",
			DB:       0,
		})
		//集群
		//RedisInstance = nil
		//var redisClus *redis.ClusterClient = nil
		//redisClus = redis.NewClusterClient(&redis.ClusterOptions{
		//	Addrs:              redisaddrs,
		//	Password:           "",
		//	MaxRedirects:       16,
		//	ReadOnly:           true,
		//	RouteByLatency:     true,
		//	DialTimeout:        10000,
		//	ReadTimeout:        30000,
		//	WriteTimeout:       30000,
		//	PoolSize:           10,
		//	PoolTimeout:        35000,
		//	IdleTimeout:        600,
		//	IdleCheckFrequency: 60,
		//})
		//_, err = redisClus.Ping().Result()
		//if err != nil {
		//	panic("redis error - " + err.Error())
		//} else {
		//	log.Debug("redis load...")
		//}
		//RedisInstance = redisClus
	}
}

//获取资源
func GetMysqlHandler() *xorm.Engine {
	for XormInstance == nil {
		MysqlInit()
	}
	return XormInstance
}
