package chapter03

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestSyncTpl(ctx *gin.Context) {
	m := map[string]interface{}{
		"name": "aaab",
		"arr":  []int{11, 22, 33, 44},
	}
	ctx.HTML(http.StatusOK, "chapter03/test.html", m)
}
