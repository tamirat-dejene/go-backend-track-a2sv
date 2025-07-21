package middlewares

import (
	"fmt"
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

		refresh_token, err := ctx.Cookie("refresh_token")
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Refresh token required"})
			return
		}

		user_name, err = data.ValidateToken(refresh_token, []byte(os.Getenv("RTS")))
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid refresh token, please log in again"})
			return
		}

		token, err := data.SignUser(&data.JWTPayload{
			UserName: user_name,
			Exp:      os.Getenv("ATE"),
		}, os.Getenv("ATE"))

		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"error": "Failed to refresh access token"})
			return
		}

		ctx.Header("Authorization", fmt.Sprintf("Bearer %s", token))
		ctx.Set("user_name", user_name)
		ctx.Next()
	}
}
