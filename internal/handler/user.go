package handler

import (
	"net/http"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (u *UserHandler) GetAllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List all users",
		"results": u.userService.GetAllUser(),
	})
}

func (u *UserHandler) GetUserByEmail(ctx *gin.Context) {
	// email, isEmailSet := ctx.GetQuery("email")
	var data models.UserEmail
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input tidak valid",
		})
	}

	exists := u.userService.GetUserByEmail(data.Email)
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User found",
		"results": gin.H{
			"email": data.Email,
		},
	})
}
