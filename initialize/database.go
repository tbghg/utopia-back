package initialize

import (
	"utopia-back/database"
	"utopia-back/model"
)

func InitDB() {
	// 初始化数据库
	database.Init()
	// 初始化数据表
	database.DB.AutoMigrate(&model.TestUser{})
}
