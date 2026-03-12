package handler

import (
	"net/http"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	fpService   *service.ForgotPasswordService
	userService *service.UserService
}

func NewForgotPasswordHandler(fpService *service.ForgotPasswordService, userService *service.UserService) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{
		fpService:   fpService,
		userService: userService,
	}
}

func (fp *ForgotPasswordHandler) GenerateOTP(ctx *gin.Context) {
	var data models.UserEmail
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input tidak valid",
		})
		return
	}

	emailExist := fp.userService.GetUserByEmail(data.Email)

	if !emailExist {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	fp.fpService.GenerateOTP(data.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "OTP successfully generated and sent to your email.",
	})
}

func (fp *ForgotPasswordHandler) VerifikasiOTP(ctx *gin.Context) {
	var data models.VerifOTP

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data input is not valid",
		})
		return
	}

	err = fp.fpService.VerifikasiOTP(data.Email, data.Code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "OTP successfully verified.",
	})
}
