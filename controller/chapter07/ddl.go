package chapter07

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// 创建表

func CreateTable(ctx *gin.Context) {
	db.Table("user").CreateTable(&User{}) // 指定表名
	db.CreateTable(&User{})               // 不指定表名,模型后面会加个s
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": 1,
	})
}

// 删除表
func DeleteTable(ctx *gin.Context) {
	//db.DropTable(&User{})       // 使用模型名
	//db.DropTable("users")       // 直接使用表名

	db.DropTableIfExists(&User{}) // 先判断是否存在再删除，可以接受多个参数，模型和字符串都可以
	db.DropTableIfExists("user")  // 先判断是否存在再删除，可以接受多个参数，模型和字符串都可以
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": 1,
	})
}

// 检查表是否存在表
func HasTable(ctx *gin.Context) {
	is_has := db.HasTable(&User{}) // 使用模型

	is_has1 := db.HasTable("users") // 使用表名
	fmt.Println(is_has)
	fmt.Println(is_has1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": 1,
	})
}
