package model

import (
	"github-notify-bot/util"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB :  数据库连接对象
var DB *gorm.DB

// InitSQLite :  初始化数据库
func InitSQLite() (err error) {
	dsn := util.GetCurrentDirectory() + "/conf/subscriber.db"
	DB, err = gorm.Open("sqlite3", dsn)
	if err != nil {
		log.Println(err)
		return err
	}

	//建立 struct - table 映射，不存在的表会自动创建
	DB.AutoMigrate(&Subscriber{})
	DB.AutoMigrate(&Repository{})

	return DB.DB().Ping()
}

// Close :  关闭数据库连接
func Close() {
	DB.Close()
}

