package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		header := ctx.GetHeader("Authorization")
		tokens := strings.Split(header, " ")
		if len(tokens) != 2 || strings.ToLower(tokens[0]) != "bearer" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"status":  401,
				"success": false,
				"data":    "The token is empty or null",
				"obj":     nil,
			})
			return
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")
		token, err := jwt.Parse(tokens[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"status":  401,
				"success": false,
				"data":    err.Error(),
				"obj":     nil,
			})
			return
		}

		// claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{
				"status":  401,
				"success": false,
				"data":    "Invalid Token",
				"obj":     err,
			})
			return
		}

		ctx.Set("user", claims["user"])
		ctx.Next()
	}
}
