// @Title 业务错误码
package logiccode

import (
	"strconv"
	"git.gumpcome.com/go_kit/strkit"
)

func New(code int, msg string) error {
	return &LogicCode{code, msg}
}

// @Title 业务错误码
// @Description
//	业务错误码由6位组成,前3位代表类别,后3位代表具体业务。
//	100XXX: 100代表通用类别;
// 		XXX<=100: 	代表DAO层错误;
// 		100<XXX<=200: 	代表logic层错误;
// 		200<XXX<=300: 	代表controller层错误;
type LogicCode struct {
	Code int `json:"code" desc:"业务错误码"`
	Msg string `json:"msg" desc:"错误描述"`
}

func (code *LogicCode) Error() string {
	return strkit.StrJoin("[", strconv.Itoa(code.Code), "]<", code.Msg, ">")
}

// @Title DB连接错误
// @Description 用于DAO层操作DB错误反馈
func DbConErrorCode() error {
	return New(100001, "db connect error")
}

// @Title DB插入操作错误
// @Description 用于DAO层操作DB错误反馈
func DbInsertErrorCode() error {
	return New(100002, "db insert error")
}

// @Title DB更新操作错误
// @Description 用于DAO层操作DB错误反馈
func DbUpdateErrorCode() error {
	return New(100003, "db update error")
}

// @Title 通过主键ID更新DB操作错误
// @Description 用于DAO层操作DB错误反馈
func DbUpdateByIdErrorCode() error {
	return New(100004, "db update by id is nil")
}

// @Title DB删除操作错误
// @Description 用于DAO层操作DB错误反馈
func DbDeleteErrorCode() error {
	return New(100005, "db delete error")
}

// @Title DB查询操作错误
// @Description 用于DAO层操作DB错误反馈
func DbQueryErrorCode() error {
	return New(100006, "db query error")
}

// @Title DB分页查询超出总页数范围
// @Description 用于DAO层操作DB错误反馈
//	页码错误、每页显示记录总数错误。
func DbPageOutErrorCode() error {
	return New(100007, "db page query out of range")
}

// @Title DB操作影响记录数为0
// @Description 用于DAO层插入、更新、删除记录时没有实际发生影响记录数
func DbZeroErrorCode() error {
	return New(100008, "db affected rows is 0")
}

// @Title 记录总数值字符串转整形异常
// @Description
func DbPageCountToIntCode() error {
	return New(100009, "page count string to int error")
}

// @Title 数据库配置名称错误
// @Description
func DbConfigNameErrorCode() error {
	return New(100010, "mysql config name is nil")
}

// @Title DB结果字符串转整型错误
// @Description 用于DAO层操作DB错误反馈
func DbItemToIntErrorCode() error {
	return New(100011, "db item to int error")
}

// @Title 请求参数值错误
// @Description 用于反射请求参数对象、参数值类型转换、必填参数校验错误反馈
func ReqParamErrorCode() error  {
	return New(100301, "param value error")
}