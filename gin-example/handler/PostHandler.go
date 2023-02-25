package handler

import (
	"fmt"
	"golang-learning/gin-example/model"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func DemoController(ctx *gin.Context) {
	//回传字符串
	ctx.String(http.StatusOK, "DemoGet")
}

func DemoGetController(ctx *gin.Context) {
	//从url中取数据
	ip := ctx.Params.ByName("ip")

	//回传json
	//使用gin.H{}直接构造需要的json
	ctx.JSON(http.StatusOK, gin.H{
		"text": ip, "" +
			"" +
			"int": 200,
	})

	//使用结构体构造json
	//ctx.JSON(http.StatusOK, model.DemoSuccess())
}

func DemoPostController(ctx *gin.Context) {
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

func DemoNextGetController(ctx *gin.Context) {
	//从body中读表单数据
	name := ctx.PostForm("name")

	//没有拿到数据
	if name == "" {
		name = "没有设置name"
	}

	//自定义结构体转json
	Demo := model.Demo{
		DemoText:   name,
		DemoInt:    200,
		DemoOption: nil,
	}
	ctx.JSON(http.StatusOK, Demo)
}
