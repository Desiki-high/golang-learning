package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.Default()
	router = indexRouter(router)

	return router
}
