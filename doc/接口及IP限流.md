
# 接口及IP限流

采用 **令牌桶** 对可针对不同接口配置不同的限流策略，同时支持对用户IP限流，防止用户恶意攻击

通过`github.com/juju/ratelimit`实现令牌桶，将接口做map的key，对应不同的*BucketConf，BucketConf包含一个Bucket，用于对接口进行限流，接口限流支持阻塞等待，配置maxWait

每隔BucketConf配置一个map[string]*ratelimit.Bucket，用于对IP进行限流，IP限流不支持阻塞等待，达到阈值时直接阻止该请求。

```go
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
```

限流中间件初始化，对不同接口和对不同接口下的IP配置不同的限流策略

```go
// InitRateLimit 注册限流
func InitRateLimit() {
	// 需要限流的接口在此注册
	// 参考文章：https://blog.csdn.net/m0_52528053/article/details/127294249
	middleware.BucketMap["/api/v1/upload/token"] = &middleware.BucketConf{
		// 每10ms产生5个token，最多存储10个token
		Bucket: ratelimit.NewBucketWithQuantum(10*time.Millisecond, 10, 5),
		// 最大等待时间
		MaxWait: 20 * time.Millisecond,

		IpRateConf: &middleware.IpRateConf{
			FillInterval: time.Second,
			Capacity:     5,
		},
	}
	middleware.BucketMap["/api/v1/upload/callback"] = &middleware.BucketConf{
		Bucket:  ratelimit.NewBucketWithQuantum(10*time.Millisecond, 10, 5),
		MaxWait: 20 * time.Millisecond,
	}
	middleware.BucketMap["default"] = &middleware.BucketConf{
		Bucket:  ratelimit.NewBucketWithQuantum(10*time.Millisecond, 100, 20),
		MaxWait: 10 * time.Millisecond,
	}
	middleware.FillDefault()
}
```

将代码注册到路由中间件，当请求失败时报错请求频繁。

```go
// 代码注册到路由中间件
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
```
