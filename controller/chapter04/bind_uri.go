package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindUri(ctx *gin.Context) {
	mmmm := struct {
		Name string `form:"name" json:"name" uri:"name"`
		Age  string `form:"age" json:"age" uri:"age"`
		Addr string `form:"addr" json:"addr" uri:"addr"`
	}{}

	ctx.ShouldBindUri(&mmmm)
	fmt.Printf("%+v\n", mmmm)
	ctx.String(http.StatusOK, "成功", nil)
}
