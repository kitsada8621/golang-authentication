package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (u *HomeController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": true,
		"data":    "Welcome to docker + golang",
	})
}
