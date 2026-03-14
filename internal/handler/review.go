package handler

import (
	"net/http"

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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":  false,
			"messages": "Failed to get all reviews.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"messages": "List all reviews",
		"results":  reviews,
	})
}
