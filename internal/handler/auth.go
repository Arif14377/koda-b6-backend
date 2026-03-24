package handler

import (
	"fmt"
	"net/http"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var data models.UserRegister
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input tidak valid",
		})
		fmt.Printf("Input tidak valid: %v", err)
		return
	}

	// lempar data ke service
	err = a.authService.Register(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Registrasi berhasil.",
	})
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var data models.UserLogin
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input tidak valid.",
		})
		fmt.Printf("Input tidak valid: %v", err)
		return
	}

	user, err := a.authService.Login(data.Email, data.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Email atau password salah.",
		})
		fmt.Printf("Input tidak valid: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login berhasil.",
		"results": gin.H{
			"email": user.Email,
		},
	})
}
