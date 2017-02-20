package logkit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"strings"
)

const (
	LogDevMode  = "dev"
	LogProdMode = "prod"
)

var log *logs.BeeLogger

func InitLog() {
	appName := beego.AppConfig.String("appname")
	if appName == "" {
		panic("config appname is empty!")
	}
	logmode := beego.AppConfig.String("logmode")
	if logmode == "" || (logmode != LogDevMode && logmode != LogProdMode) {
		panic("config logmode is empty or log mode isnot dev or prod!")
	}
	logDir := strings.TrimSpace(appName) + "_logs"
	logFile := logDir + string(os.PathSeparator) + strings.TrimSpace(appName) + ".log"
	os.MkdirAll(logDir, 0777)

	log = logs.GetBeeLogger()
	log.Async(1e3) //1000
	if logmode == LogDevMode {
		log.SetLevel(beego.LevelDebug)
	} else {
		log.SetLevel(beego.LevelInformational)
	}

	log.SetLogger(logs.AdapterMultiFile, `{
		"filename":"`+logFile+`",
		"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]
	}`)
}

func GetLog() *logs.BeeLogger {
	return log
}
