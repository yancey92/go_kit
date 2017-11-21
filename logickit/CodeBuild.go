package logickit

import (
	"github.com/astaxie/beego"
	"gitlab.gumpcome.com/common/go_kit/logiccode"
	"gitlab.gumpcome.com/common/go_kit/strkit"
	"gitlab.gumpcome.com/common/go_kit/timekit"
	"regexp"
	"strconv"
	"time"
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
		beego.Error("方案类型错误", planType)
		return "", logiccode.New(ParamValueError, "方案类型错误")
	}
	if err != nil {
		beego.Error(err)
	}
	return code, err
}

//取货码活动code码
func GetExchangPlanCode(companyId int) (string, error) {
	code, err := planCodeBuild(Exchange_Plan_code, companyId)
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
		beego.Error("方案类型错误", planType)
		return "", logiccode.New(ParamValueError, "方案类型错误")
	}
	if err != nil {
		beego.Error(err)
	}
	return code, err
}

// 广告方案code码生成（视频和图片广告）
func GetAdPlanCode(companyId int) (string, error) {
	var code string
	var err error
	code, err = planCodeBuild(Code_Plan_Ads, companyId)
	if err != nil {
		beego.Error(err)
	}
	return code, err
}

// 首页商品营销方案code码生成
func GetHomeGoodsPlanCode(companyId int) (string, error) {
	var code string
	var err error
	code, err = planCodeBuild(Code_Goods_Plan_Home, companyId)
	if err != nil {
		beego.Error(err)
	}
	return code, err
}

// 支付标签方案code码生成
func GetPayTagPlanCode(companyId int) (string, error) {
	var code string
	var err error
	code, err = planCodeBuild(Code_Plan_PayTag, companyId)
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
		return "", logiccode.New(ParamValueError, "生成方案CODE码时,方案类型错误")
	}
	if !VerificationCompanyId(companyId) {
		beego.Error("生成方案CODE码时，校验公司号错误")
		return "", logiccode.New(ParamValueError, "生成方案CODE码时,公司号错误")
	}
	_, timeStr, err := timekit.GetTimeSsAndDate(time.Now(), timekit.DateFormat_YYYYMMDDHHMMSSMS)
	if err != nil {
		beego.Error(err)
		return "", logiccode.New(FormatConvertError, "生成方案CODE码时,当前时间转换错误")
	}
	code := strkit.StringBuilder{}
	code.Append(planType).Append(strconv.Itoa(companyId)).Append(timeStr)
	return code.ToString(), nil
}
