package logickit

// 平台业务状态码
const (
	IsOk = 200

	// 数据库操作
	DbConnectError      = 100001 // 数据库连接失败
	DbInsertError       = 100002 // 插入数据库失败
	DbUpdateError       = 100003 // 更新数据库失败
	DbUpdateByIdIsNil   = 100004 // 更新数据库使用的ID是nil
	DbDeleteError       = 100005 // 数据库删除数据失败
	DbQueryError        = 100006 // 数据库查询数据失败
	DbPageQueryOutRange = 100007 // 分页查询页面展示数据数目不对
	DbNoAffectedRows    = 100008 // 数据库受影响的行数是0

	// 公司，账户
	CompanyTypeError         = 120001 //公司类型不合法
	CompanyDisable           = 120013 // 公司不可用。
	CompanyAlreadyDue        = 120014 // 公司使用系统期限已到期。
	AccountNotBelongsCompany = 120018 // 账户与公司号无管理关系。
	AccountNotBelongsGroup   = 120019 // 账户与群组无管理关系。
	NotIsDispatcher          = 120011 // 非总调度员角色。
	AccountDisable           = 120012 // 账户不可用。
	AccountAuthFail          = 120020 //账户鉴权失败

	// 设备
	SvmAuthError                       = 100302 // 售货机认证错误
	ConfigVersionUpdateError           = 120007 // 数据已更新，版本更新失败。
	DeviceCodeNotRelationMainBoardCode = 120021 // 设备编码和工控编码不一致
	NotConfigGoods                     = 120022 // 设备上该商品未配置
	WorkStatusError                    = 120026 // 工作状态错误
	SaleStatusError                    = 120027 // 售卖状态错误
	DeviceDisable                      = 120028 // 设备不可用
	DeviceNoPointLocation              = 120030 // 设备没有点位

	// 方案
	PlanStatusError         = 120003 // 方案状态不合法。
	PlanNoUsedSvm           = 120004 // 方案未配置售货机。
	PlanExpired             = 120005 // 方案已过期。
	PlanUsedSvm             = 120006 // 方案中该售货机已存在。
	MinuteTimeError         = 120008 // 分时段折扣信息错误。
	PlanGooodsExceededLimit = 120009 // 营销方案商品数量超出限制。
	PlanGoodsNotOnly        = 120010 // 营销方案售卖商品不唯一错误。

	//取货码营销方案
	CodeIsUsed    = 130011 //取货码已被使用
	CodeIsInvalid = 130012 //取货码无效
	SvmPlanError  = 130013 //获取方案配置售货机数据失败
	CodeIsLocked  = 130014 //取货码已被锁
	SvmNoGoods    = 130015 //设备无货
	// 其他
	PageCountStringToIntError = 100009 //
	ParamValueError           = 100301 // 参数错误
	HttpRequestError          = 100400 // http请求错误
	ExceededLimit             = 120000 // 超出总量限制
	DateScopeError            = 120002 // 日期范围不合法
	ContentContainSpecialChar = 120015 // 内容含有特殊字符。
	RecordNoExist             = 120016 // 记录不存在。
	FormatConvertError        = 120017 // 格式转换错误。
	PriceIsZero               = 120023 // 价格为0
	PaysCodeError             = 120024 // 支付二维码获取失败
	BufferOperationError      = 120025 // 缓存操作失败
	KeyNotExistRedis          = 100018 // redis中获取key不存在

	// 目录文件
	FileReadError  = 130001 // 文件读错误
	FileWriteError = 130002 // 文件写错误

)

const (
	// 虚拟商品营销方案类型标识 (同时对应虚拟商品类型)
	Vgoods_Plan_Qrcode   = 11 // 扫码
	Vgoods_Plan_Package  = 12 // 套餐
	Vgoods_Plan_Surprise = 13 // 惊喜

	// 常规折扣商品营销方案类型标识
	Goods_Plan_Svm     = 10 // 整机折扣
	Goods_Plan_Present = 11 // 买赠活动
	Goods_Plan_Single  = 12 // 单件购活动
	Goods_Plan_Many    = 13 // 两件购活动

	// 营销方案状态码
	Plan_Status_Not_Activated = 10 // 未激活
	Plan_Status_Activated     = 11 // 已激活
	Plan_Status_Has_Put       = 12 // 已投放
	Plan_Status_Stopped       = 13 // 已停止

	// 生成营销方案code码时，使用的方案类型code标识
	Code_Vgoods_Plan_Qrcode   = "H" // 虚拟商品扫码方案code类型
	Code_Vgoods_Plan_Package  = "I" // 虚拟商品套餐方案code类型
	Code_Vgoods_Plan_Surprise = "J" // 虚拟商品惊喜方案code类型

	Code_Goods_Plan_Svm     = "A" // 实物商品整机折扣code类型
	Code_Goods_Plan_Present = "B" // 实物商品买赠活动code类型
	Code_Goods_Plan_Single  = "C" // 实物商品单件购code类型
	Code_Goods_Plan_Many    = "D" // 实物商品多件购code类型

	Code_Goods_Plan_Home = "E" // 首页商品营销方案
	Code_Plan_Ads        = "F" // 广告营销方案
	Code_Plan_PayTag     = "G" // 支付标签营销方案

	Exchange_Plan_code = "K" //取货码方案code

)

const (
	// 工作状态 10：线上运营，11：停机保养，12：待机调试',
	SVM_WORKSTATUS_RUN   = 10
	SVM_WORKSTATUS_STOP  = 11
	SVM_WORKSTATUS_DEBUG = 12

	// 销售状态，10：正常销售，11：禁止销售
	SVM_SALESTATUS_ENABLE  = 10
	SVM_SALESTATUS_DISABLE = 11
)
