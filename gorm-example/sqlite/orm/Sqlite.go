package orm

import (
	"github.com/glebarez/sqlite"
	"golang-learning/gorm-example/sqlite/orm/entity"
	"gorm.io/gorm"
)

var dataBase *gorm.DB

func SqliteConnect() {
	db, err := gorm.Open(sqlite.Open("gorm-example/sqlite/config.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&entity.User{}, &entity.Company{})
	if err != nil {
		panic("数据库关系迁移失败")
	}
	dataBase = db
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
}
