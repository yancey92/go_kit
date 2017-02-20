// @Title 唯一ID生成工具
package idkit

import (
	"crypto/md5"
	"io"
	"encoding/hex"
)

// @Title 创建32位唯一ID
// @Description
// usage:
//	CreateUniqueId()
func CreateUniqueId() string {
	return NewV1().String()
}

// @Title 创建字符串MD5
// @Description
// usage:
//	CreateMd5("abc")
func CreateMd5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return hex.EncodeToString(m.Sum(nil))
}
