package redis

import (
	"context"
	"strconv"
	"time"
)

var ctx = context.Background()
var expire = 12 * 24 * time.Hour

const (
	TypeString = iota
	TypeInt64
	TypeBool
)

// Cache 缓存
func Cache(dalFunc func() (interface{}, error), key string, ResType int) (interface{}, error) {
	//先去redis里面查找
	resStr, err := RDB.Get(ctx, key).Result()
	var res interface{}
	if err != nil {
		//redis里面没有，去数据库里面查找
		res, err = dalFunc()
		if err != nil {
			return nil, err
		}
		//将结果存入redis
		RDB.Set(ctx, key, res, expire)
		return res, nil
	}
	//redis里面有,把字符串转换成对应的类型,然后返回
	switch ResType {
	case TypeString:
		return resStr, nil
	case TypeInt64:
		res, err = strconv.ParseInt(resStr, 10, 64)
		if err != nil {
			return 0, err
		}
	case TypeBool:
		res, err = strconv.ParseBool(resStr)
		if err != nil {
			return false, err
		}
	}

	return resStr, nil
}
