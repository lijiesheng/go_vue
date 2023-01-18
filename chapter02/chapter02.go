package chapter02

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 指定到 user.html 文件

func User(context *gin.Context) {
	name := "user"
	context.HTML(http.StatusOK, "user/user.html", name)
}
