package route

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-docker/controller"
	"github.com/golang-docker/middleware"
)

func InitAuthRoute(r *gin.Engine) {

	auth := controller.InitAuthController()
	login := controller.LoginController{}
	register := controller.RegisterController{}

	r.POST("/login", login.Login)
	r.POST("/register", register.Register)
	r.GET("/profile", middleware.Authorize(), auth.Profile)
	r.GET("/logout", middleware.Authorize(), auth.Logout)
}
