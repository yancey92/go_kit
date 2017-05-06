package common

import (
	"strings"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego"
	"net/http"
	"git.gumpcome.com/goods_mgr/constant"
	"git.gumpcome.com/goods_mgr/httpprotocol"
)

// 批量更新售货机版本号
// return:是否成功，更新版本号失败的售货机id
func ConfigVersionBatchUpdate(svmIds []int) (bool, []int) {
	// 批量修改售货机版本号可放置售货机的最大数量
	batchEditSvmIdNum := 50
	forNum := len(svmIds) / batchEditSvmIdNum
	if len(svmIds) % batchEditSvmIdNum > 0 {
		forNum = forNum + 1
	}
	// 批量更新结果
	isOk := true
	var failureSvmIds []int // 版本号更新失败的售货机id
	// 批量升级svm版本号
	for i := 1; i <= forNum; i++ {
		var svmidsBatch []int
		if i == forNum && len(svmIds) % batchEditSvmIdNum > 0 {
			svmidsBatch = svmIds[(i - 1) * batchEditSvmIdNum:]
		} else {
			svmidsBatch = svmIds[(i - 1) * batchEditSvmIdNum : batchEditSvmIdNum]
		}
		svmIdStr := strings.Replace(strings.Replace(strings.Replace(fmt.Sprint(svmidsBatch), " ", ",", batchEditSvmIdNum), "[", "", 1), "]", "", 1)
		req := httplib.Post(constant.BatchUpdateConfigVersion)
		req.Param("svm_ids", svmIdStr)
		resp := httpprotocol.RespBatchUpdateSvmConfigVersion{}
		if err := req.ToJSON(&resp); err != nil {
			// 结构体转义失败
			beego.Error(err, svmidsBatch)
			isOk = false
			failureSvmIds = append(failureSvmIds, svmidsBatch...)
		}
		if resp.Code != http.StatusOK {
			isOk = false
			beego.Error(resp)
			failureSvmIds = append(failureSvmIds, resp.Data.FailSvms...)
		}
	}
	return isOk, failureSvmIds
}


