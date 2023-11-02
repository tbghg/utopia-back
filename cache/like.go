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
	sMinVid                        = "minVid"
	sTTL                           = "ttl"
	MaxFieldNum                    = 1500 // hash中最大字段数
	SetFieldNum                    = 750  // hash重设字段数
	userLikedVideoExpireTime       = time.Hour * 24 * 15
	userLikedVideoUpdateExpireTime = time.Hour * 24 * 15 / 3
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

// HSetUserLikedVideo 用户点赞视频批量写入cache
func HSetUserLikedVideo(uid uint, vid []uint) {
	key := UserLikedVideoKeyV3(uid)

	fields := make(map[string]interface{})
	for _, v := range vid {
		fields[strconv.Itoa(int(v))] = true
	}

	resHSet := RDB.HSet(Ctx, key, fields)
	if err := resHSet.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf("HSetUserLikedVideo RDB.HMSet cmd:%v", resHSet.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("HSetUserLikedVideo RDB.HMSet cmd:%v", resHSet.String()))
	}
}

// IsUserLikedVideo 用户是否为某个视频点赞(单个)
//
// state 0 查询成功；1 不存在该key,需要新建；2 冷数据,需回源；3 查询失败，需要回源
func IsUserLikedVideo(uid uint, vid uint) (liked bool, state int) {
	key := UserLikedVideoKeyV3(uid)

	// 查询是否存在
	resHMGet := RDB.HMGet(Ctx, key, sMinVid, sTTL, strconv.Itoa(int(vid)))
	if resHMGet.Err() != nil {
		logger.Logger.Error(fmt.Sprintf("IsUserLikedVideo cmd:%v", resHMGet.String()))
		state = 3
		return
	} else {
		logger.Logger.Info(fmt.Sprintf("IsUserLikedVideo cmd:%v", resHMGet.String()))
	}

	sMinVidVal, ok1 := resHMGet.Val()[0].(string)
	sTtlVal, ok2 := resHMGet.Val()[1].(string)
	if !ok1 || !ok2 { // 不存在该key，需重新构建
		state = 1
		return
	}

	minVid, _ := strconv.Atoi(sMinVidVal)
	ttl, _ := strconv.Atoi(sTtlVal)

	// 剩余时间小于三分之一，重新续期
	// 此处失败不影响服务，打印日志即可，不用向上层传递
	if time.Unix(int64(ttl), 0).Sub(time.Now()) < time.Hour*24*15/3 {
		resExpire := RDB.Expire(Ctx, key, time.Hour*24*15)
		if err := resExpire.Err(); err != nil {
			logger.Logger.Error(fmt.Sprintf("IsUserLikedVideo RDB.Expire cmd:%v", resExpire.String()))
		} else {
			logger.Logger.Info(fmt.Sprintf("IsUserLikedVideo RDB.Expire cmd:%v", resExpire.String()))
		}
	}

	if resHMGet.Val()[2] != nil {
		liked = true
		return
	}
	// 判断是否为冷数据
	if vid < uint(minVid) {
		state = 2
		return
	}
	return
}

// IsUserLikedVideos 用户是否为某些视频点赞(批量)
//
// result key为vid  <->  0未点赞；1为点赞；2 冷数据,需回源
//
// state 0 成功；1 不存在该key,需要新建；2 查询失败,需回源
func IsUserLikedVideos(uid uint, videoIds []uint) (result map[uint]int, state int) {
	key := UserLikedVideoKeyV3(uid)

	fields := make([]string, len(videoIds)+2)
	fields[0], fields[1] = sMinVid, sTTL
	// 将整数切片转换为字符串切片
	for i, num := range videoIds {
		fields[i+2] = strconv.FormatInt(int64(num), 10)
	}

	// 查询是否存在
	resHMGet := RDB.HMGet(Ctx, key, fields...)
	if resHMGet.Err() != nil {
		logger.Logger.Error(fmt.Sprintf("IsUserLikedVideos cmd:%v", resHMGet.String()))
		state = 2
		return
	} else {
		logger.Logger.Info(fmt.Sprintf("IsUserLikedVideos cmd:%v", resHMGet.String()))
	}

	sMinVidVal, ok1 := resHMGet.Val()[0].(string)
	sTtlVal, ok2 := resHMGet.Val()[1].(string)
	if !ok1 || !ok2 { // 不存在该key，需重新构建
		state = 1
		return
	}

	minVid, _ := strconv.Atoi(sMinVidVal)
	ttl, _ := strconv.Atoi(sTtlVal)

	state = 0 // 查询成功，state标为0

	vidRes := resHMGet.Val()[2:]
	result = make(map[uint]int, len(videoIds))

	for i, vid := range videoIds {
		if _, ok := vidRes[i].(string); ok { // 查询到点赞
			result[vid] = 1
		} else if vid >= uint(minVid) { // 热数据，用户没点赞
			result[vid] = 0
		} else { // 冷数据，需要回源
			result[vid] = 2
		}
	}

	// 判断是否超过域值，是否需要续期
	judgeRebuildVideoLikedVideos(key, ttl)

	return
}

