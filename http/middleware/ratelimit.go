package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var BucketMap = make(map[string]*BucketConf)

type BucketConf struct {
	Bucket  *ratelimit.Bucket
	MaxWait time.Duration
}

func RateLimit(c *gin.Context) {
	bucket, ok := BucketMap[c.Request.URL.Path]
	if !ok {
		bucket = BucketMap["default"]
	}
	if bucket.Bucket.WaitMaxDuration(1, bucket.MaxWait) {
		c.Next()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 501,
		"msg":  "请求频繁，请稍后重试",
	})
	c.Abort()
	return
}
