package main

import (
	"fmt"
	"golang-learning/gorm-example/sqlite/orm"
)

func main() {
	orm.SqliteConnect()
	orm.AddUserAndCompany()
	for user, id := range orm.GetUsers() {
		fmt.Println(user)
		fmt.Println(id)
	}
}
