package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"utopia-back/config"
	utils "utopia-back/pkg/util"
)

const (
	ErrorTokenFormatMsg      = "token格式错误"
	ErrorTokenExpiredMsg     = "token过期"
	ErrorTokenNotValidYetMsg = "token未激活"
	ErrorTokenInvalidMsg     = "token无效"
	ErrorTokenHandleFailed   = "无法处理此Token"
)

func JwtMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	//请求头中没有Authorization
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请求头中auth为空",
		})
		c.Abort() //阻止执行
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  ErrorTokenFormatMsg,
		})
		c.Abort() //阻止执行
		return
	}
	//解析token
	tokenStr := parts[1]
	secret := []byte(config.V.GetString("jwt.secret"))
	token, err := jwt.ParseWithClaims(tokenStr, &utils.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		// 判断错误类型并处理
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			var msg string
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				msg = ErrorTokenFormatMsg //token格式错误
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				msg = ErrorTokenExpiredMsg //token过期
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				msg = ErrorTokenNotValidYetMsg //token未激活
			} else {
				msg = ErrorTokenHandleFailed //无法处理此Token
			}

			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  msg,
			})
			c.Abort() //阻止执行
			return
		}
	}
	if claims, ok := token.Claims.(*utils.MyClaims); ok && token.Valid {
		id := claims.ID
		c.Set("user_id", id)
		c.Next()
		return
	}
	//失效的token
	c.JSON(http.StatusOK, gin.H{
		"code": 401,
		"msg":  ErrorTokenInvalidMsg,
	})
	c.Abort() //阻止执行
	return

}

func JwtWithoutAbortMiddleware(c *gin.Context) {
	defer func() {
		c.Next()
	}()

	authHeader := c.Request.Header.Get("Authorization")
	//请求头中没有Authorization
	if authHeader == "" {
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return
	}
	//解析token
	tokenStr := parts[1]
	secret := []byte(config.V.GetString("jwt.secret"))
	token, err := jwt.ParseWithClaims(tokenStr, &utils.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(*utils.MyClaims); ok && token.Valid {
		id := claims.ID
		c.Set("user_id", id)
		return
	}
	return
}
