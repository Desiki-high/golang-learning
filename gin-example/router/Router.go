package router

import (
	"github.com/gin-gonic/gin"
	"golang-learning/gin-example/handler"
)

func ControllerRouter() *gin.Engine {
	router := gin.Default()

	//路由组
	RouterGroup := router.Group("/index")
	{
		RouterGroup.GET("/get", handler.DemoController)
	}

	return router
}
