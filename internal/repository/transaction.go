package repository

import (
	"context"
	"fmt"

	"github.com/arif14377/koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetHistoryByUserId(userId string) ([]models.Transaction, error) {
	query := `SELECT id, user_id, trx_code, delivery_method, full_name, email, address, sub_total, tax, total, date, status, payment_method 
              FROM transactions WHERE user_id = $1 ORDER BY date DESC`
	rows, err := r.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) GetTransactionById(id int64, userId string) (*models.Transaction, error) {
	query := `SELECT id, user_id, trx_code, delivery_method, full_name, email, address, sub_total, tax, total, date, status, payment_method 
              FROM transactions WHERE id = $1 AND user_id = $2`
	rows, err := r.db.Query(context.Background(), query, id, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transaction, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		return nil, err
	}

	// Get items
	itemQuery := `SELECT tp.id, tp.product_id, tp.transaction_id, tp.quantity, tp.size_id, tp.variant_id, tp.price, 
                         p.name as product_name, 
                         COALESCE((SELECT path FROM product_images WHERE product_id = p.id LIMIT 1), '') as image, 
                         ps.name as size_name, pv.name as variant_name
                  FROM transaction_product tp
                  JOIN products p ON tp.product_id = p.id
                  LEFT JOIN product_size ps ON tp.size_id = ps.id
                  LEFT JOIN product_variant pv ON tp.variant_id = pv.id
                  WHERE tp.transaction_id = $1`

	itemRows, err := r.db.Query(context.Background(), itemQuery, id)
	if err != nil {
		return nil, err
	}
	defer itemRows.Close()

	items, err := pgx.CollectRows(itemRows, pgx.RowToStructByName[models.TransactionItem])
	if err != nil {
		fmt.Printf("Error collecting items: %v\n", err)
		return &transaction, nil // Return transaction even if items fail
	}

	transaction.Items = items
	return &transaction, nil
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, tx pgx.Tx, trx models.Transaction) (int64, error) {
	query := `INSERT INTO transactions (user_id, trx_code, delivery_method, full_name, email, address, sub_total, tax, total, status, payment_method) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`

	var id int64
	err := tx.QueryRow(ctx, query, trx.UserId, trx.TrxCode, trx.DeliveryMethod, trx.FullName, trx.Email, trx.Address, trx.SubTotal, trx.Tax, trx.Total, trx.Status, trx.PaymentMethod).Scan(&id)
	return id, err
}

func (r *TransactionRepository) CreateTransactionProduct(ctx context.Context, tx pgx.Tx, item models.TransactionItem) error {
	query := `INSERT INTO transaction_product (product_id, transaction_id, quantity, size_id, variant_id, price) 
              VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := tx.Exec(ctx, query, item.ProductId, item.TransactionId, item.Quantity, item.SizeId, item.VariantId, item.Price)
	return err
}

func (r *TransactionRepository) Begin(ctx context.Context) (pgx.Tx, error) {
	return r.db.Begin(ctx)
}
