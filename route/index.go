package route

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {

	// router
	InitAuthRoute(r)
	InitHomeRoute(r)

}
