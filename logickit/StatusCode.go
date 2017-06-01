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


)