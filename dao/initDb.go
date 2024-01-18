package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = InitDb()

func InitDb() *gorm.DB {
	path := "root:abcdi2124Jcke23@tcp (192.168.1.8:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		log.Print("init db err", err)
		return nil
	}
	return db
}
