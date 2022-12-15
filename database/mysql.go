package database

import (
	"technical_test_skyshi/config"
	"technical_test_skyshi/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql() *gorm.DB {
	config.InitEnvDB()
	dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")" + "/" + config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Activity{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Todo{})
	if err != nil {
		panic(err)
	}

	return db
}
