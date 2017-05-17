package logickit

import (
	"regexp"
	"github.com/astaxie/beego"
	"git.gumpcome.com/go_kit/logiccode"
	"time"
	"git.gumpcome.com/go_kit/timekit"
	"strconv"
	"git.gumpcome.com/go_kit/strkit"
)

// 生成营销方案code码
func PlanCodeBuild(planType string, companyId int) (string, error) {
	planTypeMatch, _ := regexp.MatchString("^[A-Z]$", planType)
	if !planTypeMatch {
		beego.Error("生成方案CODE码时，校验方案类型错误")
		return "", logiccode.New(100301, "生成方案CODE码时,方案类型错误")
	}
	if !VerificationCompanyId(companyId) {
		beego.Error("生成方案CODE码时，校验公司号错误")
		return "", logiccode.New(100301, "生成方案CODE码时,公司号错误")
	}
	_, timeStr, err := timekit.GetTimeSsAndDate(time.Now(), timekit.DateFormat_YYYYMMDDHHMMSSMS)
	if err != nil {
		return "", logiccode.New(120017, "生成方案CODE码时,当前时间转换错误")
	}
	code := strkit.StringBuilder{}
	code.Append(planType).Append(strconv.Itoa(companyId)).Append(timeStr)
	return code.ToString(), nil
}
