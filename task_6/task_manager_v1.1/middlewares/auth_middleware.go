package middlewares

import (
	"strings"
	"t4/taskmanager/data"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth_header := ctx.GetHeader("Authorization")
		if auth_header == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Authorization token required"})
			return
		}

		auth_header_parts := strings.Split(auth_header, " ")
		if len(auth_header_parts) != 2 || strings.ToLower(auth_header_parts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		access_token := auth_header_parts[1]
		refresh_token, err := ctx.Cookie("refresh_token")

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Refresh token required"})
			return
		}

		user_name, err := data.ValidateToken(access_token, []byte(data.ATS))
		if err == nil {
			ctx.Set("user_name", user_name)
			ctx.Next()
			return
		}

		user_name, err = data.ValidateToken(refresh_token, []byte(data.RTS))
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid refresh token, please log in again"})
			return
		}
		ctx.Set("user_name", user_name)
		ctx.Next()
	}
}
