package intkit

// @Title 判断多个字符串是否为空
// @Description
// @param strs
// usage:
//	IntIsZero(0) return false
func IntIsZero(ints ...int64) bool {
	if len(ints) == 0 {
		return false
	}
	for _, v := range ints {
		if v != 0 {
			return false
		}
	}
	return true
}
