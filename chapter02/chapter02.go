package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 指定到 user.html 文件

// 字符串渲染
func User(context *gin.Context) {
	name := "ljs"
	context.HTML(http.StatusOK, "user/user.html", name)
}

// 结构体渲染
func UserInfoStruct(context *gin.Context) {
	user := struct {
		Id   int
		Name string
	}{Id: 1, Name: "ljs"}
	context.HTML(http.StatusOK, "chapter02/user_info.html", user)
}

// 数组渲染
func ArrController(context *gin.Context) {
	nums := []int{1, 2, 3, 4, 5}
	context.HTML(http.StatusOK, "chapter02/arr.html", nums)
}

// 结构体数组渲染
func ArrStructController(context *gin.Context) {
	userArr := []struct {
		Id   int
		Name string
	}{{Id: 1, Name: "ljs"}, {Id: 2, Name: "hsm"}}
	context.HTML(http.StatusOK, "chapter02/arrStruct.html", userArr)
}

// map 的渲染
func MapController(context *gin.Context) {
	m := map[string]int{
		"chinese": 80,
		"math":    100,
		"english": 120,
	}
	context.HTML(http.StatusOK, "chapter02/map.html", m)
}

func MapMapController(context *gin.Context) {
	//mm := map[string]map[string]int{
	//	"key1": map[string]int{
	//		"chinese": 80,
	//		"math":    100,
	//		"english": 120,
	//	},
	//	"key2": map[string]int{
	//		"chinese": 1,
	//		"math":    2,
	//		"english": 3,
	//	},
	//}

	mm := map[string]interface{}{
		"key1": map[string]int{
			"chinese": 80,
			"math":    100,
			"english": 120,
		},
		"key2": map[string]int{
			"chinese": 1,
			"math":    2,
			"english": 3,
		},
		"key3": 1,
	}

	context.HTML(http.StatusOK, "chapter02/map_map.html", mm)
}

// map + struct
func MapStructController(context *gin.Context) {
	user := struct {
		Id   int
		Name string
	}{
		Id:   1,
		Name: "ljs",
	}

	mm := map[string]interface{}{
		"key1": user,
		"key2": "name",
		"key3": []struct {
			Id   int
			Name string
		}{
			{Id: 10, Name: "ljs10"},
			{Id: 20, Name: "ljs20"},
			{Id: 30, Name: "ljs30"},
		},
	}
	context.HTML(http.StatusOK, "chapter02/map_struct.html", mm)
}

// slice 渲染
func SliceController(context *gin.Context) {
	ints := make([]int, 0)
	for i := 0; i < 5; i++ {
		ints = append(ints, i)
	}
	context.HTML(http.StatusOK, "chapter02/slice.html", ints)
}

// slice 渲染
func SliceStructController(context *gin.Context) {
	user_list := []struct {
		Id   int
		Name string
	}{
		{Id: 1, Name: "ljs"},
		{Id: 2, Name: "hsm"},
	}
	context.HTML(http.StatusOK, "chapter02/slice_struct.html", user_list)
}

// 获取路径上的参数【单个参数】
func Param1(context *gin.Context) {
	param := context.Param("id")
	context.String(http.StatusOK, "hello"+param) // 直接输出浏览器
}

// 获取路径上的参数【多个参数】
func Param2(context *gin.Context) {
	id := context.Param("id")
	name := context.Param("name")
	context.String(http.StatusOK, "hello"+" "+id+" "+name) // 直接输出浏览器
}

// 获取路径上的参数
func Param3(context *gin.Context) {
	id := context.Param("id")
	fmt.Println("获取到 param3 ")
	context.String(http.StatusOK, id) // 直接输出浏览器
}

// 获取 ? 后面的参数
func GetQueryData(context *gin.Context) {
	id := context.Query("id")
	name := context.DefaultQuery("name", "ljs") // 如果没有获取到，值是 ljs
	context.String(http.StatusOK, id+", "+name) // 直接输出浏览器
}

// 获取数组
func GetQueryArr(context *gin.Context) {
	values := context.QueryArray("name")
	fmt.Println(len(values))
	context.String(http.StatusOK, "%s", values) // 直接输出浏览器
}

// 获取map
func GetQueryMap(context *gin.Context) {
	values := context.QueryMap("name")
	fmt.Println("values = ", values)
	for k, v := range values {
		fmt.Println(k, " = ", v)
	}
	context.String(http.StatusOK, "%s", values) // 直接输出浏览器
}

// 进入用户添加页面
func ToUserAdd(context *gin.Context) {
	context.HTML(http.StatusOK, "chapter02/user_add.html", nil)
}

// 添加用户
func DoUserAdd(context *gin.Context) {
	// 1、获取参数
	username := context.PostForm("username")
	password := context.DefaultPostForm("password", "ljs")

	values := context.PostFormArray("love")
	for _, v := range values {
		fmt.Printf(v + " ")
	}
	fmt.Println()

	m := context.PostFormMap("user")
	for k, v := range m {
		fmt.Println(k, " = ", v)
	}

	fmt.Println(username)
	fmt.Println(password)
	fmt.Println("添加用户")
	context.String(http.StatusOK, "%s", username+" "+password) // 直接输出浏览器
}

// 需要 将jquery.min.js 放到 /static/js 中
// 使用 ajax 发送 , post 请求
func ToUserAdd3(context *gin.Context) {
	context.HTML(http.StatusOK, "chapter02/user_add3.html", nil)
}

// 添加用户
func DoUserAdd3(context *gin.Context) {
	// 1、获取参数
	username := context.PostForm("username")
	password := context.DefaultPostForm("password", "ljs")

	values := context.PostFormArray("love")
	fmt.Println("values = ", values)
	for _, v := range values {
		fmt.Printf(v + " ")
	}
	fmt.Println()

	m := context.PostFormMap("user")
	for k, v := range m {
		fmt.Println(k, " = ", v)
	}

	fmt.Println(username)
	fmt.Println(password)
	fmt.Println("添加用户")

	context.JSON(http.StatusOK, gin.H{ //  给 ajax 返回数据
		"code": 200,
		"msg":  "提交成功",
	})
}

// 参数绑定
func ToUserAdd4(context *gin.Context) {
	context.HTML(http.StatusOK, "chapter02/user_add4.html", nil)
}

func DoUserAdd4(context *gin.Context) {
	// 1、获取参数
	// 这里可以获取 form 表单的数据，也可以获取 json 的数据
	user := struct {
		Username string            `form:"username" json:"username"`
		Password string            `form:"password" json:"password"`
		Love     []string          `form:"love" json:"love"`
		User     map[string]string `form:"user" json:"user"` // 这个获取不到，可以手动获取
	}{}

	err := context.ShouldBind(&user)
	user.User = context.PostFormMap("user") // 手动获取
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	context.String(http.StatusOK, "%s", user.Username+" "+user.Password) // 直接输出浏览器
}

// 重定向
func RedirectA(ctx *gin.Context) {
	fmt.Println("AAAA")
	ctx.Redirect(http.StatusFound, "/redirectB")
}

func RedirectB(ctx *gin.Context) {
	fmt.Println("BBBB")
	ctx.String(http.StatusOK, "这是B路由")
}
