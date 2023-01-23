package main

import (
	"Gin_Vue/chapter01"
	"Gin_Vue/chapter02"
	"Gin_Vue/chapter03"
	"Gin_Vue/chapter04"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	/*------------------------------2、获取路由参数  【前端给后端数据】-------------------------------*/

	// 1、路由参数获取
	r.GET("/param/:id", chapter02.Param1)       // 获取一个参数   路由必须是/param/xxx
	r.GET("/param/:id/:name", chapter02.Param2) // 获取多个参数   路由必须是/param/xxx/xxx
	r.GET("/param2/*id", chapter02.Param3)      // 路由可以是 /param2 或者 param2/123 param2/123/123 都可以匹配到

	/*------------------------------2、获取GET请求参数  【前端给后端数据】-------------------------------*/

	// 2、? 后面的参数获取
	r.GET("/query", chapter02.GetQueryData)

	// 3、? 后面获取数组  /user?name=1,2,3,4,5
	r.GET("/query_arr", chapter02.GetQueryArr)

	// 4、? 后面获取 map /user?name["aa"]=hallen&name["bb"]=hallen2  其中 aa bb 是key ，name 是map的一个实例
	r.GET("/query_map", chapter02.GetQueryMap)

	/*------------------------------3、获取POST请求参数 【前端给后端数据】-------------------------------*/
	r.GET("/to_user_add", chapter02.ToUserAdd)  // 进入这个页面
	r.POST("/do_user_add", chapter02.DoUserAdd) // 提交数据

	// 使用 ajax
	r.GET("/to_user_add3", chapter02.ToUserAdd3)  // 进入这个页面
	r.POST("/do_user_add3", chapter02.DoUserAdd3) // 提交数据

	/*------------------------------4、参数绑定-------------------------------*/

	// post 请求参数绑定
	r.GET("/to_user_add4", chapter02.ToUserAdd4)  // 进入这个页面
	r.POST("/do_user_add4", chapter02.DoUserAdd4) // 提交数据

	// get 请求参数绑定 ShouldBindQuery

	/*------------------------------5、文件上传-------------------------------*/
	// 1、form 表单，单文件
	r.GET("/to_upload1", chapter02.ToUpload1)  // 进入这个页面
	r.POST("/do_upload1", chapter02.DoUpload1) // 提交数据

	// 2、form 表单，单文件
	r.GET("/to_upload2", chapter02.ToUpload2)  // 进入这个页面
	r.POST("/do_upload2", chapter02.DoUpload2) // 提交数据

	// 3、ajax，多文件
	r.GET("/to_upload3", chapter02.ToUpload3)  // 进入这个页面
	r.POST("/do_upload3", chapter02.DoUpload3) // 提交数据

	// 4、ajax，多文件
	r.GET("/to_upload4", chapter02.ToUpload4)  // 进入这个页面
	r.POST("/do_upload4", chapter02.DoUpload4) // 提交数据

	/*------------------------------6、其他输出格式-------------------------------*/
	// 1、json
	r.GET("/output_json", chapter02.OutJson)
	// 2、AsciiJSON
	r.GET("/output_AsciiJSON", chapter02.OutAsciiJSON)
	// 3、JSONP
	r.GET("/output_JSONP", chapter02.OutJSONP)
	// 3、OutAsciiPureJSON
	r.GET("/output_AsciiPureJSON", chapter02.OutAsciiPureJSON)
	// 3、OutAsciiSecureJSON
	r.GET("/output_asciiSecureJSON", chapter02.OutAsciiSecureJSON)
	// 3、OutAsciiXML
	r.GET("/output_outAsciiXML", chapter02.OutAsciiXML)
	// 3、OutAsciiYAML
	r.GET("/output_asciiYAML", chapter02.OutAsciiYAML)

	/*------------------------------6、重定向-------------------------------*/
	r.GET("/redirectA", chapter02.RedirectA)
	r.GET("/redirectB", chapter02.RedirectB)

	// template/* 意思是找当前项目路径下template文件夹下所有的html文件
	r.LoadHTMLGlob("template/**/*") // 所有的 html 文件都是 /**/xx.html

	// 这个 /static 对应的 <link rel="stylesheet" href="/static/css/index.css"> 中的 /static
	// root 的 static 是代表到 static 文件中去找
	r.Static("/static", "static")

	/*------------------------------7、HTTP-------------------------------*/

	//r.Run(":8083")

	//http.ListenAndServe(":8083", r) // 和上面的是一样的

	s := &http.Server{
		Addr:         ":8083",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	/*------------------------------第三章-------------------------------*/

	/*---------------------------3.1 模板语法----------------------------*/
	r.GET("/test_tpl", chapter03.TestSyncTpl)

	/*---------------------------3.2 模板函数----------------------------*/

	/*---------------------------3.2 自定义模板函数----------------------------*/

	/*------------------------------第四章-------------------------------*/
	/*-----------------------1、数据绑定-------------------------------*/
	// post
	r.GET("/to_bind_form", chapter04.ToBindForm)  // 到这个页面
	r.POST("/do_bind_form", chapter04.DoBindForm) // 提交数据

	// get
	r.GET("/bind_query_string", chapter04.BindQueryString)

	// ajax 发送 json
	r.GET("/to_bind_json", chapter04.ToBindJSON) // 到这个页面
	r.POST("/do_bind_json", chapter04.DoBindJSON)
	// uri 的绑定
	r.GET("/do_bind_uri/:name/:age/:addr", chapter04.BindUri)

	/*-----------------------2、数据验证-------------------------------*/
	r.GET("/valid", chapter04.ToValidData)
	r.POST("/do_valid", chapter04.DoValidData)
	s.ListenAndServe()
}
