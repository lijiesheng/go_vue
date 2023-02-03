package chapter11

import "github.com/gin-gonic/gin"

func Router(chap11 *gin.RouterGroup) {
	chap11.GET("/aaa", TestAxios)
}
