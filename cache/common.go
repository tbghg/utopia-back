package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

var Ctx = context.Background()
var expire = 12 * 24 * time.Hour

const (
	TypeString = iota
	TypeInt64
	TypeBool
)

// VideoLikeCountKey 视频点赞数
func VideoLikeCountKey(vid uint) string {
	return fmt.Sprintf("like:count:%d", vid)
}

// VideoLikeCountKeyV2 视频点赞数
func VideoLikeCountKeyV2(vid uint) string {
	return fmt.Sprintf("v2:like:%d", vid)
}

// GetStringCache 从redis里面获取字符串缓存，失败则根据dalFunc回源并写入DB
func GetStringCache(dalFunc func() (interface{}, error), key string, ResType int) (interface{}, error) {
	//先去redis里面查找
	resStr, err := RDB.Get(Ctx, key).Result()
	var res interface{}
	if err != nil {
		//redis里面没有，去数据库里面查找
		res, err = dalFunc()
		if err != nil {
			return nil, err
		}
		//将结果存入redis 异步
		go func() { RDB.Set(Ctx, key, res, expire) }()
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

	return res, nil
}
