package logickit

import (
	"git.gumpcome.com/go_kit/logiccode"
	"git.gumpcome.com/go_kit/strkit"
)

// 根据订单号获取分表信息
// @outTradeNo 商户订单号
// 返回值
// 1:订单存储对应的表名
// 2:生成订单的月份 YYYYMM
// 3:生成订单的日期 YYYYMMDD
// 4:生成订单的时间 YYYYMMDDHHMMSS
func GetTabelInfoByOutTradeNo(outTradeNo string) (string, string, string, string, error) {
	// 订单时间
	orderTime := strkit.SubStr(outTradeNo, 0, 14)
	if orderTime == "" || len(orderTime) != 14 {
		return "", "", "", "", logiccode.New(120017, "依据订单号获取日期错误outTradeNo="+outTradeNo)
	}
	// 订单月份
	orderMonth := strkit.SubStr(outTradeNo, 0, 6)
	// 订单日期
	orderDay := strkit.SubStr(outTradeNo, 0, 8)
	// 订单表
	orderTable := "svm_trade_" + orderMonth

	return orderTable, orderMonth, orderDay, orderTime, nil
}

// 根据订单号获取分表信息
// @outTradeNo 商户订单号
// 返回值
// 1:订单存储对应的表名
// 2:生成订单的年份 YYYY
// 3:生成订单的月份 YYYYMM
// 4:生成订单的日期 YYYYMMDD
func GetTabelInfoWithYearByOutTradeNo(outTradeNo string) (string, string, string, string, error) {
	// 订单时间
	orderTime := strkit.SubStr(outTradeNo, 0, 14)
	if orderTime == "" || len(orderTime) != 14 {
		return "", "", "", "", logiccode.New(120017, "依据订单号获取日期错误outTradeNo="+outTradeNo)
	}
	// 订单年份
	orderYear := strkit.SubStr(outTradeNo, 0, 4)
	// 订单月份
	orderMonth := strkit.SubStr(outTradeNo, 0, 6)
	// 订单日期
	orderDay := strkit.SubStr(outTradeNo, 0, 8)
	// 订单表
	orderTable := "svm_trade_" + orderMonth

	return orderTable, orderYear, orderMonth, orderDay, nil
}


// 根据订单号获取分表信息
// @outTradeNo 商户订单号
// 返回值
// 1:中控存储订单对应的表名
// 2:生成订单的年份 YYYY
// 3:生成订单的月份 YYYYMM
// 4:生成订单的日期 YYYYMMDD
func GetPayTabelInfoWithYearByOutTradeNo(outTradeNo string) (string, string, string, string, error) {
	// 订单时间
	orderTime := strkit.SubStr(outTradeNo, 0, 14)
	if orderTime == "" || len(orderTime) != 14 {
		return "", "", "", "", logiccode.New(120017, "依据订单号获取日期错误outTradeNo="+outTradeNo)
	}
	// 订单年份
	orderYear := strkit.SubStr(outTradeNo, 0, 4)
	// 订单月份
	orderMonth := strkit.SubStr(outTradeNo, 0, 6)
	// 订单日期
	orderDay := strkit.SubStr(outTradeNo, 0, 8)
	// 订单表
	orderTable := "order_status_" + orderMonth

	return orderTable, orderYear, orderMonth, orderDay, nil
}

// 根据订单号获取分表信息
// @outTradeNo 商户订单号
// 返回值
// 1:中控订单退款信息存储对应的表名
// 2:生成订单的年份 YYYY
// 3:生成订单的月份 YYYYMM
// 4:生成订单的日期 YYYYMMDD
func GetRefundTabelInfoWithYearByOutTradeNo(outTradeNo string) (string, string, string, string, error) {
	// 订单时间
	orderTime := strkit.SubStr(outTradeNo, 0, 14)
	if orderTime == "" || len(orderTime) != 14 {
		return "", "", "", "", logiccode.New(120017, "依据订单号获取日期错误outTradeNo="+outTradeNo)
	}
	// 订单年份
	orderYear := strkit.SubStr(outTradeNo, 0, 4)
	// 订单月份
	orderMonth := strkit.SubStr(outTradeNo, 0, 6)
	// 订单日期
	orderDay := strkit.SubStr(outTradeNo, 0, 8)
	// 订单表
	orderTable := "order_refund_" + orderYear

	return orderTable, orderYear, orderMonth, orderDay, nil
}
