package chapter07

import (
	"Gin_Vue/controller/chapter07/relate_tables"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 添加数据
func AddData(ctx *gin.Context) {
	db.Create(&User{Name: "hallen", Age: 18}) // 对应的 user 表数据
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": 1,
	})
}

// 查询数据
func SearchData(ctx *gin.Context) {
	var user User
	db.First(&user, 1) // 默认使用id字段
	fmt.Println(user)
	db.First(&user, "name=?", "hallen") // 查询指定字段
	fmt.Println(user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": 1,
	})
}

// 更新数据
func UpdateData(ctx *gin.Context) {
	var user User
	db.First(&user, 2) // 默认使用id字段

	db.Model(&user).Update("age", 20)
	db.Model(&user).Update("name", "zhangsan")
	// 这里的Model是查询出来的结构体对象user，不是模型
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": 1,
	})
}

func DeleteData(ctx *gin.Context) {
	var user User
	db.First(&user, 1) // 默认使用id字段

	db.Delete(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": 1,
	})
}

// 一对一
func ForeignKey(ctx *gin.Context) {
	db.AutoMigrate(&relate_tables.User{}, &relate_tables.UserProfile{})
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": 1,
	})
}
