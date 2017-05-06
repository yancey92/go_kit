package logickit

import (
	"github.com/astaxie/beego/httplib"
	"strconv"
	"github.com/astaxie/beego"
	"git.gumpcome.com/go_kit/logiccode"
	"net/http"
	"git.gumpcome.com/goods_mgr/httpprotocol"
	"fmt"
)

// 账户售货机鉴权
func AccountSvmAuth(url string, accountId, svmId int) (bool, error) {
	req := httplib.Post(url)
	req.Param("account_id", strconv.Itoa(accountId))
	req.Param("svm_id", strconv.Itoa(svmId))
	respAuth := httpprotocol.RespAccountSvmAuth{}
	if err := req.ToJSON(&respAuth); err != nil {
		beego.Error("账户下某设备的鉴权失败", err)
		beego.Error(fmt.Sprintf("account_id=%v,svm_id=%v", accountId, svmId))
		return false, logiccode.New(100400, "账户下某设备的鉴权失败")
	}
	if respAuth.Code != http.StatusOK {
		beego.Error(respAuth)
		return false, logiccode.New(respAuth.Code, respAuth.Msg)
	}
	if respAuth.Data.IsSuccess != "Y" {
		return false, logiccode.New(100302, "此账户没有管理该售货机的权限")
	}
	return true, nil
}