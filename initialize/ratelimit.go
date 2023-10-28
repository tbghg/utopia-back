package initialize

import (
	"github.com/juju/ratelimit"
	"time"
	"utopia-back/http/middleware"
)

// InitRateLimit 注册限流
func InitRateLimit() {
	// 需要限流的接口在此注册
	// 参考文章：https://blog.csdn.net/m0_52528053/article/details/127294249
	middleware.BucketMap["/api/v1/upload/token"] = &middleware.BucketConf{
		Bucket:  ratelimit.NewBucketWithQuantum(10*time.Millisecond, 10, 5),
		MaxWait: 20 * time.Millisecond,
	}
	middleware.BucketMap["/api/v1/upload/callback"] = &middleware.BucketConf{
		Bucket:  ratelimit.NewBucketWithQuantum(10*time.Millisecond, 10, 5),
		MaxWait: 20 * time.Millisecond,
	}

	middleware.BucketMap["default"] = &middleware.BucketConf{
		Bucket:  ratelimit.NewBucketWithQuantum(10*time.Millisecond, 100, 20),
		MaxWait: 10 * time.Millisecond,
	}
}
