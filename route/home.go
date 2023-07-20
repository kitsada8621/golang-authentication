package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHomeRoute(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  200,
			"success": true,
			"data":    "Welcome to golang docker",
		})
	})
}
