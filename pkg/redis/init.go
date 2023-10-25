package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"utopia-back/pkg/logger"
)

var RDB *redis.Client

func Init() {
	ctx := context.Background()
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	RDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		logger.Logger.Error("redis connect failed")
		return
	}

	logger.Logger.Info("redis connect success")

}
