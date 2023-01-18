package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", Index)
	r.GET("/user", User)
	r.GET("/news", News)

	// template/* 意思是找当前项目路径下template文件夹下所有的html文件
	r.LoadHTMLGlob("template/**/*") // 所有的 html 文件都是 /**/xx.html

	// 这个 /static 对应的 <link rel="stylesheet" href="/static/css/index.css"> 中的 /static
	// root 的 static 是代表到 static 文件中去找
	r.Static("/static", "static")

	r.Run(":8083")
}

// 指定到 index.html 文件

func Index(context *gin.Context) {
	name := "zhiliao"
	context.HTML(http.StatusOK, "index/index.html", name)
}

// 指定到 user.html 文件

func User(context *gin.Context) {
	//name := "user"
	context.HTML(http.StatusOK, "user/user.html", nil)
}

// 指定到 news.html 文件

func News(context *gin.Context) {
	//name := "news"
	context.HTML(http.StatusOK, "news/news.html", nil)
}
