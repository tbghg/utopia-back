package database

import (
	mysqlCfg "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"utopia-back/config"
)

var DB *gorm.DB

func Init() error {
	// 数据库配置
	cfg := mysqlCfg.Config{
		User:      config.V.GetString("mysql.user"),
		Passwd:    config.V.GetString("mysql.password"),
		Net:       "tcp",
		Addr:      config.V.GetString("mysql.addr"),
		DBName:    config.V.GetString("mysql.dbname"),
		Loc:       time.Local,
		ParseTime: true,
		// 允许原生密码
		AllowNativePasswords: true,
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
