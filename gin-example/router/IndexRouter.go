package router

import (
	"golang-learning/gin-example/handler"

	"github.com/gin-gonic/gin"
)

func indexRouter(router *gin.Engine) *gin.Engine {
	//路由组
	indexGroup := router.Group("/index")
	{
		indexGroup.GET("/get", handler.GetController)
		indexGroup.POST("/post", handler.PostController)
	}

	return router
}
