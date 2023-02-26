package orm

import (
	"golang-learning/gorm-example/sqlite/orm/entity"
)

//查询不对数据库产生修改，查询失败则返回的结果为零值

// GetUsers 获取所有User
func GetUsers() []entity.User {
	var users []entity.User
	dataBase.Model(&entity.User{}).First(&users)
	return users
}
