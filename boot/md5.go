package boot

import (
	"crypto/md5"
	"encoding/hex"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 加密
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// 解密
func ValidPassword(plainpwd, salt, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}
