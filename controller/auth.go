package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-docker/service"
)

type AuthController struct {
	userManager service.UserService
}

func InitAuthController() *AuthController {
	return &AuthController{
		userManager: service.NewUserService(),
	}
}

func (a *AuthController) Profile(ctx *gin.Context) {

	claims, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": true,
		"data":    "profile",
		"obj":     claims,
	})
}

func (a *AuthController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": true,
		"data":    "Success logged out",
	})
}
