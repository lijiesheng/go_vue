package chapter01

import "github.com/gin-gonic/gin"

var chapter01RouterGroup *gin.RouterGroup

func Router(chap01 *gin.RouterGroup) {
	//chapter01RouterGroup = routerGroup
	chap01.GET("/index", Index)
	chap01.GET("/news", News)
}
