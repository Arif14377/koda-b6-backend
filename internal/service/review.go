package service

import (
	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/repository"
)

type ReviewService struct {
	reviewRepo *repository.ReviewRepository
}

func NewReviewService(reviewRepo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
	}
}

func (r *ReviewService) GetAllReviews() (*[]models.Reviews, error) {
	reviews, err := r.reviewRepo.GetAllReviews()
	if err != nil {
		return &[]models.Reviews{}, err
	}

	return reviews, nil
}
