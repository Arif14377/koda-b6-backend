package handler

import (
	"log"
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
		log.Printf("Input tidak valid: %v", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input tidak valid",
		})
		return
	}

	emailExist := fp.userService.GetUserByEmail(data.Email)

	if !emailExist {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "User not found",
		})
		return
	}

	fp.fpService.GenerateOTP(data.Email)

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "OTP successfully generated and sent to your email.",
	})
}

func (fp *ForgotPasswordHandler) VerifikasiOTP(ctx *gin.Context) {
	var data models.VerifOTP

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("Input tidak valid: %v", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Data input is not valid",
		})
		return
	}

	err = fp.fpService.VerifikasiOTP(data.Email, data.Code)
	// fmt.Println(err)
	if err != nil {
		log.Printf("Gagal verifikasi OTP: %v", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Incorrect email or OTP",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "OTP successfully verified.",
	})
}

func (fp *ForgotPasswordHandler) ChangePassword(ctx *gin.Context) {
	var data *models.ForgotPassword
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Printf("Input tidak valid: %v", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Data input is not valid",
		})
		return
	}

	err = fp.fpService.ChangePassword(data)
	if err != nil {
		log.Printf("Gagal ganti password: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Successfully change password",
	})
}
