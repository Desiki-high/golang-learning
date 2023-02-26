package orm

import (
	"golang-learning/gorm-example/sqlite/orm/entity"
	"gorm.io/gorm/clause"
)

func AddUserAndCompany() {
	var user = entity.User{
		Name:      "test",
		CompanyID: 0,
		Company: entity.Company{
			ID:   0,
			Name: "testCompany",
		},
	}
	dataBase.Model(&entity.User{}).Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
}
