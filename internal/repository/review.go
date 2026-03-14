package repository

import (
	"context"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type ReviewRepository struct {
	db *pgx.Conn
}

func NewReviewRepository(db *pgx.Conn) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}

func (r *ReviewRepository) GetAllReviews() (*[]models.Reviews, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT users.full_name, reviews.messages, reviews.rating
		FROM reviews
		INNER JOIN users ON reviews.user_id = users.id
	`)
	if err != nil {
		return &[]models.Reviews{}, err
	}
	defer rows.Close()

	reviews, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[models.Reviews])
	if err != nil {
		return &[]models.Reviews{}, err
	}

	return &reviews, nil
}
