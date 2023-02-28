package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("viper/config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.Get("name"))
	fmt.Println(viper.AllKeys())
	fmt.Println(viper.AllSettings())

}
