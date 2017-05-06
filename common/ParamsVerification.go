package common

import ()
import (
	"fmt"
	"github.com/astaxie/beego"
	"unicode/utf8"
	"git.gumpcome.com/svm_mgr/constant"
)

// 校验公司号
func VerificationCompanyId(companyId int) bool {
	if !(10000 <= companyId && companyId <= 99999) {
		beego.Error("公司号错误", fmt.Sprintf("公司号：%v", companyId))
		return false
	}
	return true
}

func VerificationSVmId(svmId int) bool {
	if !(23456 <= svmId) {
		beego.Error("售货机ID错误", fmt.Sprintf("售货机ID：%v", svmId))
		return false
	}
	return true
}

func VerificationCompCategory(companyCategory int) bool {
	if !(companyCategory == constant.COMPANY_CATEGORY_PLATFORM ||
		companyCategory == constant.COMPANY_CATEGORY_SERVICE ||
		companyCategory == constant.COMPANY_CATEGORY_OPERATOR) {
		beego.Error("公司类型错误", fmt.Sprintf("公司类型=%v", companyCategory))
		return false
	}
	return true
}

func VerificationCompAccessCategory(companyCategory int) bool {
	if !(companyCategory == 0 ||
		companyCategory == constant.COMPANY_CATEGORY_PLATFORM ||
		companyCategory == constant.COMPANY_CATEGORY_SERVICE ||
		companyCategory == constant.COMPANY_CATEGORY_OPERATOR) {
		beego.Error("公司权限类型错误", fmt.Sprintf("公司权限类型=%v", companyCategory))
		return false
	}
	return true
}

func VerificationAccessId(accessId int) bool {
	if accessId < 100000 || accessId > 999999 {
		beego.Error("权限ID错误", fmt.Sprintf("权限ID：%v", accessId))
		return false
	}
	return true
}

func VerificationAccessPId(accessPId int) bool {
	if (accessPId < 100000 && accessPId != 0) || accessPId > 999999 {
		beego.Error("权限父ID错误", fmt.Sprintf("权限父ID：%v", accessPId))
		return false
	}
	return true
}

func VerificationLoginPwd(pwd string) bool {
	if len(pwd) != utf8.RuneCountInString(pwd) || len(pwd) < 8 {
		beego.Error("密码格式错误", fmt.Sprintf("密码：%v", pwd))
		return false
	}
	return true
}

func VerificationPhone(phone string) bool {
	if !(11 <= len(phone) && len(phone) <= 20) || len(phone) != utf8.RuneCountInString(phone) {
		beego.Error("手机号码格式错误", fmt.Sprintf("手机号码：%v", phone))
		return false
	}
	return true
}

func VerificationIsNo(isOrNo string) bool {
	if !(isOrNo == "Y" || isOrNo == "N") {
		beego.Error("Y/N格式错误", fmt.Sprintf("Y/N：%v", isOrNo))
		return false
	}
	return true
}

func VerificationPageNumber(number int) bool {
	if !(0 <= number && number <= 1000000) {
		beego.Error("页码错误", fmt.Sprintf("页码：%v", number))
		return false
	}
	return true
}

func VerificationPageSize(size int) bool {
	if !(0 <= size && size <= 50) {
		beego.Error("页面展示数目错误", fmt.Sprintf("页面展示数目：%v", size))
		return false
	}
	return true
}

func VerificationDiscountType(discountType int) bool {
	if !(discountType == 10 || discountType == 11 || discountType == 12 || discountType == 13) {
		beego.Error("方案类型错误", fmt.Sprintf("类型：%v", discountType))
		return false
	}
	return true
}

func VerificationPlanStatus(status int) bool {
	if !(status == 10 || status == 11 || status == 12 || status == 13) {
		beego.Error("方案状态错误", fmt.Sprintf("方案状态：%v", status))
		return false
	}
	return true
}

func VerificationVgoodsType(vgoodsType int) bool {
	// '扫码虚拟商品：11，套餐虚拟商品：12，惊喜虚拟商品：13'
	if !( vgoodsType == 11 || vgoodsType == 12 || vgoodsType == 13) {
		beego.Error("虚拟商品方案类型错误", fmt.Sprintf("虚拟商品方案类型：%v", vgoodsType))
		return false
	}
	return true
}