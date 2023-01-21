package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// form表单， 单文件

func ToUpload1(context *gin.Context) {
	context.HTML(http.StatusOK, "chapter02/test_upload1.html", nil)
}

func DoUpload1(context *gin.Context) {
	file, err := context.FormFile("file") // 和前端的 <input type="file" name="file">  对应
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Filename) // 文件名

	// 保存文件
	//dst := "upload/" + time.Now().Format("2006-01-02 15:04:05") + "_" + file.Filename
	time_unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int, 10)
	dst := "upload/" + time_unix_str + file.Filename
	context.SaveUploadedFile(file, dst)
	context.String(http.StatusOK, "上传成功")
}

// form表单， 多文件

func ToUpload2(context *gin.Context) {
	context.HTML(http.StatusOK, "chapter02/test_upload2.html", nil)
}

func DoUpload2(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	file_list := form.File["file"]
	for _, v := range file_list {
		// 保存文件
		//dst := "upload/" + time.Now().Format("2006-01-02 15:04:05") + "_" + file.Filename
		time_unix_int := time.Now().Unix()
		time_unix_str := strconv.FormatInt(time_unix_int, 10)
		dst := "upload/" + time_unix_str + v.Filename
		context.SaveUploadedFile(v, dst)
	}
	context.String(http.StatusOK, "上传成功")
}

// ajax 单文件
func ToUpload3(context *gin.Context) {
	context.HTML(http.StatusOK, "chapter02/test_upload3.html", nil)
}

func DoUpload3(context *gin.Context) {
	name := context.PostForm("name") // <input type="text" id="name">
	fmt.Println("name = ", name)
	file, err := context.FormFile("file") // <input type="file" id="file">
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(file.Filename) // 文件名

	// 保存文件
	//dst := "upload/" + time.Now().Format("2006-01-02 15:04:05") + "_" + file.Filename
	time_unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int, 10)
	dst := "upload/" + time_unix_str + file.Filename
	context.SaveUploadedFile(file, dst)
	context.String(http.StatusOK, "上传成功")
}

// ajax 多文件
func ToUpload4(context *gin.Context) {
	context.HTML(http.StatusOK, "chapter02/test_upload4.html", nil)
}

func DoUpload4(context *gin.Context) {
	form, err := context.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	name := context.PostForm("name")
	fmt.Println("name = ", name)
	file_list := form.File["file"]
	for _, v := range file_list {
		// 保存文件
		//dst := "upload/" + time.Now().Format("2006-01-02 15:04:05") + "_" + file.Filename
		time_unix_int := time.Now().Unix()
		time_unix_str := strconv.FormatInt(time_unix_int, 10)
		dst := "upload/" + time_unix_str + v.Filename
		context.SaveUploadedFile(v, dst)
	}
	context.String(http.StatusOK, "上传成功")
}
