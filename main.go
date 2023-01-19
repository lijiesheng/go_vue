package main

import (
	"Gin_Vue/chapter01"
	"Gin_Vue/chapter02"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", chapter01.Index)
	r.GET("/news", chapter01.News)

	/*---------------------------------1、数据的渲染  【后端给前端数据】-------------------------------*/

	// 字符串渲染
	r.GET("/user", chapter02.User)

	// 结构体渲染
	r.GET("/user_struct", chapter02.UserInfoStruct)

	// 数组渲染
	r.GET("/arr", chapter02.ArrController)

	// 结构体数组渲染
	r.GET("/arr_struct", chapter02.ArrStructController)

	// map的渲染
	r.GET("/map1", chapter02.MapController)

	// map 嵌套map
	r.GET("/map_map", chapter02.MapMapController)

	// map + 结构体
	r.GET("/map_struct", chapter02.MapStructController)

	// 切片
	r.GET("/slice", chapter02.SliceController)

	r.GET("/slice_struct", chapter02.SliceStructController)

	/*------------------------------2、获取GET请求参数  【前端给后端数据】-------------------------------*/
	r.GET("/param/:id", chapter02.Param1)       // 获取一个参数   路由必须是/param/xxx
	r.GET("/param/:id/:name", chapter02.Param2) // 获取多个参数   路由必须是/param/xxx/xxx
	r.GET("/param2/*id", chapter02.Param3)      // 路由可以是 /param2 或者 param2/123 param2/123/123 都可以匹配到
	/*------------------------------3、获取POST请求参数 【前端给后端数据】-------------------------------*/

	/*------------------------------4、参数绑定-------------------------------*/

	// template/* 意思是找当前项目路径下template文件夹下所有的html文件
	r.LoadHTMLGlob("template/**/*") // 所有的 html 文件都是 /**/xx.html

	// 这个 /static 对应的 <link rel="stylesheet" href="/static/css/index.css"> 中的 /static
	// root 的 static 是代表到 static 文件中去找
	r.Static("/static", "static")

	r.Run(":8083")
}
