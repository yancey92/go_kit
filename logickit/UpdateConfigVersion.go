package logickit

import (
	"strings"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego"
	"net/http"
	"strconv"
)

// 批量更新售货机版本号
type RespBatchUpdateSvmConfigVersion struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		     IsAllSuccess string   `json:"is_all_success" desc:"Y:全部成功，N:全部失败"`
		     FailSvms     []int    `json:"fail_svms" desc:"版本号更新失败的售货机id集合"`
	     } `json:"data"`
}
// 更新单个售货机版本号
type RespUpdateSvmConfigVersion struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		     IsSuccess string   `json:"is_success" desc:"Y：更新成功，N：更新失败"`
		     SvmId     int    `json:"svm_id" desc:"版本号更新失败的售货机id"`
	     } `json:"data"`
}


// 批量更新售货机版本号（只要有一个售货机版本号更新失败，则返回false）
// return:是否成功，更新版本号失败的售货机id
func ConfigVersionBatchUpdate(url string, svmIds []int) (bool, []int) {
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
		resp := RespBatchUpdateSvmConfigVersion{}
		beego.Info("批量修改售货机版本号URl：",url)
		if err := httplib.Post(url).
			Param("svm_ids", svmIdStr).
			ToJSON(&resp);
			err != nil {
			// 结构体转义失败
			beego.Error(err, svmidsBatch)
			isOk = false
			failureSvmIds = append(failureSvmIds, svmidsBatch...)
		}
		if resp.Code != http.StatusOK {
			isOk = false
			beego.Error(resp)
			if len(resp.Data.FailSvms) > 0 {
				failureSvmIds = append(failureSvmIds, resp.Data.FailSvms...)
			} else {
				failureSvmIds = append(failureSvmIds, svmidsBatch...)
			}
		} else {
			failureSvmIds = append(failureSvmIds, resp.Data.FailSvms...)
		}
	}
	return isOk, failureSvmIds
}

//单个更新售货机版本号
func ConfigVersionUpdate(url string, svmId int) (bool, int) {
	httpClient := httplib.Post(url).
		Param("svm_id", strconv.Itoa(svmId))
	resp := RespUpdateSvmConfigVersion{}
	if err := httpClient.ToJSON(&resp); err != nil {
		// 结构体转义失败
		beego.Error(fmt.Sprintf("更新单个售货机版本号返回体解析错误svmid=%#v", svmId), err)
		return false, svmId
	}
	if resp.Code != http.StatusOK || resp.Data.IsSuccess != "Y" {
		beego.Error("更新单个售货机版本号失败", fmt.Sprintf("%#v", resp))
		return false, svmId
	}
	return true, svmId
}