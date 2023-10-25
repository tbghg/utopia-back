package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http/httputil"
	"os"
	"strings"
	"time"
	"utopia-back/pkg/logger"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//获取用户的请求信息
				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				//记录日志
				//判断是否是链接中断
				var brokenPipe bool
				if ne, ok := err.(*gin.Error); ok {
					var se *os.SyscallError
					if errors.As(ne.Err, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") {
							brokenPipe = true
						}
					}
				}
				// 如果是链接中断，不需要记录堆栈信息
				if brokenPipe {
					logger.Logger.Error(c.Request.URL.Path,
						zap.Time("time", time.Now()),
						zap.String("method", c.Request.Method),
						zap.String("ip", c.ClientIP()),
						zap.String("error", err.(string)),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				//如果不是链接中断，记录堆栈信息
				logger.Logger.Error("[Recovery from panic]",
					zap.Time("time", time.Now()),
					zap.String("method", c.Request.Method),
					zap.String("ip", c.ClientIP()),
					zap.String("error", err.(string)),
					zap.String("request", string(httpRequest)),
					zap.Stack("stacktrace"),
				)

			}
		}()
		c.Next()
	}
}
