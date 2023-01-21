package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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
	dst := "upload/" + time.Now().Format("2006-01-02 15:04:05") + "_" + file.Filename
	context.SaveUploadedFile(file, dst)
	context.String(http.StatusOK, "上传成功")
}
