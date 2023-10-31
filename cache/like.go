package cache

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"utopia-back/pkg/logger"
)

const (
	sMinVid     = "minVid"
	sTTL        = "ttl"
	maxFieldNum = 1500 // hash中最大字段数
	setFieldNum = 750  // hash重设字段数
)

// VideoLikeCountKey 视频点赞数
func VideoLikeCountKey(vid uint) string {
	return fmt.Sprintf("like:count:%d", vid)
}

// VideoLikeCountKeyV2 视频点赞数
func VideoLikeCountKeyV2(vid uint) string {
	return fmt.Sprintf("v2:like:%d", vid)
}

// UserLikedVideoKeyV3 用户点赞情况
//
// 用户维度缓存点赞情况
func UserLikedVideoKeyV3(uid uint) string {
	return fmt.Sprintf("v3:like:%d", uid)
}

// IsUserLikedVideo 用户是否为某个视频点赞(单个)
//
// state 1 查询成功；2 不存在该key,需要新建；3 冷数据,需回源
func IsUserLikedVideo(uid uint, vid uint) (liked bool, state int) {
	key := UserLikedVideoKeyV3(uid)

	// 查询是否存在
	res := RDB.HMGet(Ctx, key, sMinVid, sTTL, strconv.Itoa(int(vid)))
	if res.Err() != nil {
		logger.Logger.Error(fmt.Sprintf("IsUserLikedVideo cmd:%v", res.String()))
		return
	} else {
		logger.Logger.Info(fmt.Sprintf("IsUserLikedVideo cmd:%v", res.String()))
	}

	minVid, ok1 := res.Val()[0].(uint)
	ttl, ok2 := res.Val()[1].(time.Time)

	if !ok1 || !ok2 { // 不存在该key，需重新构建
		state = 2
		return
	}

	// 剩余时间小于三分之一，重新续期
	if time.Now().Sub(ttl) < time.Hour*24*15/3 {
		res := RDB.Expire(Ctx, key, time.Hour*24*15)
		if err := res.Err(); err != nil {
			logger.Logger.Error(fmt.Sprintf("BuildUserLikedVideos RDB.Expire cmd:%v", res.String()))
		} else {
			logger.Logger.Info(fmt.Sprintf("BuildUserLikedVideos RDB.Expire cmd:%v", res.String()))
		}
	}

	_, ok := res.Val()[2].(uint)
	// 如果不是冷数据，此时就是最终结果
	if vid >= minVid {
		state = 1
		liked = ok
		return
	}
	// 冷数据，查到则返回true，未查到设置state
	if ok {
		liked = true
	} else {
		state = 3
	}
	return
}

// IsUserLikedVideos 用户是否为某些视频点赞(批量)
//
// result key为vid  <->  0未点赞；1为点赞；2 冷数据,需回源
// state 1 成功；2 不存在该key,需要新建
func IsUserLikedVideos(uid uint, videoIds []uint) (result map[uint]int, state int) {
	key := UserLikedVideoKeyV3(uid)

	var builder strings.Builder
	for i := range videoIds {
		builder.WriteString(strconv.Itoa(int(videoIds[i])))
		builder.WriteString(" ")
	}
	field := builder.String()

	// 查询是否存在
	res := RDB.HMGet(Ctx, key, sMinVid, sTTL, field)
	if res.Err() != nil {
		logger.Logger.Error(fmt.Sprintf("IsUserLikedVideo cmd:%v", res.String()))
		return
	} else {
		logger.Logger.Info(fmt.Sprintf("IsUserLikedVideo cmd:%v", res.String()))
	}

	minVid, ok1 := res.Val()[0].(uint)
	ttl, ok2 := res.Val()[1].(time.Time)

	if !ok1 || !ok2 { // 不存在该key，需重新构建
		state = 2
		return
	}

	state = 1

	// 判断数量是否超过域值，超过则重建
	lenRes := RDB.HLen(Ctx, key)
	if res.Err() != nil {
		logger.Logger.Error(fmt.Sprintf("IsUserLikedVideo RDB.HLen cmd:%v", lenRes.String()))
		return
	} else {
		logger.Logger.Info(fmt.Sprintf("IsUserLikedVideo RDB.HLen cmd:%v", lenRes.String()))
	}

	if lenRes.Val() > maxFieldNum { // 数量超过域值
		go rebuildUserLikedVideos(key)

	} else if time.Now().Sub(ttl) < time.Hour*24*15/3 { // 剩余时间小于三分之一，重新续期

		res := RDB.Expire(Ctx, key, time.Hour*24*15)
		if err := res.Err(); err != nil {
			logger.Logger.Error(fmt.Sprintf("BuildUserLikedVideos RDB.Expire cmd:%v", res.String()))
		} else {
			logger.Logger.Info(fmt.Sprintf("BuildUserLikedVideos RDB.Expire cmd:%v", res.String()))
		}
	}

	vidRes := res.Val()[2:]
	result = make(map[uint]int, len(videoIds))
	for i, vid := range videoIds {
		if _, ok := vidRes[i].(uint); ok { // 查询到点赞
			result[vid] = 1
		} else if vid >= minVid { // 热数据，用户没点赞
			result[vid] = 0
		} else { // 冷数据，需要回源
			result[vid] = 2
		}
	}
	return
}

// BuildUserLikedVideos 构建用户点赞视频缓存
func BuildUserLikedVideos(key string, vid []uint, minVid uint) {
	var err error

	ttl := time.Now().AddDate(0, 0, 15) // 一个月的过期时间

	var builder strings.Builder
	for i := range vid {
		builder.WriteString(strconv.Itoa(int(vid[i])))
		builder.WriteString(" 1 ")
	}
	field := builder.String()

	res := RDB.HMSet(Ctx, key, sTTL, ttl, sMinVid, minVid, field)
	if err = res.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf("BuildUserLikedVideos RDB.HMSet cmd:%v", res.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("BuildUserLikedVideos RDB.HMSet cmd:%v", res.String()))
	}

	res = RDB.Expire(Ctx, key, time.Hour*24*15)
	if err = res.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf("BuildUserLikedVideos RDB.Expire cmd:%v", res.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("BuildUserLikedVideos RDB.Expire cmd:%v", res.String()))
	}
}

func rebuildUserLikedVideos(key string) {
	res := RDB.HGetAll(Ctx, key)
	if err := res.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf("rebuildUserLikedVideos RDB.HGetAll cmd:%v", res.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("rebuildUserLikedVideos RDB.HGetAll cmd:%v", res.String()))
	}

	videoIds := make([]uint, 0, len(res.Val()))
	for k := range res.Val() {
		num, _ := strconv.Atoi(k)
		videoIds = append(videoIds, uint(num))
	}
	sort.Slice(videoIds, func(i, j int) bool {
		return videoIds[i] > videoIds[j]
	})

	BuildUserLikedVideos(key, videoIds[:setFieldNum], videoIds[setFieldNum-1])
}
