package handler

import (
	"golang-learning/gin-example/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetController(ctx *gin.Context) {
	//使用结构体构造json
	ctx.JSON(http.StatusOK, model.DemoSuccess())
}
