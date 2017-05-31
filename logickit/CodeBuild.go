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

const (
	// 一下各个方案状态在其他项目中也有用到
	// 各个方案对应类型标识
	Vgoods_Plan_Qrcode = 11 // 扫码
	Vgoods_Plan_Package = 12  // 套餐
	Vgoods_Plan_Surprise = 13 // 惊喜

	Goods_Plan_Svm = 10 // 整机折扣
	Goods_Plan_Present = 11 // 买赠活动
	Goods_Plan_Single = 12 // 单件购活动
	Goods_Plan_Many = 13 // 两件购活动


	// 生成营销方案code码时，使用的方案类型code标识
	Code_Vgoods_Plan_Qrcode = "H" // 虚拟商品扫码方案code类型
	Code_Vgoods_Plan_Package = "I" // 虚拟商品套餐方案code类型
	Code_Vgoods_Plan_Surprise = "J" // 虚拟商品惊喜方案code类型

	Code_Goods_Plan_Svm = "A" // 实物商品整机折扣code类型
	Code_Goods_Plan_Present = "B" // 实物商品买赠活动code类型
	Code_Goods_Plan_Single = "C" // 实物商品单件购code类型
	Code_Goods_Plan_Many = "D" // 实物商品多件购code类型

	Code_Goods_Plan_Home = "E" // 首页商品营销方案
	Code_Plan_Ads = "F" // 广告营销方案
	Code_Plan_PayTag = "G" // 广告营销方案

)

// 常规实物营销方案code获取
func GetCommonDiscountPlanCode(planType int, companyId int) (string, error) {
	var code string
	var err error
	switch planType {
	case Goods_Plan_Svm:
		code, err = planCodeBuild(Code_Goods_Plan_Svm, companyId)
	case Goods_Plan_Present:
		code, err = planCodeBuild(Code_Goods_Plan_Present, companyId)
	case Goods_Plan_Single:
		code, err = planCodeBuild(Code_Goods_Plan_Single, companyId)
	case Goods_Plan_Many:
		code, err = planCodeBuild(Code_Goods_Plan_Many, companyId)
	default:
		return "", logiccode.New(100301, "方案类型错误")
	}
	if err != nil {
		beego.Error(err)
	}
	return code, err
}

// 复杂营销方案（套餐，扫码虚拟商品，惊喜虚拟商品）
func GetComplexPlanCode(planType int, companyId int) (string, error) {
	var code string
	var err error
	switch planType {
	case Vgoods_Plan_Qrcode:
		code, err = planCodeBuild(Code_Vgoods_Plan_Qrcode, companyId)
	case Vgoods_Plan_Package:
		code, err = planCodeBuild(Code_Vgoods_Plan_Package, companyId)
	case Vgoods_Plan_Surprise:
		code, err = planCodeBuild(Code_Vgoods_Plan_Surprise, companyId)
	default:
		return "", logiccode.New(100301, "方案类型错误")
	}
	if err != nil {
		beego.Error(err)
	}
	return code, err
}

// 生成营销方案code码
func planCodeBuild(planType string, companyId int) (string, error) {
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
