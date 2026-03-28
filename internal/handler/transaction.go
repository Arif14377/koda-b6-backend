package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) GetHistory(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	if userId == "" {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	history, err := h.service.GetHistory(userId)
	if err != nil {
		log.Printf("Gagal mendapatkan riwayat transaksi: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Berhasil mendapatkan riwayat transaksi.",
		Results: history,
	})
}

func (h *TransactionHandler) GetDetail(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	if userId == "" {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "ID transaksi tidak valid",
		})
		return
	}

	detail, err := h.service.GetDetail(id, userId)
	if err != nil {
		log.Printf("Gagal mendapatkan detail transaksi: %v", err)
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Transaksi tidak ditemukan.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Berhasil mendapatkan detail transaksi.",
		Results: detail,
	})
}

func (h *TransactionHandler) Checkout(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	if userId == "" {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	var trx models.Transaction
	if err := ctx.ShouldBindJSON(&trx); err != nil {
		log.Printf("Gagal bind JSON checkout: %v. Body: %v", err, ctx.Request.Body)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Input tidak valid",
		})
		return
	}

	err := h.service.Checkout(ctx.Request.Context(), userId, trx)
	if err != nil {
		log.Printf("Gagal melakukan checkout: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Checkout berhasil.",
	})
}
