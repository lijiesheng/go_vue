package chapter11

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestAxios(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "hello",
	})
}

// 字符串渲染
func User(context *gin.Context) {
	name := "ljs"
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": name,
	})
}

// 结构体渲染
func UserInfoStruct(context *gin.Context) {
	user := struct {
		Id   int
		Name string
	}{Id: 1, Name: "ljs"}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": user,
	})
}

// 数组渲染
func ArrController(context *gin.Context) {
	nums := []int{1, 2, 3, 4, 5}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": nums,
	})
}

// 结构体数组渲染
func ArrStructController(context *gin.Context) {
	userArr := []struct {
		Id   int
		Name string
	}{{Id: 1, Name: "ljs"}, {Id: 2, Name: "hsm"}}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": userArr,
	})
}

// map 的渲染
func MapController(context *gin.Context) {
	m := map[string]int{
		"chinese": 80,
		"math":    100,
		"english": 120,
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": m,
	})
}

// map+map
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

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": mm,
	})
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
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": mm,
	})
}

// slice 渲染
func SliceController(context *gin.Context) {
	ints := make([]int, 0)
	for i := 0; i < 5; i++ {
		ints = append(ints, i)
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": ints,
	})
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
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": user_list,
	})
}

func DynamicRouter(context *gin.Context) {
	user_list := []struct {
		//Id      int      `json:"id"`
		Key     string   `json:"key"`
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Address string   `json:"address"`
		Tags    []string `json:"tags"`
	}{
		{Key: "1", Name: "John Brown", Age: 32, Address: "New York No. 1 Lake Park", Tags: []string{"nice", "developer"}},
		{Key: "2", Name: "Jim Green", Age: 42, Address: "London No. 1 Lake Park", Tags: []string{"loser"}},
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": user_list,
	})
}

func GetBookDetails(context *gin.Context) {
	// 获取 id
	id := context.Query("name")
	fmt.Println(id)
	user_list := struct {
		//Id      int      `json:"id"`
		Key     string   `json:"key"`
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Address string   `json:"address"`
		Tags    []string `json:"tags"`
	}{
		Key: "1", Name: "John Brown", Age: 32, Address: "New York No. 1 Lake Park", Tags: []string{"nice", "developer"},
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功",
		"user": user_list,
	})
}
