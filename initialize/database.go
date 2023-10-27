package initialize

import (
	"utopia-back/database"
	"utopia-back/model"
)

func InitDB() {
	// 初始化数据库
	_ = database.Init()
	// 初始化数据表
	_ = database.DB.AutoMigrate(&model.TestUser{})
	_ = database.DB.AutoMigrate(&model.User{})
	_ = database.DB.AutoMigrate(&model.Video{})
	_ = database.DB.AutoMigrate(&model.Like{})
	_ = database.DB.AutoMigrate(&model.Favorite{})
	_ = database.DB.AutoMigrate(&model.Follow{})
}
