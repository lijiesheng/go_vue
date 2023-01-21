package chapter02

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OutJson(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}

func OutJSONP(context *gin.Context) {
	context.JSONP(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}
