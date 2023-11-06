package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
	"utopia-back/job/common"
	"utopia-back/pkg/logger"
)

type VideoPopularItem struct {
	Vid   uint
	Score float64
}

// PopularVideoKey 热门视频Key
func PopularVideoKey(version int) string {
	return fmt.Sprintf("video:popular:%d", version)
}

// VideoInfoKey 视频信息Key
func VideoInfoKey(videoId uint) string {
	return fmt.Sprintf("video:info:%v", videoId)
}

// BuildPopularVideo 构建热门视频
func BuildPopularVideo(key string, videoItem []*VideoPopularItem) {
	members := make([]redis.Z, len(videoItem))
	for i := range videoItem {
		members[i].Member = videoItem[i].Vid
		members[i].Score = videoItem[i].Score
	}

	// 使用事务执行删除和重新写入 ZSET 的操作
	err := RDB.Watch(Ctx, func(tx *redis.Tx) error {
		// 开始事务
		_, err := tx.TxPipelined(Ctx, func(pipe redis.Pipeliner) error {
			// 删除 ZSET
			pipe.ZRemRangeByRank(Ctx, key, 0, -1)
			// 重新写入 ZSET
			pipe.ZAdd(Ctx, key, members...)
			return nil
		})
		return err
	})

	if err == redis.TxFailedErr {
		logger.Logger.Error(fmt.Sprintf("BuildPopularVideo RDB.Watch redis.TxFailedErr err:%+v", err))
		return
	} else if err != nil {
		logger.Logger.Error(fmt.Sprintf("BuildPopularVideo RDB.Watch err:%+v", err))
		return
	}
	return
}

func GetPopularVideo(version int, score float64, count int64) (videoIds []uint, nextScore float64, nextVersion int, ok bool) {
	var (
		key      string
		maxScore string
	)

	if version == 0 {
		version = common.GetPopularVideoVersion()
	}
	nextVersion = version
	key = PopularVideoKey(version)

	if score == 0 {
		maxScore = "+inf"
	} else {
		maxScore = fmt.Sprintf("(%f", score)
	}

	// 从 ZSET 中按分数从大到小获取前 count 个成员
	result, err := RDB.ZRevRangeByScoreWithScores(Ctx, key, &redis.ZRangeBy{
		Max:    maxScore,
		Offset: 0,
		Count:  count - 1,
	}).Result()

	if err != nil || len(result) == 0 {
		logger.Logger.Error(fmt.Sprintf("GetPopularVideo RDB.ZRevRangeByScoreWithScores err:%+v", err))
		return nil, -1, nextVersion, err == nil
	}

	for _, z := range result {
		vid, _ := strconv.Atoi(z.Member.(string))
		videoIds = append(videoIds, uint(vid))
	}

	nextScore = result[len(result)-1].Score
	return videoIds, nextScore, nextVersion, true
}

func VideoPersistentKey(inputKey string) string {
	return fmt.Sprintf("video:persistent:%s", inputKey)
}

func SetVideoPersistent(inputKey string, vid uint) {
	err := RDB.Set(Ctx, VideoPersistentKey(inputKey), vid, 30*time.Minute).Err()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("SetVideoPersistent RDB.Set err:%+v", err))
	}
}

func GetVideoPersistent(inputKey string) uint {
	result := RDB.Get(Ctx, VideoPersistentKey(inputKey))
	if result.Err() != nil {
		logger.Logger.Error(fmt.Sprintf("SetVideoPersistent RDB.Set err:%+v", result.Err()))
	}
	vid, _ := strconv.Atoi(result.Val())
	return uint(vid)
}
