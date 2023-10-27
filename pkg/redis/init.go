package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"utopia-back/config"
	"utopia-back/pkg/logger"
)

var RDB *redis.Client

func Init() {
	ctx := context.Background()

	RDB = redis.NewClient(&redis.Options{
		Addr:     config.V.GetString("redis.addr"),
		Password: config.V.GetString("redis.password"),
		DB:       config.V.GetInt("redis.db"),
	})
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		logger.Logger.Error("redis连接失败", zap.String("err", err.Error()))
		return
	}

	logger.Logger.Info("redis连接成功")

}
