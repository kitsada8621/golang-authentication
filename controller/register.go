package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-docker/dto"
	"github.com/golang-docker/model"
	"github.com/golang-docker/service"
)

var (
	hashManager service.AuthService = service.NewAuthService()
	userManager service.UserService = service.NewUserService()
)

type RegisterController struct{}

func (r *RegisterController) Register(ctx *gin.Context) {

	request := &dto.RegisterDto{}
	if err := ctx.BindJSON(request); err != nil {
		errors := map[string]interface{}{"data": err.Error()}
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  422,
			"success": false,
			"data":    "Bad Request",
			"obj":     errors,
		})
		return
	}

	_, err := userManager.FindOne("email", request.Email)
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"success": false,
			"data":    "The email is already exists.",
			"obj":     err,
		})
		return
	}

	payload := &model.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Username: request.Username,
		Password: request.Password,
	}

	_, err = userManager.Create(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"success": false,
			"data":    "Failed create user",
			"obj":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": true,
		"data":    "Successfully created",
		"obj":     nil,
	})
	return

}
