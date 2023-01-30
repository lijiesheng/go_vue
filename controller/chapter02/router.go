package chapter02

import "github.com/gin-gonic/gin"

func Router(chap02 *gin.RouterGroup) {
	/*---------------------------------1、数据的渲染  【后端给前端数据】-------------------------------*/

	// 字符串渲染
	chap02.GET("/user", User)

	// 结构体渲染
	chap02.GET("/user_struct", UserInfoStruct)

	// 数组渲染
	chap02.GET("/arr", ArrController)

	// 结构体数组渲染
	chap02.GET("/arr_struct", ArrStructController)

	// map的渲染
	chap02.GET("/map1", MapController)

	// map 嵌套map
	chap02.GET("/map_map", MapMapController)

	// map + 结构体
	chap02.GET("/map_struct", MapStructController)

	// 切片
	chap02.GET("/slice", SliceController)

	chap02.GET("/slice_struct", SliceStructController)

	/*------------------------------2、获取路由参数  【前端给后端数据】-------------------------------*/

	// 1、路由参数获取
	chap02.GET("/param/:id", Param1)       // 获取一个参数   路由必须是/param/xxx
	chap02.GET("/param/:id/:name", Param2) // 获取多个参数   路由必须是/param/xxx/xxx
	chap02.GET("/param2/*id", Param3)      // 路由可以是 /param2 或者 param2/123 param2/123/123 都可以匹配到

	/*------------------------------2、获取GET请求参数  【前端给后端数据】-------------------------------*/

	// 2、? 后面的参数获取
	chap02.GET("/query", GetQueryData)

	// 3、? 后面获取数组  /user?name=1,2,3,4,5
	chap02.GET("/query_arr", GetQueryArr)

	// 4、? 后面获取 map /user?name["aa"]=hallen&name["bb"]=hallen2  其中 aa bb 是key ，name 是map的一个实例
	chap02.GET("/query_map", GetQueryMap)

	/*------------------------------3、获取POST请求参数 【前端给后端数据】-------------------------------*/
	chap02.GET("/to_user_add", ToUserAdd)  // 进入这个页面
	chap02.POST("/do_user_add", DoUserAdd) // 提交数据

	// 使用 ajax
	chap02.GET("/to_user_add3", ToUserAdd3)  // 进入这个页面
	chap02.POST("/do_user_add3", DoUserAdd3) // 提交数据

	/*------------------------------4、参数绑定-------------------------------*/

	// post 请求参数绑定
	chap02.GET("/to_user_add4", ToUserAdd4)  // 进入这个页面
	chap02.POST("/do_user_add4", DoUserAdd4) // 提交数据

	// get 请求参数绑定 ShouldBindQuery

	/*------------------------------5、文件上传-------------------------------*/
	// 1、form 表单，单文件
	chap02.GET("/to_upload1", ToUpload1)  // 进入这个页面
	chap02.POST("/do_upload1", DoUpload1) // 提交数据

	// 2、form 表单，单文件
	chap02.GET("/to_upload2", ToUpload2)  // 进入这个页面
	chap02.POST("/do_upload2", DoUpload2) // 提交数据

	// 3、ajax，多文件
	chap02.GET("/to_upload3", ToUpload3)  // 进入这个页面
	chap02.POST("/do_upload3", DoUpload3) // 提交数据

	// 4、ajax，多文件
	chap02.GET("/to_upload4", ToUpload4)  // 进入这个页面
	chap02.POST("/do_upload4", DoUpload4) // 提交数据

	/*------------------------------6、其他输出格式-------------------------------*/
	// 1、json
	chap02.GET("/output_json", OutJson)
	// 2、AsciiJSON
	chap02.GET("/output_AsciiJSON", OutAsciiJSON)
	// 3、JSONP
	chap02.GET("/output_JSONP", OutJSONP)
	// 3、OutAsciiPureJSON
	chap02.GET("/output_AsciiPureJSON", OutAsciiPureJSON)
	// 3、OutAsciiSecureJSON
	chap02.GET("/output_asciiSecureJSON", OutAsciiSecureJSON)
	// 3、OutAsciiXML
	chap02.GET("/output_outAsciiXML", OutAsciiXML)
	// 3、OutAsciiYAML
	chap02.GET("/output_asciiYAML", OutAsciiYAML)

	/*------------------------------6、重定向-------------------------------*/
	chap02.GET("/redirectA", RedirectA)
	chap02.GET("/redirectB", RedirectB)
}
