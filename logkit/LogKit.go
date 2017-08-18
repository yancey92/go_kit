package logkit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
)

const (
	LogDevMode  = "dev"
	LogProdMode = "prod"
)

func InitLog() {
	logmode := beego.AppConfig.String("logmode")
	if logmode == "" || (logmode != LogDevMode && logmode != LogProdMode) {
		panic("config logmode is empty or log mode isnot dev or prod!")
	}

	logDir := "logs"
	logFile := logDir + string(os.PathSeparator) + "server.log"
	os.MkdirAll(logDir, 0777)

	if logmode == LogDevMode {
		beego.SetLevel(beego.LevelInformational)
	} else {
		beego.SetLevel(beego.LevelInformational)
	}

	beego.SetLogger(logs.AdapterMultiFile, `{
		"filename":"`+logFile+`",
		"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]
	}`)
}
