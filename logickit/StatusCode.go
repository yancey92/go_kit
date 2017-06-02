package logickit

const (
	IsOk = 200
	DbConnectError = 100001 // 数据库连接失败
	DbInsertError = 100002 // 插入数据库失败
	DbUpdateError = 100003 // 更新数据库失败
	DbUpdateByIdIsNil = 100004 // 更新数据库使用的ID是nil
	DbDeleteError = 100005 // 数据库删除数据失败
	DbQueryError = 100006 // 数据库查询数据失败
	DbPageQueryOutRange = 100007 // 分页查询页面展示数据数目不对
	DbNoAffectedRows = 100008 // 数据库受影响的行数是0

	PageCountStringToIntError = 100009 //
	ParamValueError = 100301 // 参数错误
	SvmAuthError = 100302 // 售货机认证错误
	HttpRequestError = 100400 // http请求错误

	ExceededLimit = 120000 // 超出总量限制
	CompanyTypeError = 120001 //公司类型不合法
	DateScopeError = 120002 // 日期范围不合法
	PlanStatusError = 120003 // 方案状态不合法。
	PlanNoUsedSvm = 120004 // 方案未配置售货机。
	PlanExpired = 120005 // 方案已过期。
	PlanUsedSvm = 120006 // 方案中该售货机已存在。
	ConfigVersionUpdateError = 120007 // 数据已更新，版本更新失败。
	MinuteTimeError = 120008 // 分时段折扣信息错误。
	PlanGooodsExceededLimit = 120009 // 营销方案商品数量超出限制。
	PlanGoodsNotOnly = 120010 // 营销方案售卖商品不唯一错误。
	NotIsDispatcher = 120011 // 非总调度员角色。
	AccountDisable = 120012 // 账户不可用。
	CompanyDisable = 120013 // 公司不可用。
	CompanyAlreadyDue = 120014 // 公司使用系统期限已到期。
	ContentContainSpecialChar = 120015 // 内容含有特殊字符。
	RecordNoExist = 120016 // 记录不存在。
	FormatConvertError = 120017 // 格式转换错误。
	AccountNotBelongsCompany = 120018 // 账户与公司号无管理关系。
	AccountNotBelongsGroup = 120019 // 账户与群组无管理关系。
	DeviceNotExist = 120020 // 设备不存在
	DeviceCodeNotRelationMainBoardCode = 120021 // 设备编码和工控编码不一致
	NotConfigGoods = 120022 // 设备上该商品未配置
	PriceIsZero = 120023 // 价格为0
	PaysCodeError = 120024 // 支付二维码获取失败
	BufferOperationError = 120025 // 缓存操作失败
	WorkStatusError = 120026 // 工作状态错误
	SaleStatusError = 120027 // 售卖状态错误
	DeviceDisable = 120028 // 设备不可用

	SerCompanyCannotOperOtherSvm = 130000 // 服务商不能授权添加其他服务商的售货机。

)