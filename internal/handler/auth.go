package handler

import (
	"log"
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
		log.Printf("Input tidak valid: %v", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input tidak valid",
		})
		return
	}

	// lempar data ke service
	err = a.authService.Register(&data)
	if err != nil {
		log.Printf("Gagal register: %v", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Registrasi berhasil.",
	})
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var data models.UserLogin
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("Input tidak valid: %v", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input tidak valid.",
		})
		return
	}

	user, token, err := a.authService.Login(data.Email, data.Password)
	if err != nil {
		log.Printf("Gagal login: %v", err)
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Email atau password salah.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Login berhasil.",
		Results: gin.H{
			"id":    user.Id,
			"email": user.Email,
			"token": token,
		},
	})
}
