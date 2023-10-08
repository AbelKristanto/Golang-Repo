package middleware

import (
	"diary-api/model/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := helper.ValidateJWT(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
