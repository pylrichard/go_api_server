package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt 使用bcrypt进行加密
func Encrypt(src string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)

	return string(hashedBytes), err
}

// Compare 比较密码和加密密码
func Compare(hashedPwd, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}