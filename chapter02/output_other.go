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

func OutAsciiJSON(context *gin.Context) {
	context.JSONP(http.StatusOK, gin.H{
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

func OutAsciiPureJSON(context *gin.Context) {
	context.PureJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}

func OutAsciiSecureJSON(context *gin.Context) {
	context.PureJSON(http.StatusOK, gin.H{
		"code":  200,
		"names": []string{"lena", "austin", "foo"},
	})
}

func OutAsciiXML(context *gin.Context) {
	context.XML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}

func OutAsciiYAML(context *gin.Context) {
	context.YAML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"user": gin.H{"name": "zhiliao", "age": 18},
		"html": "<b>Hello, world!</b>",
	})
}
