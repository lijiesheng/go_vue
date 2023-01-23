package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToBindForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_form.html", nil)
}

func DoBindForm(ctx *gin.Context) {
	m := struct {
		Username string `form:"username" json:"username"`
		Age      int    `form:"age" json:"age"`
		Addr     string `form:"addr" json:"addr"`
	}{}
	err := ctx.ShouldBind(&m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", m)
	fmt.Println()
	ctx.String(http.StatusOK, "提交成功")
}
