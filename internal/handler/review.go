package handler

import (
	"log"
	"net/http"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler(reviewService *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

func (r *ReviewHandler) GetAllReviews(ctx *gin.Context) {
	reviews, err := r.reviewService.GetAllReviews()
	if err != nil {
		log.Printf("Gagal mendapatkan semua review: %v", err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Ada kesalahan pada server.",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "List all reviews",
		Results: reviews,
	})
}
