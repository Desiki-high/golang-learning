package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func main() {
	i := 0
	c := cron.New(cron.WithSeconds())
	spec := fmt.Sprintf("@every %s", 3*time.Second)

	if _, err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	}); err != nil {
		fmt.Println(err)
		return
	}

	spec = "*/5 * * * * *" // 每隔1s执行一次，cron格式（秒，分，时，日，月，周(0-6 周日到周六)）,秒需要在创建实例的时选择开启
	if _, err := c.AddFunc(spec, func() {
		log.Printf("time = %d\n", time.Now().Unix())
	}); err != nil {
		fmt.Println(err)
		return
	}
	c.Start()

	select {} //refuse the ending of the main func
}
