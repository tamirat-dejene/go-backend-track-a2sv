package controllers

import (
	"net/http"
	"t7/taskmanager/Delivery/bootstrap"
	domain "t7/taskmanager/Domain"
	infrustructure "t7/taskmanager/Infrustructure"
	"t7/taskmanager/Infrustructure/constants"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

func (uc *UserController) Delete(ctx *gin.Context) {
	var id string

	if err := ctx.ShouldBindUri(id); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorRespone{Message: "Invalid user ID", Error: err.Error()})
		return
	}

	id = ctx.Param("id")
	err := uc.UserUsecase.Delete(ctx, id)

	if err != nil {
		if err.Error() == constants.NOT_FOUND {
			// user not found
			ctx.JSON(http.StatusNotFound, domain.ErrorRespone{
				Message: "User not found. Please verify the user ID and try again.",
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, domain.ErrorRespone{Message: "Deletion failed", Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "User deleted successfully.",
	})
}

func (uc *UserController) Register(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.IndentedJSON(400, gin.H{"message": "Validation failed: please check your input fields", "error": err.Error()})
		return
	}

	id, err := uc.UserUsecase.Register(ctx, &user)
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

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserUsecase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve users", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(400, domain.ErrorRespone{Message: "Invalid request", Error: err.Error()})
		return
	}

	_, err := uc.UserUsecase.Login(ctx, &user)

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
	access_token, err1 := infrustructure.CreateToken(user.UserName, uc.Env.AccTE, uc.Env.AccTS)
	refresh_token, err2 := infrustructure.CreateToken(user.UserName, uc.Env.RefTE, uc.Env.RefTS)

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

func (uc *UserController) Refresh(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token missing or invalid. Try to re login"})
		return
	}

	user_name, err := infrustructure.IsAuthorized(refreshToken, []byte(uc.Env.RefTS))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid refresh token", "error": err.Error()})
		return
	}

	accessToken, err := infrustructure.CreateToken(user_name, uc.Env.AccTE, uc.Env.AccTS)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate access token", "error": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+accessToken)
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Access token refreshed successfully.",
		"access_token": accessToken,
	})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	ctx.SetCookie("refresh_token", "", -1, "/", "", false, true)
	ctx.Header("Authorization", "")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout successful. Tokens have been cleared.",
	})
}