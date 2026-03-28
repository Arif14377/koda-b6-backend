package service

import (
	"context"
	"fmt"
	"time"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/arif14377/koda-b6-backend/internal/repository"
)

type TransactionService struct {
	repo     *repository.TransactionRepository
	cartRepo *repository.CartRepository
}

func NewTransactionService(repo *repository.TransactionRepository, cartRepo *repository.CartRepository) *TransactionService {
	return &TransactionService{repo: repo, cartRepo: cartRepo}
}

func (s *TransactionService) GetHistory(userId string) ([]models.Transaction, error) {
	return s.repo.GetHistoryByUserId(userId)
}

func (s *TransactionService) GetDeliveryMethods() ([]models.DeliveryMethod, error) {
	return s.repo.GetDeliveryMethods()
}

func (s *TransactionService) GetDetail(id int64, userId string) (*models.Transaction, error) {
	return s.repo.GetTransactionById(id, userId)
}

func (s *TransactionService) Checkout(ctx context.Context, userId string, trx models.Transaction) error {
	// 1. Get Cart items
	cartItems, err := s.cartRepo.GetCartByUserId(ctx, userId)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	// 2. Begin Transaction
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// 3. Create Transaction record
	trx.UserId = userId
	trx.TrxCode = fmt.Sprintf("TRX-%d", time.Now().Unix())
	trx.Status = "Pending"

	trxId, err := s.repo.CreateTransaction(ctx, tx, trx)
	if err != nil {
		return err
	}

	// 4. Create Transaction Products
	for _, item := range cartItems {
		trxItem := models.TransactionItem{
			ProductId:     item.ProductId,
			TransactionId: trxId,
			Quantity:      item.Quantity,
			SizeId:        item.SizeId,
			VariantId:     item.VariantId,
			Price:         item.Price,
		}
		if err := s.repo.CreateTransactionProduct(ctx, tx, trxItem); err != nil {
			return err
		}
	}

	// 5. Clear Cart
	if err := s.cartRepo.ClearCart(ctx, userId); err != nil {
		return err
	}

	// 6. Commit Transaction
	return tx.Commit(ctx)
}
