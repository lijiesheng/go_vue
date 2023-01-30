package chapter04

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Router(chap04 *gin.RouterGroup) {
	/*------------------------------第四章-------------------------------*/
	/*-----------------------1、数据绑定-------------------------------*/
	// post
	chap04.GET("/to_bind_form", ToBindForm)  // 到这个页面
	chap04.POST("/do_bind_form", DoBindForm) // 提交数据

	// get
	chap04.GET("/bind_query_string", BindQueryString)

	// ajax 发送 json
	chap04.GET("/to_bind_json", ToBindJSON) // 到这个页面
	chap04.POST("/do_bind_json", DoBindJSON)
	// uri 的绑定
	chap04.GET("/do_bind_uri/:name/:age/:addr", BindUri)

	/*-----------------------2、数据验证-------------------------------*/
	chap04.GET("/valid", ToValidData)
	chap04.POST("/do_valid", DoValidData)

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("len_valid", Len6Valid)
	}

	// beego 数据验证 略过
}
