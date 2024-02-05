package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB = InitDb()

func InitDb() *gorm.DB {
	path := "root:root@tcp(106.14.116.62:3306)/oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Print("init db err", err)
		return nil
	}
	return db
}
