package setting

import (
	"os"
	"time"

	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	RedisHost string
	RedisPort int
)

func init() {
	LoadLog()
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Fail to parse")
	}
	Cfg.BlockMode = false //配置文件加锁字段，纯读可以关闭
	LoadBase()
	LoadServer()
	LoadApp()
	LoadRedis()
}

func LoadLog() {
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true) // 显示行号等信息
	// log.SetFormatter(&log.JSONFormatter{
	// 	PrettyPrint:     true,                  //格式化json
	// 	TimestampFormat: "2006-01-02 15:04:05", //时间格式化
	// })
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间格式化
	})
	log.SetOutput(os.Stdout)

}

func LoadBase() {
	RunMode = Cfg.Section("env").Key("RUN_MODE").MustString("debug")

	log.WithFields(log.Fields{
		"error": nil,
	}).Info("Success to get env")

	log.WithFields(log.Fields{
		"error": nil,
	}).Warn("Success to get env")

	log.WithFields(log.Fields{
		"error": nil,
	}).Debug("Success to get env")

	log.WithFields(log.Fields{
		"error": nil,
	}).Error("Success to get env")

	//print and exit
	// log.WithFields(log.Fields{
	// 	"error": nil,
	// }).Fatal("Success to get env")

	//print and panic
	// log.WithFields(log.Fields{
	// 	"error": nil,
	// }).Panic("Success to get env")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Fail to get section")
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("page")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Panic("Fail to get section")
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Fail to get section")
	}
	RedisHost = sec.Key("HOST").MustString("127.0.0.1")
	RedisPort = sec.Key("PORT").MustInt(6379)
}
