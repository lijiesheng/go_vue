package chapter11

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestAxios(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": true,
	})
}
