package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
)

// Md5Encode 生成md5
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	res := hex.EncodeToString(h.Sum(nil))
	return res
}

// RandomSalt 生成随机盐
func RandomSalt() string {
	return fmt.Sprint(rand.Int31n(1000000))
}

// Md5EncodeWithSalt 生成带盐的md5
func Md5EncodeWithSalt(data string, salt string) string {
	return Md5Encode(data + salt)
}

// ValidMd5EncodeWithSalt  验证带盐的md5
func ValidMd5EncodeWithSalt(data string, salt string, md5 string) bool {
	return Md5Encode(data+salt) == md5
}

// Md5EncodeUpper 生成大写的md5
func Md5EncodeUpper(data string) string {
	return Md5Encode(data)
}
