package handler

import (
	"fmt"
	"golang-learning/gin-example/model"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func PostController(ctx *gin.Context) {
	var demo model.Demo
	err := ctx.ShouldBind(&demo)
	if err != nil { //err不为空 说明将json映射成struct失败了 说明Demo中没有demoText 或者为空
		ctx.String(http.StatusBadRequest, `the body format is wrong`)
	} else {
		//遍历结构体对象
		t := reflect.TypeOf(demo)
		v := reflect.ValueOf(demo)
		for i := 0; i < t.NumField(); i++ {
			fmt.Printf("%s -- %v \n", t.Field(i).Tag, v.Field(i).Interface())
		}
		ctx.JSON(http.StatusOK, model.DemoSuccess())
	}
}
