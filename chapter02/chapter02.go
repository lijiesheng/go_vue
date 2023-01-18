package chapter02

import (
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
