package main

import (
	"golang-learning/gin/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//Debug运行模式
	gin.SetMode(gin.DebugMode)
	//生产运行模式
	//gin.SetMode(gin.ReleaseMode)

	//gin.DisableConsoleColor()

	// 记录到文件。
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//Listen and Server in localhost:10000
	r := router.Router()
	s := &http.Server{
		Addr:         ":8848",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}
