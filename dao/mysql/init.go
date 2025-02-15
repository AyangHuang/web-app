package mysql

import (
	"fmt"

	"aichat/common/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(conf *config.MySQLConfig) {
	var err error
	db, err = gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database, err: %w", err))
	}
	_db, err := db.DB()
	if err != nil {
		panic(err)
	}
	_db.SetMaxOpenConns(conf.MaxOpenConns) //设置数据库连接池最大连接数
	_db.SetMaxIdleConns(conf.MaxIdleConns) //连接池最大允许的空闲连接数
}
