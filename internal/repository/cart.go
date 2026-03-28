package repository

import (
	"context"
	"errors"

	"github.com/arif14377/koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) GetCartByUserId(ctx context.Context, userId string) ([]models.CartItemResponse, error) {
	query := `
		SELECT 
			c.id, c.product_id, p.name, c.quantity, 
			(p.price + COALESCE(ps.add_price, 0) + COALESCE(pv.add_price, 0)) as price,
			COALESCE((SELECT path FROM product_images WHERE product_id = p.id LIMIT 1), '') as image,
			ps.name as size, pv.name as variant,
			c.size_id, c.variant_id,
			(c.quantity * (p.price + COALESCE(ps.add_price, 0) + COALESCE(pv.add_price, 0))) as total
		FROM cart c
		JOIN products p ON c.product_id = p.id
		LEFT JOIN product_size ps ON c.size_id = ps.id
		LEFT JOIN product_variant pv ON c.variant_id = pv.id
		WHERE c.user_id = $1
		ORDER BY c.created_at DESC
	`
	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.CartItemResponse])
}

func (r *CartRepository) AddToCart(ctx context.Context, cart models.Cart) error {
	// Check if item already exists in cart with same size and variant
	var existingId int
	var existingQty int
	checkQuery := `
		SELECT id, quantity FROM cart 
		WHERE user_id = $1 AND product_id = $2 
		AND (size_id IS NOT DISTINCT FROM $3)
		AND (variant_id IS NOT DISTINCT FROM $4)
	`
	err := r.db.QueryRow(ctx, checkQuery, cart.UserId, cart.ProductId, cart.SizeId, cart.VariantId).Scan(&existingId, &existingQty)

	if err == nil {
		// Update quantity
		updateQuery := `UPDATE cart SET quantity = $1 WHERE id = $2`
		_, err = r.db.Exec(ctx, updateQuery, existingQty+cart.Quantity, existingId)
		return err
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	// Insert new item
	insertQuery := `
		INSERT INTO cart (user_id, product_id, quantity, size_id, variant_id)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = r.db.Exec(ctx, insertQuery, cart.UserId, cart.ProductId, cart.Quantity, cart.SizeId, cart.VariantId)
	return err
}

func (r *CartRepository) UpdateQuantity(ctx context.Context, cartId int, userId string, quantity int) error {
	query := `UPDATE cart SET quantity = $1 WHERE id = $2 AND user_id = $3`
	_, err := r.db.Exec(ctx, query, quantity, cartId, userId)
	return err
}

func (r *CartRepository) RemoveFromCart(ctx context.Context, cartId int, userId string) error {
	query := `DELETE FROM cart WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(ctx, query, cartId, userId)
	return err
}

func (r *CartRepository) ClearCart(ctx context.Context, userId string) error {
	query := `DELETE FROM cart WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, userId)
	return err
}
