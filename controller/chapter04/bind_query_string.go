package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindQueryString(ctx *gin.Context) {
	mm := struct {
		Username string `form:"username" json:"username"`
		Age      int    `form:"age" json:"age"`
		Addr     string `form:"addr" json:"addr"`
	}{}

	//err := ctx.ShouldBindQuery(&mm)  // 这个可以
	err := ctx.ShouldBind(&mm) // 这个也可以
	if err != nil {
		fmt.Println("绑定错误")
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", mm)
	ctx.String(http.StatusOK, "get 请求获取参数")
}
