package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// GetRandomString 生成随机7位字符串
func GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 7; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetMD5 生成32位MD5
func GetMD5() string {
	ctx := md5.New()
	ctx.Write([]byte(GetRandomString()))
	return hex.EncodeToString(ctx.Sum(nil))
}
