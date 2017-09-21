package logickit

import (
	"fmt"
	"git.gumpcome.com/go_kit/strkit"
	"github.com/astaxie/beego"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 校验公司号
func VerificationCompanyId(companyId int) bool {
	if !(10000 <= companyId && companyId <= 99999) {
		beego.Error("公司号错误", fmt.Sprintf("公司号：%v", companyId))
		return false
	}
	return true
}

// 校验公司号（父id）
func VerificationCompanyPId(companyPId int) bool {
	if !(10000 <= companyPId && companyPId <= 99999 || companyPId == 0) {
		beego.Error("父公司号错误", fmt.Sprintf("公司号：%v", companyPId))
		return false
	}
	return true
}

// 校验【查询服务商运营商关系】查询方式
func VerificationOwnerQueryType(queryType int) bool {
	if !(queryType == 0 || queryType == 10 || queryType == 11) {
		beego.Error("查询服务商运营商关系:查询方式错误", fmt.Sprintf("query_type：%v", queryType))
		return false
	}
	return true
}

// 校验公司类型
func VerificationCompanyType(companyType int) bool {
	if !(companyType == 10 ||
		companyType == 11) {
		beego.Error("公司类型错误", fmt.Sprintf("公司类型=%v", companyType))
		return false
	}
	return true
}

// 校验支付方式
func VerificationPayWay(payWay int) bool {
	if !(-1 <= payWay && payWay <= 14) {
		beego.Error("支付方式错误", fmt.Sprintf("支付方式=%v", payWay))
		return false
	}
	return true
}

// 校验支付方式
func VerificationPayWayStrs(payWayStr string) (bool, []int) {
	if strkit.StrIsBlank(payWayStr) {
		beego.Error("支付方式错误", fmt.Sprintf("支付方式=%v", payWayStr))
		return false, nil
	}
	payWays := strings.Split(payWayStr, ",")
	payWayList := make([]int, 0, len(payWays))
	payWaysMap := make(map[int]int)
	for _, v := range payWays {
		if payWay, err := strconv.Atoi(v); err != nil {
			beego.Error(err)
			beego.Error("支付方式错误", fmt.Sprintf("支付方式=%v", payWayStr))
			return false, nil
		} else if VerificationPayWay(payWay) {
			payWaysMap[payWay] = payWay
			payWayList = append(payWayList, payWay)
		} else {
			beego.Error("支付方式错误", fmt.Sprintf("支付方式=%v", payWay))
			return false, nil
		}
	}
	if len(payWaysMap) != len(payWays) {
		beego.Error("存在重复的支付方式")
		return false, nil
	}
	return true, payWayList
}

// 校验售货机ID
func VerificationSvmId(svmId int) bool {
	if !(23456 <= svmId) {
		beego.Error("售货机ID错误", fmt.Sprintf("售货机ID：%v", svmId))
		return false
	}
	return true
}

// 校验售货机销售状态
func VerificationSvmSaleStatus(saleStatus int) bool {
	if !(saleStatus == 10 || saleStatus == 11) {
		beego.Error("校验售货机销售状态错误", fmt.Sprintf("售货机销售状态：%v", saleStatus))
		return false
	}
	return true
}

// 校验售货机运营状态
func VerificationSvmWorkStatus(workStatus int) bool {
	if !(workStatus == 10 || workStatus == 11 || workStatus == 12) {
		beego.Error("校验售货机运营状态错误", fmt.Sprintf("售货机运营状态：%v", workStatus))
		return false
	}
	return true
}

// 校验售货机经（纬）度
func VerificationSvmLonOrLat(lonOrLat string) bool {
	if utf8.RuneCountInString(lonOrLat) > 15 {
		beego.Error("售货机经度或维度长度错误", fmt.Sprintf("售货机经度或维度：%v", lonOrLat))
		return false
	} else if _, err := strconv.ParseFloat(lonOrLat, 10); err != nil {
		beego.Error("售货机经度或维度格式错误", fmt.Sprintf("售货机经度或维度：%v", lonOrLat))
		return false
	}
	return true
}

// 校验权限ID
func VerificationAccessId(accessId int) bool {
	if accessId < 100000 || accessId > 999999 {
		beego.Error("权限ID错误", fmt.Sprintf("权限ID：%v", accessId))
		return false
	}
	return true
}

// 校验权限父ID
func VerificationAccessPId(accessPId int) bool {
	if (accessPId < 100000 && accessPId != 0) || accessPId > 999999 {
		beego.Error("权限父ID错误", fmt.Sprintf("权限父ID：%v", accessPId))
		return false
	}
	return true
}

// 校验登录密码
func VerificationLoginPwd(pwd string) bool {
	if len(pwd) != utf8.RuneCountInString(pwd) || len(pwd) < 6 || len(pwd) > 20 {
		beego.Error("密码格式错误", fmt.Sprintf("密码：%v", pwd))
		return false
	}
	return true
}

// 校验登录账号
func VerificationLoginName(loginName string) bool {
	if len(loginName) != utf8.RuneCountInString(loginName) || len(loginName) < 6 || len(loginName) > 30 {
		beego.Error("账号格式错误", fmt.Sprintf("账号：%v", loginName))
		return false
	}
	return true
}

// 校验账号类型
func VerificationAccountType(accountType int) bool {
	if !(accountType == 10 || accountType == 11 || accountType == 12) {
		beego.Error("账号类型错误", fmt.Sprintf("账号类型：%v", accountType))
		return false
	}
	return true
}

// 校验手机号
func VerificationPhone(phone string) bool {
	if !(11 <= len(phone) && len(phone) <= 20) || len(phone) != utf8.RuneCountInString(phone) {
		beego.Error("手机号码格式错误", fmt.Sprintf("手机号码：%v", phone))
		return false
	}
	return true
}

func VerificationEmail(email string) bool {
	if len(email) != utf8.RuneCountInString(email) || len(email) < 5 || len(email) > 50 {
		beego.Error("邮箱格式错误", fmt.Sprintf("邮箱：%v", email))
		return false
	}
	emailMatch, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+.([a-zA-Z0-9_-])+$", email)
	if !emailMatch {
		beego.Error("邮箱格式错误", fmt.Sprintf("邮箱：%v", email))
		return false
	}
	return true
}

// 校验是否
func VerificationIsNo(isOrNo string) bool {
	if !(isOrNo == "Y" || isOrNo == "N") {
		beego.Error("Y/N格式错误", fmt.Sprintf("Y/N：%v", isOrNo))
		return false
	}
	return true
}

// 校验页码
func VerificationPageNumber(number int) bool {
	if !(0 <= number && number <= 1000000) {
		beego.Error("页码错误", fmt.Sprintf("页码：%v", number))
		return false
	}
	return true
}

// 校验页面展示数量
func VerificationPageSize(size int) bool {
	if !(0 <= size && size <= 100) {
		beego.Error("页面展示数目错误", fmt.Sprintf("页面展示数目：%v", size))
		return false
	}
	return true
}

// 校验普通折扣方案类型
func VerificationDiscountType(discountType int) bool {
	if !(discountType == Goods_Plan_Svm ||
		discountType == Goods_Plan_Present ||
		discountType == Goods_Plan_Single ||
		discountType == Goods_Plan_Many) {
		beego.Error("方案类型错误", fmt.Sprintf("类型：%v", discountType))
		return false
	}
	return true
}

// 校验方案状态(实物商品常规折扣营销方案，虚拟商品常规折扣营销方案)
func VerificationPlanStatus(status int) bool {
	if !(status == Plan_Status_Not_Activated ||
		status == Plan_Status_Activated ||
		status == Plan_Status_Has_Put ||
		status == Plan_Status_Stopped) {
		beego.Error("方案状态错误", fmt.Sprintf("方案状态：%v", status))
		return false
	}
	return true
}

// 校验虚拟商品方案类型
func VerificationVgoodsType(vgoodsType int) bool {
	// '扫码虚拟商品：11，套餐虚拟商品：12，惊喜虚拟商品：13'
	if !(vgoodsType == Vgoods_Plan_Qrcode ||
		vgoodsType == Vgoods_Plan_Package ||
		vgoodsType == Vgoods_Plan_Surprise) {
		beego.Error("虚拟商品方案类型错误", fmt.Sprintf("虚拟商品方案类型：%v", vgoodsType))
		return false
	}
	return true
}

// 校验虚拟商品方案类型
func VerificationAdType(adType int) bool {
	// 10：图片广告 11：视频广告
	if !(adType == 10 ||
		adType == 11) {
		beego.Error("广告类型错误", fmt.Sprintf("广告类型：%v", adType))
		return false
	}
	return true
}

func VerificationGoodsStatus(status int) bool {
	// 10：上架，11：下架
	if !(status == 10 ||
		status == 11) {
		beego.Error("商品状态错误", fmt.Sprintf("商品状态：%v", status))
		return false
	}
	return true
}

// 校验url
func VerificationUrl(url string) bool {
	if utf8.RuneCountInString(url) > 200 {
		beego.Error("url长度大于200", fmt.Sprintf("url长度：%v", utf8.RuneCountInString(url)))
		return false
	}
	return true
}

// 校验售货机拉去
func VerificationLatestConfigAction(action int) bool {
	if !(action == 10 || action == 11) {
		beego.Error("售货机拉去配置action状态值错误", fmt.Sprintf("action：%v", action))
		return false
	}
	return true
}
