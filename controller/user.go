package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-docker/dto"
)

type UserController struct{}

func (u *UserController) Create(ctx *gin.Context) {

	user := dto.UserCreateDto{}
	ctx.BindJSON(user)

	ctx.JSON(http.StatusOK, gin.H{
		"success": 200,
		"status":  true,
		"data":    "Successfully created",
		"obj":     user,
	})
}
