package controllers

import (
	"net/http"
	"os"
	"t4/taskmanager/constants"
	"t4/taskmanager/data"
	"t4/taskmanager/models"

	"github.com/gin-gonic/gin"
)

func LoginUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	_, err := data.LoginUser(&user)

	if err != nil {
		if err.Error() == constants.INVALID_CREDETNTIALS {
			ctx.JSON(401, gin.H{"message": "Invalid credentials"})
			return
		} else if err.Error() == constants.NOT_FOUND {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error", "error": err.Error()})
		return
	}

	// Sign user, with jwt
	access_token, err1 := data.SignUser(&data.JWTPayload{
		UserName: user.UserName,
		Exp:      os.Getenv("ATE"),
	}, os.Getenv("ATS"))

	refresh_token, err2 := data.SignUser(&data.JWTPayload{
		UserName: user.UserName,
		Exp:      os.Getenv("RTE"),
	}, os.Getenv("RTS"))

	if err1 != nil || err2 != nil {
		var tokenErr error
		if err1 != nil {
			tokenErr = err1
		} else {
			tokenErr = err2
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token", "error": tokenErr.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+access_token)
	ctx.SetCookie("refresh_token", refresh_token, 86400, "/", "", false, true)
	ctx.IndentedJSON(200, gin.H{
		"message": "Welcome back, " + user.UserName + "!",
		"token_usage": gin.H{
			"access_token":  "The access token is included in the 'Authorization' header as 'Bearer <token>' for each authenticated request.",
			"refresh_token": "The refresh token is set as an HTTP-only cookie named 'refresh_token'. Use it to obtain a new access token when the current one expires.",
			"note":          "Do not share your tokens. The access token is required for accessing protected endpoints, while the refresh token is used for token renewal.",
		},
		"user": gin.H{
			"username": user.UserName,
		},
		"access_token":  access_token,
		"refresh_token": refresh_token,
	})
}

func RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": "Validation failed: please check your input fields", "error": err.Error()})
		return
	}

	id, err := data.RegisterUser(&user)
	if err != nil {
		if err.Error() == constants.DUPLICATE_USERNAME {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "Username already exists. Please choose a different username or login if you already have an account.",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Registration failed", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully. Please proceed to login.",
		"user_id": id,
	})
}

func DeleteUser(ctx *gin.Context) {
	var id string

	if err := ctx.ShouldBindUri(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID", "error": err.Error()})
		return
	}

	id = ctx.Param("id")
	err := data.DeleteUser(id)

	if err != nil {
		if err.Error() == constants.NOT_FOUND {
			// user not found
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "User not found. Please verify the user ID and try again.",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Deletion failed", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully.",
		"user_id": id,
	})
}

func GetUsers(ctx *gin.Context) {
	users, err := data.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve users", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func RefreshUserToken(ctx *gin.Context) {

}

func LogoutUser(ctx *gin.Context) {
}
