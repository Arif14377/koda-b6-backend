package handler

import (
	"log"
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
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "List all users",
		Results: u.userService.GetAllUser(),
	})
}

func (u *UserHandler) GetUserByEmail(ctx *gin.Context) {
	// email, isEmailSet := ctx.GetQuery("email")
	var data models.UserEmail
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("Input tidak valid: %v", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input tidak valid",
		})
		return
	}

	exists := u.userService.GetUserByEmail(data.Email)
	if !exists {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User found",
		Results: gin.H{
			"email": data.Email,
		},
	})
}
