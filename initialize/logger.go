package initialize

import (
	"utopia-back/config"
	"utopia-back/pkg/logger"
)

// InitLogger 初始化日志
func InitLogger() {
	logger.InitLogger(
		config.V.GetInt("log.max_size"),
		config.V.GetInt("log.max_backup"),
		config.V.GetInt("log.max_age"),
		config.V.GetBool("log.compress"),
		config.V.GetString("log.level"),
	)
}
