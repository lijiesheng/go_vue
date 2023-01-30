package chapter04

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToBindJSON(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_json.html", nil)
}

func DoBindJSON(ctx *gin.Context) {
	mmm := struct {
		Username string   `form:"username" json:"username"`
		Password string   `form:"password" json:"password"`
		Loves    []string `form:"loves" json:"loves"`
	}{}

	err := ctx.ShouldBind(&mmm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("mmm = %+v\n", mmm)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "提交成功",
	})
}
