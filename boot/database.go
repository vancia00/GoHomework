package boot

import (
	"GoHomework/global"
	"GoHomework/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func InitMysql() {

	dsn := "root:sheide5201314.1@tcp(127.0.0.1:3306)/mysql_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("falied to connect to database！")
	}
	sqlDB := db.DB()
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxIdleTime(180 * time.Second)
	sqlDB.SetConnMaxLifetime(1800 * time.Second)
	global.GloDb = db
	global.GloDb.AutoMigrate(&model.User{}, &model.Question{}, &model.Comment{})

}
