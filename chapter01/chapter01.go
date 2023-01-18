package chapter01

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 指定到 index.html 文件

func Index(context *gin.Context) {
	name := "zhiliao"
	context.HTML(http.StatusOK, "index/index.html", name)
}

// 指定到 news.html 文件

func News(context *gin.Context) {
	//name := "news"
	context.HTML(http.StatusOK, "news/news.html", nil)
}
