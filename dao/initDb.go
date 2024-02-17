package dao

import (
	"fmt"
	config2 "github.com/TJxiaobao/OJ/utils/config"
	"github.com/TJxiaobao/OJ/utils/logger_oj"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB = InitDb()
var Driver = "mysql"
var config = config2.LoadConfig()

func InitDb() *gorm.DB {
	if config.Db.Driver != Driver {
		logger_oj.LogToFile("mysql connect error")
		return nil
	}
	path := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.Db.User, config.Db.Password, config.Db.Address, config.Db.Port, config.Db.Database, config.Db.Charset)
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Print("init db err", err)
		return nil
	}
	return db
}
