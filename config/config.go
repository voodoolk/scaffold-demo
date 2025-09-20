package config

import (
	"scaffold-demo/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

var (
	Port       string
	JwtSignKey string
	JwtExpTime int
)

// 设置日志级别
func SetLogLevel(loglevel string) {
	if loglevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	//添加文件名和行号
	logrus.SetReportCaller(true)
	//文件格式改为json
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimeFormat})
}

func init() {
	logs.Debug(nil, "配置初始化...")
	//加载程序配置
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8080") //注意要加冒号，设置系统环境变量的时候也要加冒号
	viper.SetDefault("JWT_SIGN_KEY", "rabbit")
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	viper.AutomaticEnv()
	//获取程序配置
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpTime = viper.GetInt("JWT_EXPIRE_TIME")
	SetLogLevel(logLevel)
}
