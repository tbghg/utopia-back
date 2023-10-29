package implement

import (
	mysqlCfg "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"utopia-back/config"
	"utopia-back/model"
	"utopia-back/pkg/logger"
)

var myDb *gorm.DB

type CenterDal struct {
	FavoriteDal *FavoriteDal
	UserDal     *UserDal
	VideoDal    *VideoDal
	LikeDal     *LikeDal
	FollowDal   *FollowDal
	TestUserDal *TestUserDal
}

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
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{
		Logger: logger.NewGormLogger(logger.Logger, gormLogger.Warn),
	})
	if err != nil {
		return err
	}

	myDb = db
	return nil
}

func TestInit() error {
	// 数据库配置
	cfg := mysqlCfg.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_ADDR"),
		DBName:               os.Getenv("MYSQL_DBNAME"),
		Loc:                  time.Local,
		ParseTime:            true,
		AllowNativePasswords: true, // 允许原生密码
	}
	// 日志配置
	testLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second,     // 慢 SQL 阈值
			LogLevel:      gormLogger.Info, // Log level
			Colorful:      true,            // 彩色打印
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{
		Logger: testLogger,
	})
	if err != nil {
		return err
	}

	myDb = db
	return nil
}

func InitTable() {
	// 初始化数据表
	_ = myDb.AutoMigrate(&model.TestUser{})
	_ = myDb.AutoMigrate(&model.User{})
	_ = myDb.AutoMigrate(&model.Video{})
	_ = myDb.AutoMigrate(&model.Like{})
	_ = myDb.AutoMigrate(&model.Favorite{})
	_ = myDb.AutoMigrate(&model.Follow{})
}

func NewCenterDal() *CenterDal {
	return &CenterDal{
		FavoriteDal: &FavoriteDal{
			Db: myDb,
		},
		UserDal: &UserDal{
			Db: myDb,
		},
		VideoDal: &VideoDal{
			Db: myDb,
		},
		LikeDal: &LikeDal{
			Db: myDb,
		},
		FollowDal: &FollowDal{
			Db: myDb,
		},
		TestUserDal: &TestUserDal{
			Db: myDb,
		},
	}
}
