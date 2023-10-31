package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var BucketMap = make(map[string]*BucketConf)
var defaultIpRateConf = &IpRateConf{
	FillInterval: 100 * time.Millisecond,
	Capacity:     5,
}

type BucketConf struct {
	Bucket  *ratelimit.Bucket
	MaxWait time.Duration

	IpRateConf   *IpRateConf
	IpRateBucket map[string]*ratelimit.Bucket
}

type IpRateConf struct {
	FillInterval time.Duration
	Capacity     int64
}

func RateLimit(c *gin.Context) {
	bucketConf, ok := BucketMap[c.Request.URL.Path]
	if !ok {
		bucketConf = BucketMap["default"]
	}

	if bucketConf.Bucket.WaitMaxDuration(1, bucketConf.MaxWait) {
		ipRateValidate(c.ClientIP(), bucketConf.IpRateConf, bucketConf.IpRateBucket)
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

func FillDefault() {
	for _, bConf := range BucketMap {
		if bConf.IpRateConf == nil {
			bConf.IpRateConf = defaultIpRateConf
		}
	}
}

func ipRateValidate(clientIp string, ipRateConf *IpRateConf, IpRateBucket map[string]*ratelimit.Bucket) {
	b, ok := IpRateBucket[clientIp]
	if ok {
		b.Wait(1)
	} else {
		IpRateBucket = make(map[string]*ratelimit.Bucket)
		IpRateBucket[clientIp] = ratelimit.NewBucket(ipRateConf.FillInterval, ipRateConf.Capacity)
	}
}
