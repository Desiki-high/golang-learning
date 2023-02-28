package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type User struct {
	Id   int
	Name string
}

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

func main() {
	//解析MySQL连接参数
	yamlFile, err := os.ReadFile("gorm/mysql/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	cfg := Config{}
	err = yaml.Unmarshal(yamlFile, &cfg)

	//拼接dsn参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	var user User
	result := db.Where("id", 1).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(user)
}
