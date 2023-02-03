package main

import (
	"Gin_Vue/middle_ware"
	"Gin_Vue/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(middle_ware.Cors()) // 前端跨域问题
	router.Router(r)

	// template/* 意思是找当前项目路径下template文件夹下所有的html文件
	r.LoadHTMLGlob("template/**/*") // 所有的 html 文件都是 /**/xx.html

	// 这个 /static 对应的 <link rel="stylesheet" href="/static/css/index.css"> 中的 /static
	// root 的 static 是代表到 static 文件中去找
	r.Static("/static", "static")

	/*------------------------------7、HTTP-------------------------------*/

	//r.Run(":8083")

	//http.ListenAndServe(":8083", r) // 和上面的是一样的

	s := &http.Server{
		Addr:         ":8090", // 这里不能使用 8080，因为 8080 被前端使用了
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
