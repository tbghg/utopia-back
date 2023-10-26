package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"utopia-back/config"
	"utopia-back/pkg/logger"
)

type MyClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

// GenToken 生成 Token
func GenToken(id int) string {
	var mySecret = []byte(config.V.GetString("jwt.secret"))

	c := MyClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.V.GetString("jwt.issuer"),                                       // 签发人
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.V.GetDuration("jwt.expire"))), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                         // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                         // 生效时间
		}}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenStr, err := token.SignedString(mySecret)
	if err != nil {
		logger.Logger.Error("生成Token错误" + err.Error())
	}
	return tokenStr
}