// BuildUserLikedVideos 构建用户点赞视频缓存
func BuildUserLikedVideos(key string, vid []uint, minVid uint) {
	var err error

	ttl := time.Now().AddDate(0, 0, 15).Unix() // 一个月的过期时间

	fields := make(map[string]interface{})
	fields[sMinVid] = minVid
	fields[sTTL] = ttl
	for _, v := range vid {
		fields[strconv.Itoa(int(v))] = true
	}

	resHSet := RDB.HSet(Ctx, key, fields)
	if err = resHSet.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf("BuildUserLikedVideos RDB.HMSet cmd:%v", resHSet.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("BuildUserLikedVideos RDB.HMSet cmd:%v", resHSet.String()))
	}

	resExpire := RDB.Expire(Ctx, key, time.Hour*24*15)
	if err = resExpire.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf("BuildUserLikedVideos RDB.Expire cmd:%v", resExpire.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("BuildUserLikedVideos RDB.Expire cmd:%v", resExpire.String()))
	}
}

func rebuildUserLikedVideos(key string) {
	resHGetAll := RDB.HGetAll(Ctx, key)
	if err := resHGetAll.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf("rebuildUserLikedVideos RDB.HGetAll cmd:%v", resHGetAll.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("rebuildUserLikedVideos RDB.HGetAll cmd:%v", resHGetAll.String()))
	}

	videoIds := make([]uint, 0, len(resHGetAll.Val())-2)
	for k := range resHGetAll.Val() {
		if num, err := strconv.Atoi(k); err == nil {
			videoIds = append(videoIds, uint(num))
		}
	}

	sort.Slice(videoIds, func(i, j int) bool {
		return videoIds[i] > videoIds[j]
	})

	if len(videoIds) < MaxFieldNum {
		logger.Logger.Info(fmt.Sprintf("rebuildUserLikedVideos len(videoIds):%v < MaxFieldNum:%v 不需要重建", len(videoIds), MaxFieldNum))
		return
	}

	// 只是为了更新minVid、ttl、过期时间，不用传vid切片
	BuildUserLikedVideos(key, nil, videoIds[SetFieldNum-1])

	var builder strings.Builder
	lastVideos := videoIds[SetFieldNum:]
	for i, v := range lastVideos {
		builder.WriteString(strconv.Itoa(int(v)))
		if i != len(lastVideos)-1 {
			builder.WriteString(" ")
		}
	}
	fields := builder.String()

	resDel := RDB.HDel(Ctx, key, fields)
	if err := resDel.Err(); err != nil {
		logger.Logger.Error(fmt.Sprintf("rebuildUserLikedVideos RDB.HDel cmd:%v", resDel.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("rebuildUserLikedVideos RDB.HDel cmd:%v", resDel.String()))
	}
}

// 判断数量是否超过域值，是否需要续期
// 此处失败不影响服务，打印日志即可，不用向上层传递
func judgeRebuildVideoLikedVideos(key string, ttl int) {
	resLen := RDB.HLen(Ctx, key)
	if resLen.Err() != nil {
		logger.Logger.Error(fmt.Sprintf("judgeRebuildVideoLikedVideos RDB.HLen cmd:%v", resLen.String()))
	} else {
		logger.Logger.Info(fmt.Sprintf("judgeRebuildVideoLikedVideos RDB.HLen cmd:%v", resLen.String()))
	}

	if resLen.Val() > MaxFieldNum { // 数量超过域值
		go rebuildUserLikedVideos(key)

	} else if time.Unix(int64(ttl), 0).Sub(time.Now()) < userLikedVideoUpdateExpireTime { // 剩余时间小于三分之一，重新续期

		resExpire := RDB.Expire(Ctx, key, userLikedVideoExpireTime)
		if err := resExpire.Err(); err != nil {
			logger.Logger.Error(fmt.Sprintf("judgeRebuildVideoLikedVideos RDB.Expire cmd:%v", resExpire.String()))
		} else {
			logger.Logger.Info(fmt.Sprintf("judgeRebuildVideoLikedVideos RDB.Expire cmd:%v", resExpire.String()))
		}
	}
}
