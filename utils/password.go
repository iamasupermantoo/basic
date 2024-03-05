package utils

import (
	"crypto/md5"
	"fmt"
)

// PasswordEncrypt 密码加密
func PasswordEncrypt(password string) string {
	if password == "" {
		return ""
	}

	has1 := md5.Sum([]byte(password))
	has2 := md5.Sum([]byte(fmt.Sprintf("%x", has1)))
	return fmt.Sprintf("%x", has2)
}
