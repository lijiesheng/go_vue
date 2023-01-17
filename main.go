package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", Index)

	// template/* 意思是找当前项目路径下template文件夹下所有的html文件
	r.LoadHTMLGlob("template/*")

	r.Run(":8083")
}

// 指定到 index.html 文件
func Index(context *gin.Context) {
	name := "zhiliao"
	context.HTML(http.StatusOK, "index.html", name)
}
