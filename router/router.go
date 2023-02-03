package router

import (
	"Gin_Vue/controller/chapter01"
	"Gin_Vue/controller/chapter02"
	"Gin_Vue/controller/chapter03"
	"Gin_Vue/controller/chapter04"
	"Gin_Vue/controller/chapter07"
	"Gin_Vue/controller/chapter11"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	chap01 := engine.Group("/chapter01")
	chapter01.Router(chap01)

	chap02 := engine.Group("/chapter02")
	chapter02.Router(chap02)

	chap03 := engine.Group("/chapter03")
	chapter03.Router(chap03)
	//
	chap04 := engine.Group("/chapter04")
	chapter04.Router(chap04)

	chap07 := engine.Group("/chapter07")
	chapter07.Router(chap07)

	chap11 := engine.Group("/chapter11")
	chapter11.Router(chap11)
}
