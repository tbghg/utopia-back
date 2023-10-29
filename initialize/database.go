package initialize

import (
	"utopia-back/database/implement"
)

func InitDB() {
	// 初始化数据库
	_ = implement.Init()
	// 初始化数据库表
	implement.InitTable()
}
