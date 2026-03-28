package service

import (
	"context"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) GetCartByUserId(ctx context.Context, userId string) ([]models.CartItemResponse, error) {
	return s.repo.GetCartByUserId(ctx, userId)
}

func (s *CartService) AddToCart(ctx context.Context, cart models.Cart) error {
	return s.repo.AddToCart(ctx, cart)
}

func (s *CartService) UpdateQuantity(ctx context.Context, cartId int, userId string, quantity int) error {
	return s.repo.UpdateQuantity(ctx, cartId, userId, quantity)
}

func (s *CartService) RemoveFromCart(ctx context.Context, cartId int, userId string) error {
	return s.repo.RemoveFromCart(ctx, cartId, userId)
}

func (s *CartService) ClearCart(ctx context.Context, userId string) error {
	return s.repo.ClearCart(ctx, userId)
}
