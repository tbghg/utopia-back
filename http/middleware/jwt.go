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
		//如果是feed请求，直接放行
		if c.Request.URL.Path == "/api/v1/feed" {
			c.Next()
			return
		}
		// 判断错误类型并处理
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 { //token格式错误
				c.JSON(http.StatusOK, gin.H{
					"code": 401,
					"msg":  ErrorTokenFormatMsg,
				})
				c.Abort() //阻止执行
				return
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 { //token过期
				c.JSON(http.StatusOK, gin.H{
					"code": 401,
					"msg":  ErrorTokenExpiredMsg,
				})
				c.Abort() //阻止执行
				return
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { //token未激活
				c.JSON(http.StatusOK, gin.H{
					"code": 401,
					"msg":  ErrorTokenNotValidYetMsg,
				})
				c.Abort() //阻止执行
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 401,
					"msg":  "无法处理此Token",
				})
				c.Abort() //阻止执行
				return
			}
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
