package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	id      int
	Title   string
	Content string
	Desc    string
}

func ToValidData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/valid_data.html", nil)
}

func DoValidData(ctx *gin.Context) {
	mmmm := struct {
		Title   string `form:"title" json:"title" uri:"title" binding:"required"`     // 不能为空的字段
		Content string `form:"content" json:"content" uri:"content" binding:"min=10"` // 最小长度是10
		Desc    string `form:"desc" json:"desc" uri:"desc"`
	}{}

	err := ctx.ShouldBind(&mmmm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", mmmm)
	ctx.String(http.StatusOK, "验证成功", nil)
}
