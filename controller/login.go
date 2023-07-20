package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-docker/dto"
	"github.com/golang-docker/service"
	"github.com/golang-jwt/jwt"
)

var (
	authManager service.AuthService = service.NewAuthService()
)

type LoginController struct{}

func (l *LoginController) Login(ctx *gin.Context) {

	input := &dto.Credentials{}
	if err := ctx.BindJSON(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"status":  422,
			"data":    "Unprocessed Entities",
			"obj":     interface{}(err.Error()),
		})
		return
	}

	attempt := authManager.Attempt(input)
	if attempt == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"success": false,
			"data":    "Invalid username or password.",
			"obj":     nil,
		})
		return
	}

	// generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(5 * time.Minute).Unix()
	claims["authorized"] = true
	claims["user"] = attempt.Id

	secret_key := os.Getenv("JWT_SECRET_KEY")
	ss, err := token.SignedString([]byte(secret_key))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"success": false,
			"data":    "InternalServer",
			"obj":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  200,
		"success": true,
		"data":    "Successfully logged in.",
		"obj": gin.H{
			"token": ss,
			"user":  attempt,
		},
	})
	return

}
