package middlewares

import (
	"os"
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
		user_name, err := data.ValidateToken(access_token, []byte(os.Getenv("ATS")))
		if err == nil {
			ctx.Set("user_name", user_name)
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(401, gin.H{"error": "Access token invalid. Please re-login"})
	}
}
